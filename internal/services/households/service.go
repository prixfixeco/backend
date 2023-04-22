package households

import (
	"fmt"
	"net/http"

	"github.com/prixfixeco/backend/internal/encoding"
	"github.com/prixfixeco/backend/internal/messagequeue"
	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/metrics"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/internal/routing"
	authservice "github.com/prixfixeco/backend/internal/services/authentication"
	"github.com/prixfixeco/backend/pkg/types"
)

const (
	counterName        metrics.CounterName = "households"
	counterDescription string              = "the number of households managed by the households service"
	serviceName        string              = "households_service"
)

var _ types.HouseholdDataService = (*service)(nil)

type (
	// service handles to-do list households.
	service struct {
		logger                         logging.Logger
		householdDataManager           types.HouseholdDataManager
		householdInvitationDataManager types.HouseholdInvitationDataManager
		householdMembershipDataManager types.HouseholdUserMembershipDataManager
		tracer                         tracing.Tracer
		encoderDecoder                 encoding.ServerEncoderDecoder
		dataChangesPublisher           messagequeue.Publisher
		sessionContextDataFetcher      func(*http.Request) (*types.SessionContextData, error)
		userIDFetcher                  func(*http.Request) string
		householdIDFetcher             func(*http.Request) string
	}
)

// ProvideService builds a new HouseholdsService.
func ProvideService(
	logger logging.Logger,
	cfg Config,
	householdDataManager types.HouseholdDataManager,
	householdInvitationDataManager types.HouseholdInvitationDataManager,
	householdMembershipDataManager types.HouseholdUserMembershipDataManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
) (types.HouseholdDataService, error) {
	dataChangesPublisher, err := publisherProvider.ProvidePublisher(cfg.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up household service data changes publisher: %w", err)
	}

	s := &service{
		logger:                         logging.EnsureLogger(logger).WithName(serviceName),
		householdIDFetcher:             routeParamManager.BuildRouteParamStringIDFetcher(HouseholdIDURIParamKey),
		userIDFetcher:                  routeParamManager.BuildRouteParamStringIDFetcher(UserIDURIParamKey),
		sessionContextDataFetcher:      authservice.FetchContextFromRequest,
		householdDataManager:           householdDataManager,
		householdInvitationDataManager: householdInvitationDataManager,
		householdMembershipDataManager: householdMembershipDataManager,
		encoderDecoder:                 encoder,
		dataChangesPublisher:           dataChangesPublisher,
		tracer:                         tracing.NewTracer(tracerProvider.Tracer(serviceName)),
	}

	return s, nil
}
