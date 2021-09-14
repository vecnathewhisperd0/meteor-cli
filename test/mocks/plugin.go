package mocks

import (
	"context"

	"github.com/odpf/meteor/models"
	"github.com/odpf/meteor/plugins"
	"github.com/stretchr/testify/mock"
)

type Plugin struct {
	mock.Mock
}

func (m *Plugin) Info() plugins.Info {
	args := m.Called()
	return args.Get(0).(plugins.Info)
}

func (m *Plugin) Validate(config map[string]interface{}) error {
	args := m.Called(config)
	return args.Error(0)
}

func (m *Plugin) Init(ctx context.Context, config map[string]interface{}) error {
	args := m.Called(ctx, config)
	return args.Error(0)
}

type Extractor struct {
	Plugin
	records []models.Record
}

func NewExtractor() *Extractor {
	return &Extractor{}
}

func (m *Extractor) SetEmit(records []models.Record) {
	m.records = records
}

func (m *Extractor) Extract(ctx context.Context, push plugins.PushFunc) error {
	args := m.Called(ctx, push)

	for _, r := range m.records {
		push(r)
	}

	return args.Error(0)
}

type Processor struct {
	Plugin
}

func NewProcessor() *Processor {
	return &Processor{}
}

func (m *Processor) Process(ctx context.Context, src models.Record) (models.Record, error) {
	args := m.Called(ctx, src)
	return args.Get(0).(models.Record), args.Error(1)
}

type Sink struct {
	Plugin
}

func NewSink() *Sink {
	return &Sink{}
}

func (m *Sink) Sink(ctx context.Context, batch []models.Record) error {
	args := m.Called(ctx, batch)
	return args.Error(0)
}

type PushFunc struct {
	data []models.Record
}

func NewPushFunc() *PushFunc {
	return &PushFunc{}
}

func (m *PushFunc) Push(record models.Record) {
	m.data = append(m.data, record)
}

func (m *PushFunc) Get() []models.Record {
	return m.data
}
