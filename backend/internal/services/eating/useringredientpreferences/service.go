package useringredientpreferences

import (
	"fmt"
	"net/http"

	"github.com/dinnerdonebetter/backend/internal/authentication"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	"github.com/dinnerdonebetter/backend/internal/messagequeue"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/routing"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

const (
	serviceName string = "valid_preparations_service"
)

var _ types.UserIngredientPreferenceDataService = (*service)(nil)

type (
	// service handles user ingredient preferences.
	service struct {
		logger                              logging.Logger
		userIngredientPreferenceDataManager types.UserIngredientPreferenceDataManager
		userIngredientPreferenceIDFetcher   func(*http.Request) string
		sessionContextDataFetcher           func(*http.Request) (*types.SessionContextData, error)
		dataChangesPublisher                messagequeue.Publisher
		encoderDecoder                      encoding.ServerEncoderDecoder
		tracer                              tracing.Tracer
	}
)

// ProvideService builds a new UserIngredientPreferencesService.
func ProvideService(
	logger logging.Logger,
	userIngredientPreferenceDataManager types.UserIngredientPreferenceDataManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
	queueConfig *msgconfig.QueuesConfig,
) (types.UserIngredientPreferenceDataService, error) {
	if queueConfig == nil {
		return nil, fmt.Errorf("nil queue config provided")
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(queueConfig.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up %s data changes publisher: %w", serviceName, err)
	}

	svc := &service{
		logger:                              logging.EnsureLogger(logger).WithName(serviceName),
		userIngredientPreferenceIDFetcher:   routeParamManager.BuildRouteParamStringIDFetcher(UserIngredientPreferenceIDURIParamKey),
		sessionContextDataFetcher:           authentication.FetchContextFromRequest,
		userIngredientPreferenceDataManager: userIngredientPreferenceDataManager,
		dataChangesPublisher:                dataChangesPublisher,
		encoderDecoder:                      encoder,
		tracer:                              tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer(serviceName)),
	}

	return svc, nil
}