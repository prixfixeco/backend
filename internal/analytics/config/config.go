package config

import (
	"context"
	"strings"

	"github.com/dinnerdonebetter/backend/internal/analytics"
	"github.com/dinnerdonebetter/backend/internal/analytics/rudderstack"
	"github.com/dinnerdonebetter/backend/internal/analytics/segment"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// ProviderSegment represents Segment.
	ProviderSegment = "segment"
	// ProviderRudderstack represents Rudderstack.
	ProviderRudderstack = "rudderstack"
)

type (
	// Config is the configuration structure.
	Config struct {
		Segment     *segment.Config     `json:"segment"     mapstructure:"segment"     toml:"segment,omitempty"`
		Rudderstack *rudderstack.Config `json:"rudderstack" mapstructure:"rudderstack" toml:"rudderstack,omitempty"`
		Provider    string              `json:"provider"    mapstructure:"provider"    toml:"provider,omitempty"`
	}
)

var _ validation.ValidatableWithContext = (*Config)(nil)

// ValidateWithContext validates a Config struct.
func (cfg *Config) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, cfg,
		validation.Field(&cfg.Provider, validation.In(ProviderSegment, ProviderRudderstack)),
		validation.Field(&cfg.Segment, validation.When(cfg.Provider == ProviderSegment, validation.Required)),
		validation.Field(&cfg.Rudderstack, validation.When(cfg.Provider == ProviderRudderstack, validation.Required)),
	)
}

// ProvideCollector provides a collector.
func (cfg *Config) ProvideCollector(logger logging.Logger, tracerProvider tracing.TracerProvider) (analytics.EventReporter, error) {
	switch strings.ToLower(strings.TrimSpace(cfg.Provider)) {
	case ProviderSegment:
		return segment.NewSegmentEventReporter(logger, tracerProvider, cfg.Segment.APIToken)
	case ProviderRudderstack:
		return rudderstack.NewRudderstackEventReporter(logger, tracerProvider, cfg.Rudderstack)
	default:
		return analytics.NewNoopEventReporter(), nil
	}
}
