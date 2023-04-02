package observability

import (
	"context"
	"testing"

	"github.com/prixfixeco/backend/internal/observability/metrics/config"
	"github.com/prixfixeco/backend/internal/observability/metrics/prometheus"
	tracingcfg "github.com/prixfixeco/backend/internal/observability/tracing/config"
	"github.com/prixfixeco/backend/internal/observability/tracing/jaeger"

	"github.com/stretchr/testify/assert"
)

func TestConfig_ValidateWithContext(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		cfg := &Config{
			Tracing: tracingcfg.Config{
				Provider: tracingcfg.ProviderJaeger,
				Jaeger: &jaeger.Config{
					CollectorEndpoint:         "0.0.0.0",
					ServiceName:               t.Name(),
					SpanCollectionProbability: 1,
				},
			},
			Metrics: config.Config{
				Provider: config.ProviderPrometheus,
				Prometheus: &prometheus.Config{
					RuntimeMetricsCollectionInterval: config.DefaultMetricsCollectionInterval,
				},
			},
		}

		assert.NoError(t, cfg.ValidateWithContext(ctx))
	})
}
