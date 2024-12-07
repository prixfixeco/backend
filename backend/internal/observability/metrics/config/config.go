package config

import (
	"context"
	"strings"

	"github.com/dinnerdonebetter/backend/internal/observability/metrics"
	"github.com/dinnerdonebetter/backend/internal/observability/metrics/otelgrpc"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// ProviderOtel represents the open source tracing server.
	ProviderOtel = "otelgrpc"
)

type (
	// Config contains settings related to tracing.
	Config struct {
		_ struct{} `json:"-"`

		Otel     *otelgrpc.Config `json:"otelgrpc,omitempty" toml:"otelgrpc,omitempty"`
		Provider string           `json:"provider,omitempty" toml:"provider,omitempty"`
	}
)

// ProvideMetricsProvider provides a metrics provider.
func (c *Config) ProvideMetricsProvider(ctx context.Context) (metrics.Provider, error) {
	p := strings.TrimSpace(strings.ToLower(c.Provider))

	switch p {
	case ProviderOtel:
		return otelgrpc.ProvideMetricsProvider(ctx, c.Otel)
	default:
		return metrics.NewNoopMetricsProvider(), nil
	}
}

var _ validation.ValidatableWithContext = (*Config)(nil)

// ValidateWithContext validates the config struct.
func (c *Config) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, c,
		validation.Field(&c.Provider, validation.In("", ProviderOtel)),
		validation.Field(&c.Otel, validation.When(c.Provider == ProviderOtel, validation.Required).Else(validation.Nil)),
	)
}