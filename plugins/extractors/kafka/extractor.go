package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/odpf/meteor/core"
	"github.com/odpf/meteor/core/extractor"
	"github.com/odpf/meteor/plugins"

	"github.com/odpf/meteor/proto/odpf/meta"
	"github.com/odpf/meteor/utils"
)

type Config struct {
	Broker string `mapstructure:"broker" validate:"required"`
}

type Extractor struct {
	// internal states
	out    chan<- interface{}
	client *kafka.AdminClient

	// dependencies
	logger plugins.Logger
}

func New(logger plugins.Logger) *Extractor {
	return &Extractor{
		logger: logger,
	}
}

func (e *Extractor) Extract(ctx context.Context, configMap map[string]interface{}, out chan<- interface{}) (err error) {
	e.out = out

	// build config
	var config Config
	err = utils.BuildConfig(configMap, &config)
	if err != nil {
		return extractor.InvalidConfigError{}
	}

	// create client
	client, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"metadata.broker.list": config.Broker,
	})
	if err != nil {
		return err
	}
	defer client.Close()
	e.client = client

	return e.extract()
}

// Extract and output metadata from all topics in a broker
func (e *Extractor) extract() (err error) {
	// Fetch kafka metadata
	metadata, err := e.client.GetMetadata(nil, true, 1000)
	if err != nil {
		return
	}

	// Build and output topic metadata from topic name
	for topic_name := range metadata.Topics {
		e.out <- e.buildTopic(topic_name)
	}

	return
}

// Build topic metadata model using a topic name
func (e *Extractor) buildTopic(topic_name string) meta.Topic {
	return meta.Topic{
		Urn:    topic_name,
		Name:   topic_name,
		Source: "kafka",
	}
}

func init() {
	if err := extractor.Catalog.Register("kafka", func() core.Extractor {
		return New(plugins.Log)
	}); err != nil {
		panic(err)
	}
}
