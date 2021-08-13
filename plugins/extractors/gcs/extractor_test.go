//+build integration

package gcs_test

import (
	"context"
	"testing"

	"github.com/odpf/meteor/plugins"
	"github.com/odpf/meteor/plugins/extractors/gcs"
	logger "github.com/odpf/salt/log"
	"github.com/stretchr/testify/assert"
)

var log = logger.NewLogrus()

func TestExtract(t *testing.T) {
	t.Run("should return error if no project_id in config", func(t *testing.T) {
		err := gcs.New(log).Extract(context.TODO(), map[string]interface{}{
			"wrong-config": "sample-project",
		}, make(chan interface{}))

		assert.Equal(t, plugins.InvalidConfigError{}, err)
	})
}
