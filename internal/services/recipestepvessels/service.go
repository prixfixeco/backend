package recipestepvessels

import (
	"fmt"
	"net/http"

	"github.com/prixfixeco/backend/internal/encoding"
	"github.com/prixfixeco/backend/internal/messagequeue"
	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/internal/routing"
	authservice "github.com/prixfixeco/backend/internal/services/authentication"
	recipesservice "github.com/prixfixeco/backend/internal/services/recipes"
	recipestepsservice "github.com/prixfixeco/backend/internal/services/recipesteps"
	"github.com/prixfixeco/backend/pkg/types"
)

const (
	serviceName string = "recipe_step_vessels_service"
)

var _ types.RecipeStepVesselDataService = (*service)(nil)

type (
	// service handles recipe step vessels.
	service struct {
		logger                      logging.Logger
		recipeStepVesselDataManager types.RecipeStepVesselDataManager
		recipeIDFetcher             func(*http.Request) string
		recipeStepIDFetcher         func(*http.Request) string
		recipeStepVesselIDFetcher   func(*http.Request) string
		sessionContextDataFetcher   func(*http.Request) (*types.SessionContextData, error)
		dataChangesPublisher        messagequeue.Publisher
		encoderDecoder              encoding.ServerEncoderDecoder
		tracer                      tracing.Tracer
	}
)

// ProvideService builds a new RecipeStepVesselsService.
func ProvideService(
	logger logging.Logger,
	cfg *Config,
	recipeStepVesselDataManager types.RecipeStepVesselDataManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
) (types.RecipeStepVesselDataService, error) {
	dataChangesPublisher, err := publisherProvider.ProvidePublisher(cfg.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up recipe step vessels service data changes publisher: %w", err)
	}

	svc := &service{
		logger:                      logging.EnsureLogger(logger).WithName(serviceName),
		recipeIDFetcher:             routeParamManager.BuildRouteParamStringIDFetcher(recipesservice.RecipeIDURIParamKey),
		recipeStepIDFetcher:         routeParamManager.BuildRouteParamStringIDFetcher(recipestepsservice.RecipeStepIDURIParamKey),
		recipeStepVesselIDFetcher:   routeParamManager.BuildRouteParamStringIDFetcher(RecipeStepVesselIDURIParamKey),
		sessionContextDataFetcher:   authservice.FetchContextFromRequest,
		recipeStepVesselDataManager: recipeStepVesselDataManager,
		dataChangesPublisher:        dataChangesPublisher,
		encoderDecoder:              encoder,
		tracer:                      tracing.NewTracer(tracerProvider.Tracer(serviceName)),
	}

	return svc, nil
}
