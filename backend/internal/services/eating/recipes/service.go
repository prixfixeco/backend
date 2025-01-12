package recipes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dinnerdonebetter/backend/internal/authentication"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	"github.com/dinnerdonebetter/backend/internal/features/recipeanalysis"
	"github.com/dinnerdonebetter/backend/internal/messagequeue"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/routing"
	"github.com/dinnerdonebetter/backend/internal/search/text"
	textsearchcfg "github.com/dinnerdonebetter/backend/internal/search/text/config"
	"github.com/dinnerdonebetter/backend/internal/uploads"
	"github.com/dinnerdonebetter/backend/internal/uploads/images"
	"github.com/dinnerdonebetter/backend/internal/uploads/objectstorage"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

const (
	serviceName string = "recipes_service"
)

var _ types.RecipeDataService = (*service)(nil)

type (
	// service handles recipes.
	service struct {
		logger                    logging.Logger
		tracer                    tracing.Tracer
		recipeDataManager         types.RecipeDataManager
		recipeMediaDataManager    types.RecipeMediaDataManager
		recipeAnalyzer            recipeanalysis.RecipeAnalyzer
		imageUploadProcessor      images.MediaUploadProcessor
		encoderDecoder            encoding.ServerEncoderDecoder
		dataChangesPublisher      messagequeue.Publisher
		searchIndex               textsearch.IndexSearcher[types.RecipeSearchSubset]
		uploadManager             uploads.UploadManager
		sessionContextDataFetcher func(*http.Request) (*types.SessionContextData, error)
		recipeIDFetcher           func(*http.Request) string
		cfg                       *Config
	}
)

var errInvalidConfig = errors.New("config cannot be nil")

// ProvideService builds a new RecipesService.
func ProvideService(
	ctx context.Context,
	logger logging.Logger,
	cfg *Config,
	searchConfig *textsearchcfg.Config,
	recipeDataManager types.RecipeDataManager,
	recipeMediaDataManager types.RecipeMediaDataManager,
	recipeGrapher recipeanalysis.RecipeAnalyzer,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	imageUploadProcessor images.MediaUploadProcessor,
	tracerProvider tracing.TracerProvider,
	queueConfig *msgconfig.QueuesConfig,
) (types.RecipeDataService, error) {
	if queueConfig == nil {
		return nil, fmt.Errorf("nil queue config provided")
	}

	if cfg == nil {
		return nil, errInvalidConfig
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(queueConfig.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up %s data changes publisher: %w", serviceName, err)
	}

	uploader, err := objectstorage.NewUploadManager(ctx, logger, tracerProvider, &cfg.Uploads.Storage, routeParamManager)
	if err != nil {
		return nil, fmt.Errorf("initializing %s upload manager: %w", serviceName, err)
	}

	searchIndex, err := textsearchcfg.ProvideIndex[types.RecipeSearchSubset](ctx, logger, tracerProvider, searchConfig, textsearch.IndexTypeRecipes)
	if err != nil {
		return nil, observability.PrepareError(err, nil, "initializing recipe index manager")
	}

	svc := &service{
		logger:                    logging.EnsureLogger(logger).WithName(serviceName),
		recipeIDFetcher:           routeParamManager.BuildRouteParamStringIDFetcher(RecipeIDURIParamKey),
		sessionContextDataFetcher: authentication.FetchContextFromRequest,
		recipeDataManager:         recipeDataManager,
		cfg:                       cfg,
		recipeMediaDataManager:    recipeMediaDataManager,
		dataChangesPublisher:      dataChangesPublisher,
		encoderDecoder:            encoder,
		recipeAnalyzer:            recipeGrapher,
		uploadManager:             uploader,
		imageUploadProcessor:      imageUploadProcessor,
		tracer:                    tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer(serviceName)),
		searchIndex:               searchIndex,
	}

	return svc, nil
}