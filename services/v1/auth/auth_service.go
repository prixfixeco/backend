package auth

import (
	"context"
	"net/http"

	"gitlab.com/prixfixe/prixfixe/internal/v1/auth"
	"gitlab.com/prixfixe/prixfixe/internal/v1/config"
	"gitlab.com/prixfixe/prixfixe/internal/v1/encoding"
	models "gitlab.com/prixfixe/prixfixe/models/v1"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/securecookie"
	"gitlab.com/verygoodsoftwarenotvirus/logging/v1"
)

const (
	serviceName = "auth_service"
)

type (
	// OAuth2ClientValidator is a stand-in interface, where we needed to abstract
	// a regular structure with an interface for testing purposes.
	OAuth2ClientValidator interface {
		ExtractOAuth2ClientFromRequest(ctx context.Context, req *http.Request) (*models.OAuth2Client, error)
	}

	// cookieEncoderDecoder is a stand-in interface for gorilla/securecookie
	cookieEncoderDecoder interface {
		Encode(name string, value interface{}) (string, error)
		Decode(name, value string, dst interface{}) error
	}

	// Service handles authentication service-wide
	Service struct {
		config               config.AuthSettings
		logger               logging.Logger
		authenticator        auth.Authenticator
		userDB               models.UserDataManager
		oauth2ClientsService OAuth2ClientValidator
		encoderDecoder       encoding.EncoderDecoder
		cookieManager        cookieEncoderDecoder
		sessionManager       *scs.SessionManager
	}
)

// ProvideAuthService builds a new AuthService.
func ProvideAuthService(
	logger logging.Logger,
	cfg config.AuthSettings,
	authenticator auth.Authenticator,
	database models.UserDataManager,
	oauth2ClientsService OAuth2ClientValidator,
	sessionManager *scs.SessionManager,
	encoder encoding.EncoderDecoder,
) (*Service, error) {
	svc := &Service{
		logger:               logger.WithName(serviceName),
		encoderDecoder:       encoder,
		config:               cfg,
		userDB:               database,
		oauth2ClientsService: oauth2ClientsService,
		authenticator:        authenticator,
		sessionManager:       sessionManager,
		cookieManager: securecookie.New(
			securecookie.GenerateRandomKey(64),
			[]byte(cfg.CookieSecret),
		),
	}
	svc.sessionManager.Lifetime = cfg.CookieLifetime

	return svc, nil
}
