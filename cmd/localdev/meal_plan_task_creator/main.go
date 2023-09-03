package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	analyticsconfig "github.com/dinnerdonebetter/backend/internal/analytics/config"
	"github.com/dinnerdonebetter/backend/internal/config"
	"github.com/dinnerdonebetter/backend/internal/database/postgres"
	"github.com/dinnerdonebetter/backend/internal/features/recipeanalysis"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/messagequeue/redis"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	logcfg "github.com/dinnerdonebetter/backend/internal/observability/logging/config"
	"github.com/dinnerdonebetter/backend/internal/workers"

	_ "go.uber.org/automaxprocs"
)

const (
	dataChangesTopicName      = "data_changes"
	mealPlanTaskCreationTopic = "meal_plan_task_creation"

	configFilepathEnvVar = "CONFIGURATION_FILEPATH"
)

func main() {
	ctx := context.Background()

	logger := (&logcfg.Config{Level: logging.DebugLevel, Provider: logcfg.ProviderSlog}).ProvideLogger()

	logger.Info("starting meal plan task worker...")

	// find and validate our configuration filepath.
	configFilepath := os.Getenv(configFilepathEnvVar)
	if configFilepath == "" {
		log.Fatal("no config provided")
	}

	configBytes, err := os.ReadFile(configFilepath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg *config.InstanceConfig
	if err = json.NewDecoder(bytes.NewReader(configBytes)).Decode(&cfg); err != nil || cfg == nil {
		log.Fatal(err)
	}

	cfg.Observability.Tracing.Otel.ServiceName = "meal_plan_task_creation_workers"

	tracerProvider, initializeTracerErr := cfg.Observability.Tracing.ProvideTracerProvider(ctx, logger)
	if initializeTracerErr != nil {
		logger.Error(initializeTracerErr, "initializing tracer")
	}

	cfg.Database.RunMigrations = false

	cdp, err := analyticsconfig.ProvideEventReporter(&cfg.Analytics, logger, tracerProvider)
	if err != nil {
		log.Fatal(err)
	}

	dataManager, err := postgres.ProvideDatabaseClient(ctx, logger, &cfg.Database, tracerProvider)
	if err != nil {
		log.Fatal(err)
	}

	consumerProvider := redis.ProvideRedisConsumerProvider(logger, tracerProvider, cfg.Events.Consumers.Redis)

	publisherProvider, err := msgconfig.ProvidePublisherProvider(ctx, logger, tracerProvider, &cfg.Events)
	if err != nil {
		log.Fatal(err)
	}

	dataChangesPublisher, err := publisherProvider.ProvidePublisher(dataChangesTopicName)
	if err != nil {
		log.Fatal(err)
	}

	mealPlanTaskCreationEnsurerWorker := workers.ProvideMealPlanTaskCreationEnsurerWorker(
		logger,
		dataManager,
		recipeanalysis.NewRecipeAnalyzer(logger, tracerProvider),
		dataChangesPublisher,
		cdp,
		tracerProvider,
	)

	mealPlanTaskCreationConsumer, err := consumerProvider.ProvideConsumer(ctx, mealPlanTaskCreationTopic, mealPlanTaskCreationEnsurerWorker.CreateMealPlanTasksForFinalizedMealPlans)
	if err != nil {
		log.Fatal(err)
	}

	go mealPlanTaskCreationConsumer.Consume(nil, nil)

	logger.Info("working...")

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
}
