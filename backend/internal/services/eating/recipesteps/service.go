package recipesteps

import (
	"context"
	"errors"
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
	"github.com/dinnerdonebetter/backend/internal/uploads"
	"github.com/dinnerdonebetter/backend/internal/uploads/images"
	"github.com/dinnerdonebetter/backend/internal/uploads/objectstorage"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

const (
	serviceName string = "recipe_steps_service"
)

var _ types.RecipeStepDataService = (*service)(nil)

type (
	// service handles recipe steps.
	service struct {
		logger                    logging.Logger
		dataChangesPublisher      messagequeue.Publisher
		recipeStepDataManager     types.RecipeStepDataManager
		recipeMediaDataManager    types.RecipeMediaDataManager
		uploadManager             uploads.UploadManager
		tracer                    tracing.Tracer
		encoderDecoder            encoding.ServerEncoderDecoder
		imageUploadProcessor      images.MediaUploadProcessor
		sessionContextDataFetcher func(*http.Request) (*types.SessionContextData, error)
		recipeIDFetcher           func(*http.Request) string
		recipeStepIDFetcher       func(*http.Request) string
		cfg                       Config
	}
)

// ProvideService builds a new RecipeStepsService.
func ProvideService(
	ctx context.Context,
	logger logging.Logger,
	cfg *Config,
	recipeStepDataManager types.RecipeStepDataManager,
	recipeMediaDataManager types.RecipeMediaDataManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
	imageUploadProcessor images.MediaUploadProcessor,
	queueConfig *msgconfig.QueuesConfig,
) (types.RecipeStepDataService, error) {
	if cfg == nil {
		return nil, errors.New("nil config provided to recipe steps service")
	}

	if queueConfig == nil {
		return nil, fmt.Errorf("nil queue config provided")
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(queueConfig.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up %s data changes publisher: %w", serviceName, err)
	}

	uploader, err := objectstorage.NewUploadManager(ctx, logger, tracerProvider, &cfg.Uploads.Storage, routeParamManager)
	if err != nil {
		return nil, fmt.Errorf("initializing recipe steps service upload manager: %w", err)
	}

	svc := &service{
		logger:                    logging.EnsureLogger(logger).WithName(serviceName),
		recipeIDFetcher:           routeParamManager.BuildRouteParamStringIDFetcher(recipesservice.RecipeIDURIParamKey),
		recipeStepIDFetcher:       routeParamManager.BuildRouteParamStringIDFetcher(RecipeStepIDURIParamKey),
		sessionContextDataFetcher: authentication.FetchContextFromRequest,
		recipeStepDataManager:     recipeStepDataManager,
		recipeMediaDataManager:    recipeMediaDataManager,
		dataChangesPublisher:      dataChangesPublisher,
		encoderDecoder:            encoder,
		cfg:                       *cfg,
		imageUploadProcessor:      imageUploadProcessor,
		uploadManager:             uploader,
		tracer:                    tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer(serviceName)),
	}

	return svc, nil
}