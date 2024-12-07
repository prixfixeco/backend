package otel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_ValidateWithContext(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		cfg := &Config{
			CollectorEndpoint: t.Name(),
			ServiceName:       t.Name(),
		}

		assert.NoError(t, cfg.ValidateWithContext(ctx))
	})
}