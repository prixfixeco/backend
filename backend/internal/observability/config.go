package observability

import (
	"context"

	loggingcfg "github.com/dinnerdonebetter/backend/internal/observability/logging/config"
	metricscfg "github.com/dinnerdonebetter/backend/internal/observability/metrics/config"
	tracingcfg "github.com/dinnerdonebetter/backend/internal/observability/tracing/config"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	// Config contains settings about how we report our metrics.
	Config struct {
		_ struct{} `json:"-"`

		Logging loggingcfg.Config `envPrefix:"LOGGING_" json:"logging"`
		Metrics metricscfg.Config `envPrefix:"METRICS_" json:"metrics"`
		Tracing tracingcfg.Config `envPrefix:"TRACING_" json:"tracing"`
	}
)

var _ validation.ValidatableWithContext = (*Config)(nil)

// ValidateWithContext validates a Config struct.
func (cfg *Config) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, cfg,
		validation.Field(&cfg.Logging),
		validation.Field(&cfg.Metrics),
		validation.Field(&cfg.Tracing),
	)
}
