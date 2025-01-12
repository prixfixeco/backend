package validmeasurementunitconversions

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
	serviceName string = "valid_measurement_conversion_service"
)

var _ types.ValidMeasurementUnitConversionDataService = (*service)(nil)

type (
	// service handles valid measurement conversions.
	service struct {
		logger                                    logging.Logger
		validMeasurementUnitConversionDataManager types.ValidMeasurementUnitConversionDataManager
		validMeasurementUnitConversionIDFetcher   func(*http.Request) string
		validMeasurementUnitIDFetcher             func(*http.Request) string
		sessionContextDataFetcher                 func(*http.Request) (*types.SessionContextData, error)
		dataChangesPublisher                      messagequeue.Publisher
		encoderDecoder                            encoding.ServerEncoderDecoder
		tracer                                    tracing.Tracer
	}
)

// ProvideService builds a new ValidMeasurementUnitConversionsService.
func ProvideService(
	logger logging.Logger,
	validMeasurementUnitConversionDataManager types.ValidMeasurementUnitConversionDataManager,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
	queueConfig *msgconfig.QueuesConfig,
) (types.ValidMeasurementUnitConversionDataService, error) {
	if queueConfig == nil {
		return nil, fmt.Errorf("nil queue config provided")
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(queueConfig.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up %s data changes publisher: %w", serviceName, err)
	}

	svc := &service{
		logger:                                    logging.EnsureLogger(logger).WithName(serviceName),
		validMeasurementUnitConversionIDFetcher:   routeParamManager.BuildRouteParamStringIDFetcher(ValidMeasurementUnitConversionIDURIParamKey),
		validMeasurementUnitIDFetcher:             routeParamManager.BuildRouteParamStringIDFetcher(ValidMeasurementUnitIDURIParamKey),
		sessionContextDataFetcher:                 authentication.FetchContextFromRequest,
		validMeasurementUnitConversionDataManager: validMeasurementUnitConversionDataManager,
		dataChangesPublisher:                      dataChangesPublisher,
		encoderDecoder:                            encoder,
		tracer:                                    tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer(serviceName)),
	}

	return svc, nil
}