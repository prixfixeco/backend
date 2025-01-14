package usernotifications

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
	serviceName string = "user_notifications_service"
)

var _ types.UserNotificationDataService = (*service)(nil)

type (
	// service handles user notifications.
	service struct {
		logger                      logging.Logger
		dataChangesPublisher        messagequeue.Publisher
		tracer                      tracing.Tracer
		encoderDecoder              encoding.ServerEncoderDecoder
		userNotificationDataManager types.UserNotificationDataManager
		sessionContextDataFetcher   func(*http.Request) (*types.SessionContextData, error)
		userNotificationIDFetcher   func(*http.Request) string
	}
)

// ProvideService builds a new UserNotificationsService.
func ProvideService(
	logger logging.Logger,
	encoder encoding.ServerEncoderDecoder,
	routeParamManager routing.RouteParamManager,
	publisherProvider messagequeue.PublisherProvider,
	tracerProvider tracing.TracerProvider,
	userNotificationDataManager types.UserNotificationDataManager,
	queueConfig *msgconfig.QueuesConfig,
) (types.UserNotificationDataService, error) {
	if queueConfig == nil {
		return nil, fmt.Errorf("nil queue config provided")
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(queueConfig.DataChangesTopicName)
	if err != nil {
		return nil, fmt.Errorf("setting up %s data changes publisher: %w", serviceName, err)
	}

	svc := &service{
		logger:                      logging.EnsureLogger(logger).WithName(serviceName),
		userNotificationIDFetcher:   routeParamManager.BuildRouteParamStringIDFetcher(UserNotificationIDURIParamKey),
		sessionContextDataFetcher:   authentication.FetchContextFromRequest,
		dataChangesPublisher:        dataChangesPublisher,
		encoderDecoder:              encoder,
		userNotificationDataManager: userNotificationDataManager,
		tracer:                      tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer(serviceName)),
	}

	return svc, nil
}
