package config

import (
	"context"
	"net/http"
	"strings"

	"github.com/prixfixeco/backend/internal/email"
	"github.com/prixfixeco/backend/internal/email/mailgun"
	"github.com/prixfixeco/backend/internal/email/mailjet"
	"github.com/prixfixeco/backend/internal/email/sendgrid"
	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/tracing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// ProviderSendgrid represents SendGrid.
	ProviderSendgrid = "sendgrid"
	// ProviderMailgun represents Mailgun.
	ProviderMailgun = "mailgun"
	// ProviderMailjet represents Mailjet.
	ProviderMailjet = "mailjet"
)

type (
	// Config is the configuration structure.
	Config struct {
		Sendgrid *sendgrid.Config `json:"sendgrid" mapstructure:"sendgrid" toml:"sendgrid,omitempty"`
		Mailgun  *mailgun.Config  `json:"mailgun" mapstructure:"mailgun" toml:"mailgun,omitempty"`
		Mailjet  *mailjet.Config  `json:"mailjet" mapstructure:"mailjet" toml:"mailjet,omitempty"`
		Provider string           `json:"provider" mapstructure:"provider" toml:"provider,omitempty"`
	}
)

var _ validation.ValidatableWithContext = (*Config)(nil)

// ValidateWithContext validates a Config struct.
func (cfg *Config) ValidateWithContext(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, cfg,
		validation.Field(&cfg.Sendgrid, validation.When(strings.EqualFold(strings.TrimSpace(cfg.Provider), ProviderSendgrid), validation.Required)),
	)
}

// ProvideEmailer provides an emailer.
func (cfg *Config) ProvideEmailer(logger logging.Logger, tracerProvider tracing.TracerProvider, client *http.Client) (email.Emailer, error) {
	switch strings.ToLower(strings.TrimSpace(cfg.Provider)) {
	case ProviderSendgrid:
		return sendgrid.NewSendGridEmailer(cfg.Sendgrid, logger, tracerProvider, client)
	case ProviderMailgun:
		return mailgun.NewMailgunEmailer(cfg.Mailgun, logger, tracerProvider, client)
	case ProviderMailjet:
		return mailjet.NewMailjetEmailer(cfg.Mailjet, logger, tracerProvider, client)
	default:
		logger.Debug("providing noop emailer")
		return email.NewNoopEmailer()
	}
}
