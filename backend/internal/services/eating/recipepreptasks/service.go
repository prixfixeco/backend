package recipepreptasks

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
	recipesservice "github.com/dinnerdonebetter/backend/internal/services/eating/recipes"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

const (
	serviceName string = "recipe_prep_tasks_service"
)

var _ types.RecipePrepTaskDataService = (*service)(nil)

type (
	// service handles recipe prep tasks.
	service struct {
		logger                    logging.Logger
		recipePrepTaskDataManager types.RecipePrepTaskDataManager
		recipeIDFetcher           func(*http.Request) string
		recipePrepTaskIDFetcher   func(*http.Request) string
		sessionContextDataFetcher func(*http.Request) (*types.SessionContextData, error)
		dataChangesPublisher      messagequeue.Publisher
		encoderDecoder            encoding.ServerEncoderDecoder
		tracer                    tracing.Tracer
	}
)

// ProvideService builds a new RecipePrepTasksService.
func ProvideService(
	logger logging.Logger,
	recipePrepTaskDataManager types.RecipePrepTaskDataManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
	queueConfig *msgconfig.QueuesConfig,
) (types.RecipePrepTaskDataService, error) {
	if queueConfig == nil {
		return nil, fmt.Errorf("nil queue config provided")
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(queueConfig.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up %s data changes publisher: %w", serviceName, err)
	}

	svc := &service{
		logger:                    logging.EnsureLogger(logger).WithName(serviceName),
		recipeIDFetcher:           routeParamManager.BuildRouteParamStringIDFetcher(recipesservice.RecipeIDURIParamKey),
		recipePrepTaskIDFetcher:   routeParamManager.BuildRouteParamStringIDFetcher(RecipePrepTaskIDURIParamKey),
		sessionContextDataFetcher: authentication.FetchContextFromRequest,
		recipePrepTaskDataManager: recipePrepTaskDataManager,
		dataChangesPublisher:      dataChangesPublisher,
		encoderDecoder:            encoder,
		tracer:                    tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer(serviceName)),
	}

	return svc, nil
}