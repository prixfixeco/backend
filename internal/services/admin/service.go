package admin

import (
	"net/http"

	"github.com/prixfixeco/backend/internal/authentication"
	"github.com/prixfixeco/backend/internal/encoding"
	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/internal/routing"
	authservice "github.com/prixfixeco/backend/internal/services/authentication"
	"github.com/prixfixeco/backend/pkg/types"

	"github.com/alexedwards/scs/v2"
)

const (
	serviceName = "auth_service"
)

type (
	// service handles passwords service-wide.
	service struct {
		config                    *authservice.Config
		logger                    logging.Logger
		authenticator             authentication.Authenticator
		userDB                    types.AdminUserDataManager
		encoderDecoder            encoding.ServerEncoderDecoder
		sessionManager            *scs.SessionManager
		sessionContextDataFetcher func(*http.Request) (*types.SessionContextData, error)
		userIDFetcher             func(*http.Request) string
		tracer                    tracing.Tracer
	}
)

// ProvideService builds a new AuthService.
func ProvideService(
	logger logging.Logger,
	cfg *authservice.Config,
	authenticator authentication.Authenticator,
	userDataManager types.AdminUserDataManager,
	sessionManager *scs.SessionManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	tracerProvider tracing.TracerProvider,
) types.AdminService {
	svc := &service{
		logger:                    logging.EnsureLogger(logger).WithName(serviceName),
		encoderDecoder:            encoder,
		config:                    cfg,
		userDB:                    userDataManager,
		authenticator:             authenticator,
		sessionManager:            sessionManager,
		sessionContextDataFetcher: authservice.FetchContextFromRequest,
		userIDFetcher:             routeParamManager.BuildRouteParamStringIDFetcher(UserIDURIParamKey),
		tracer:                    tracing.NewTracer(tracerProvider.Tracer(serviceName)),
	}
	svc.sessionManager.Lifetime = cfg.Cookies.Lifetime

	return svc
}
