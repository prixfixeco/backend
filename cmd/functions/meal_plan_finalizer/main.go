package mealplanfinalizerfunction

import (
	"context"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/prixfixeco/api_server/internal/config"
	"github.com/prixfixeco/api_server/internal/database"
	"github.com/prixfixeco/api_server/internal/database/queriers/postgres"
	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/logging"
	"github.com/prixfixeco/api_server/internal/observability/logging/zerolog"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
)

const (
	dataChangesTopicName = "data_changes"
)

func finalizeMealPlans(ctx context.Context, logger logging.Logger, tracer tracing.Tracer, dataManager database.DataManager) error {
	_, span := tracer.StartSpan(ctx)
	defer span.End()

	logger.Debug("finalize meal plan chore invoked")

	mealPlans, fetchMealPlansErr := dataManager.GetUnfinalizedMealPlansWithExpiredVotingPeriods(ctx)
	if fetchMealPlansErr != nil {
		logger.Fatal(fetchMealPlansErr)
	}

	for _, mealPlan := range mealPlans {
		changed, err := dataManager.FinalizeMealPlanWithExpiredVotingPeriod(ctx, mealPlan.ID, mealPlan.BelongsToHousehold)
		if err != nil {
			return observability.PrepareError(err, logger, span, "finalizing meal plan")
		}

		if !changed {
			logger.Debug("meal plan was not changed")
		}
	}

	return nil
}

// FinalizeMealPlans is our cloud function entrypoint.
func FinalizeMealPlans() {
	ctx := context.Background()
	logger := zerolog.NewZerologLogger()

	cfg, err := config.GetConfigFromGoogleCloudSecretManager(ctx)
	if err != nil {
		logger.Fatal(err)
	}
	cfg.Database.RunMigrations = false

	tracerProvider := trace.NewNoopTracerProvider()
	otel.SetTracerProvider(tracerProvider)
	tracer := tracing.NewTracer(tracerProvider.Tracer("meal_plan_finalizer"))

	dataManager, err := postgres.ProvideDatabaseClient(ctx, logger, &cfg.Database, tracerProvider)
	if err != nil {
		logger.Fatal(err)
	}

	// msgconfig.ProvidePublisherProvider
	// publisherProvider.ProviderPublisher
	// emailconfig.ProvideEmailer
	// customerdataconfig.ProvideCollector

	if err = finalizeMealPlans(ctx, logger, tracer, dataManager); err != nil {
		logger.Fatal(err)
	}

	logger.Info("all done")
}
