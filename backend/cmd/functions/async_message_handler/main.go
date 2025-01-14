package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	analyticscfg "github.com/dinnerdonebetter/backend/internal/analytics/config"
	"github.com/dinnerdonebetter/backend/internal/config"
	"github.com/dinnerdonebetter/backend/internal/database/postgres"
	"github.com/dinnerdonebetter/backend/internal/email"
	emailcfg "github.com/dinnerdonebetter/backend/internal/email/config"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/routing/chi"
	"github.com/dinnerdonebetter/backend/internal/search/text/indexing"
	"github.com/dinnerdonebetter/backend/internal/uploads/objectstorage"
	"github.com/dinnerdonebetter/backend/pkg/types"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	_ "go.uber.org/automaxprocs"
)

var (
	errRequiredDataIsNil = errors.New("required data is nil")

	nonWebhookEventTypes = []types.ServiceEventType{
		types.UserSignedUpServiceEventType,
		types.UserArchivedServiceEventType,
		types.TwoFactorSecretVerifiedServiceEventType,
		types.TwoFactorDeactivatedServiceEventType,
		types.TwoFactorSecretChangedServiceEventType,
		types.PasswordResetTokenCreatedEventType,
		types.PasswordResetTokenRedeemedEventType,
		types.PasswordChangedEventType,
		types.EmailAddressChangedEventType,
		types.UsernameChangedEventType,
		types.UserDetailsChangedEventType,
		types.UsernameReminderRequestedEventType,
		types.UserLoggedInServiceEventType,
		types.UserLoggedOutServiceEventType,
		types.UserChangedActiveHouseholdServiceEventType,
		types.UserEmailAddressVerifiedEventType,
		types.UserEmailAddressVerificationEmailRequestedEventType,
		types.HouseholdMemberRemovedServiceEventType,
		types.HouseholdMembershipPermissionsUpdatedServiceEventType,
		types.HouseholdOwnershipTransferredServiceEventType,
		types.OAuth2ClientCreatedServiceEventType,
		types.OAuth2ClientArchivedServiceEventType,
	}
)

func main() {
	ctx := context.Background()

	if config.ShouldCeaseOperation() {
		slog.Info("CEASE_OPERATION is set to true, exiting")
		os.Exit(0)
	}

	cfg, err := config.LoadConfigFromEnvironment[config.AsyncMessageHandlerConfig]()
	if err != nil {
		log.Fatalf("error getting config: %v", err)
	}
	cfg.Database.RunMigrations = false

	logger := cfg.Observability.Logging.ProvideLogger()

	tracerProvider, err := cfg.Observability.Tracing.ProvideTracerProvider(ctx, logger)
	if err != nil {
		logger.Error("initializing tracer", err)
	}
	otel.SetTracerProvider(tracerProvider)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	stopChan := make(chan bool)
	errorsChan := make(chan error)

	if err = doTheThing(ctx, logger, tracerProvider, cfg, stopChan, errorsChan); err != nil {
		log.Fatal(err)
	}

	// os.Interrupt
	<-signalChan

	go func() {
		// os.Kill
		<-signalChan
		stopChan <- true
	}()
}

func doTheThing(
	ctx context.Context,
	logger logging.Logger,
	tracerProvider tracing.TracerProvider,
	cfg *config.AsyncMessageHandlerConfig,
	stopChan chan bool,
	errorsChan chan error,
) error {
	tracer := tracing.NewTracer(tracing.EnsureTracerProvider(tracerProvider).Tracer("async_message_handler"))

	ctx, span := tracer.StartSpan(ctx)
	defer span.End()

	// connect to database

	dbConnectionContext, cancel := context.WithTimeout(ctx, 15*time.Second)
	dataManager, err := postgres.ProvideDatabaseClient(dbConnectionContext, logger, tracerProvider, &cfg.Database)
	if err != nil {
		cancel()
		return observability.PrepareAndLogError(err, logger, span, "establishing database connection")
	}

	cancel()
	defer dataManager.Close()

	// setup baseline messaging providers

	consumerProvider, err := msgconfig.ProvideConsumerProvider(ctx, logger, &cfg.Events)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring queue manager")
	}

	publisherProvider, err := msgconfig.ProvidePublisherProvider(ctx, logger, tracerProvider, &cfg.Events)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring queue manager")
	}

	defer publisherProvider.Close()

	// set up myriad publishers

	analyticsEventReporter, err := analyticscfg.ProvideEventReporter(&cfg.Analytics, logger, tracerProvider)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "setting up customer data collector")
	}

	defer analyticsEventReporter.Close()

	outboundEmailsPublisher, err := publisherProvider.ProvidePublisher(cfg.Queues.OutboundEmailsTopicName)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring outbound emails publisher")
	}

	defer outboundEmailsPublisher.Stop()

	searchDataIndexPublisher, err := publisherProvider.ProvidePublisher(cfg.Queues.SearchIndexRequestsTopicName)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring search indexing publisher")
	}

	defer searchDataIndexPublisher.Stop()

	webhookExecutionRequestPublisher, err := publisherProvider.ProvidePublisher(cfg.Queues.WebhookExecutionRequestsTopicName)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring webhook execution requests publisher")
	}

	defer webhookExecutionRequestPublisher.Stop()

	// setup emailer

	emailer, err := emailcfg.ProvideEmailer(&cfg.Email, logger, tracerProvider, otelhttp.DefaultClient)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring outbound emailer")
	}

	// setup uploader

	uploadManager, err := objectstorage.NewUploadManager(ctx, logger, tracerProvider, &cfg.Storage, chi.NewRouteParamManager())
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "creating upload manager")
	}

	// setup message listeners

	dataChangesConsumer, err := consumerProvider.ProvideConsumer(ctx, cfg.Queues.DataChangesTopicName, func(ctx context.Context, rawMsg []byte) error {
		var dataChangeMessage types.DataChangeMessage
		if err = json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&dataChangeMessage); err != nil {
			return fmt.Errorf("decoding JSON body: %w", err)
		}
		return handleDataChangeMessage(ctx, logger, tracer, dataManager, analyticsEventReporter, webhookExecutionRequestPublisher, outboundEmailsPublisher, searchDataIndexPublisher, &dataChangeMessage)
	})
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring data changes consumer")
	}

	outboundEmailsConsumer, err := consumerProvider.ProvideConsumer(ctx, cfg.Queues.OutboundEmailsTopicName, func(ctx context.Context, rawMsg []byte) error {
		var emailMessage email.DeliveryRequest
		if err = json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&emailMessage); err != nil {
			return fmt.Errorf("decoding JSON body: %w", err)
		}

		return handleEmailRequest(ctx, logger, tracer, dataManager, emailer, analyticsEventReporter, &emailMessage)
	})
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring outbound emails consumer")
	}

	searchIndexRequestsConsumer, err := consumerProvider.ProvideConsumer(ctx, cfg.Queues.SearchIndexRequestsTopicName, func(ctx context.Context, rawMsg []byte) error {
		var searchIndexRequest indexing.IndexRequest
		if err = json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&searchIndexRequest); err != nil {
			return fmt.Errorf("decoding JSON body: %w", err)
		}

		// we don't want to retry indexing perpetually in the event of a fundamental error, so we just log it and move on
		return indexing.HandleIndexRequest(ctx, logger, tracerProvider, &cfg.Search, dataManager, &searchIndexRequest)
	})
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring search index requests consumer")
	}

	userDataAggregationConsumer, err := consumerProvider.ProvideConsumer(ctx, cfg.Queues.UserDataAggregationTopicName, func(ctx context.Context, rawMsg []byte) error {
		var userDataAggregationRequest types.UserDataAggregationRequest
		if err = json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&userDataAggregationRequest); err != nil {
			return fmt.Errorf("decoding JSON body: %w", err)
		}

		return handleUserDataRequest(ctx, logger, tracer, uploadManager, dataManager, &userDataAggregationRequest)
	})
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring user data aggregation consumer")
	}

	webhookExecutionRequestsConsumer, err := consumerProvider.ProvideConsumer(ctx, cfg.Queues.WebhookExecutionRequestsTopicName, func(ctx context.Context, rawMsg []byte) error {
		var webhookExecutionRequest types.WebhookExecutionRequest
		if err = json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&webhookExecutionRequest); err != nil {
			return fmt.Errorf("decoding JSON body: %w", err)
		}

		return handleWebhookExecutionRequest(ctx, logger, tracer, dataManager, &webhookExecutionRequest)
	})
	if err != nil {
		return observability.PrepareAndLogError(err, logger, span, "configuring webhook execution requests consumer")
	}

	go dataChangesConsumer.Consume(
		stopChan,
		errorsChan,
	)

	go outboundEmailsConsumer.Consume(
		stopChan,
		errorsChan,
	)

	go searchIndexRequestsConsumer.Consume(
		stopChan,
		errorsChan,
	)

	go userDataAggregationConsumer.Consume(
		stopChan,
		errorsChan,
	)

	go webhookExecutionRequestsConsumer.Consume(
		stopChan,
		errorsChan,
	)

	go func() {
		for e := range errorsChan {
			logger.Error("consuming message", e)
		}
	}()

	return nil
}
