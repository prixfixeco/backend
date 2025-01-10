// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package build

import (
	"context"

	"github.com/dinnerdonebetter/backend/internal/analytics/config"
	"github.com/dinnerdonebetter/backend/internal/authentication"
	"github.com/dinnerdonebetter/backend/internal/config"
	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/database/postgres"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	"github.com/dinnerdonebetter/backend/internal/featureflags/config"
	"github.com/dinnerdonebetter/backend/internal/features/recipeanalysis"
	"github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/observability/logging/config"
	"github.com/dinnerdonebetter/backend/internal/observability/metrics/config"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing/config"
	"github.com/dinnerdonebetter/backend/internal/pkg/random"
	"github.com/dinnerdonebetter/backend/internal/routing/config"
	"github.com/dinnerdonebetter/backend/internal/server/http"
	"github.com/dinnerdonebetter/backend/internal/services/admin"
	"github.com/dinnerdonebetter/backend/internal/services/auditlogentries"
	authentication2 "github.com/dinnerdonebetter/backend/internal/services/authentication"
	"github.com/dinnerdonebetter/backend/internal/services/dataprivacy"
	"github.com/dinnerdonebetter/backend/internal/services/householdinstrumentownerships"
	"github.com/dinnerdonebetter/backend/internal/services/householdinvitations"
	"github.com/dinnerdonebetter/backend/internal/services/households"
	"github.com/dinnerdonebetter/backend/internal/services/mealplanevents"
	"github.com/dinnerdonebetter/backend/internal/services/mealplangrocerylistitems"
	"github.com/dinnerdonebetter/backend/internal/services/mealplanoptions"
	"github.com/dinnerdonebetter/backend/internal/services/mealplanoptionvotes"
	"github.com/dinnerdonebetter/backend/internal/services/mealplans"
	"github.com/dinnerdonebetter/backend/internal/services/mealplantasks"
	"github.com/dinnerdonebetter/backend/internal/services/meals"
	"github.com/dinnerdonebetter/backend/internal/services/oauth2clients"
	"github.com/dinnerdonebetter/backend/internal/services/recipepreptasks"
	"github.com/dinnerdonebetter/backend/internal/services/reciperatings"
	"github.com/dinnerdonebetter/backend/internal/services/recipes"
	"github.com/dinnerdonebetter/backend/internal/services/recipestepcompletionconditions"
	"github.com/dinnerdonebetter/backend/internal/services/recipestepingredients"
	"github.com/dinnerdonebetter/backend/internal/services/recipestepinstruments"
	"github.com/dinnerdonebetter/backend/internal/services/recipestepproducts"
	"github.com/dinnerdonebetter/backend/internal/services/recipesteps"
	"github.com/dinnerdonebetter/backend/internal/services/recipestepvessels"
	"github.com/dinnerdonebetter/backend/internal/services/servicesettingconfigurations"
	"github.com/dinnerdonebetter/backend/internal/services/servicesettings"
	"github.com/dinnerdonebetter/backend/internal/services/useringredientpreferences"
	"github.com/dinnerdonebetter/backend/internal/services/usernotifications"
	"github.com/dinnerdonebetter/backend/internal/services/users"
	"github.com/dinnerdonebetter/backend/internal/services/validingredientgroups"
	"github.com/dinnerdonebetter/backend/internal/services/validingredientmeasurementunits"
	"github.com/dinnerdonebetter/backend/internal/services/validingredientpreparations"
	"github.com/dinnerdonebetter/backend/internal/services/validingredients"
	"github.com/dinnerdonebetter/backend/internal/services/validingredientstateingredients"
	"github.com/dinnerdonebetter/backend/internal/services/validingredientstates"
	"github.com/dinnerdonebetter/backend/internal/services/validinstruments"
	"github.com/dinnerdonebetter/backend/internal/services/validmeasurementunitconversions"
	"github.com/dinnerdonebetter/backend/internal/services/validmeasurementunits"
	"github.com/dinnerdonebetter/backend/internal/services/validpreparationinstruments"
	"github.com/dinnerdonebetter/backend/internal/services/validpreparations"
	"github.com/dinnerdonebetter/backend/internal/services/validpreparationvessels"
	"github.com/dinnerdonebetter/backend/internal/services/validvessels"
	"github.com/dinnerdonebetter/backend/internal/services/webhooks"
	"github.com/dinnerdonebetter/backend/internal/services/workers"
	"github.com/dinnerdonebetter/backend/internal/uploads/images"
)

// Injectors from build.go:

// Build builds a server.
func Build(ctx context.Context, cfg *config.APIServiceConfig) (http.Server, error) {
	httpConfig := cfg.Server
	observabilityConfig := &cfg.Observability
	loggingcfgConfig := &observabilityConfig.Logging
	logger := loggingcfg.ProvideLogger(loggingcfgConfig)
	tracingcfgConfig := &observabilityConfig.Tracing
	tracerProvider, err := tracingcfg.ProvideTracerProvider(ctx, tracingcfgConfig, logger)
	if err != nil {
		return nil, err
	}
	databasecfgConfig := &cfg.Database
	dataManager, err := postgres.ProvideDatabaseClient(ctx, logger, tracerProvider, databasecfgConfig)
	if err != nil {
		return nil, err
	}
	routingcfgConfig := &cfg.Routing
	metricscfgConfig := &observabilityConfig.Metrics
	provider, err := metricscfg.ProvideMetricsProvider(ctx, logger, metricscfgConfig)
	if err != nil {
		return nil, err
	}
	router, err := routingcfg.ProvideRouter(routingcfgConfig, logger, tracerProvider, provider)
	if err != nil {
		return nil, err
	}
	servicesConfig := &cfg.Services
	authenticationConfig := &servicesConfig.Auth
	authenticator := authentication.ProvideArgon2Authenticator(logger, tracerProvider)
	householdUserMembershipDataManager := database.ProvideHouseholdUserMembershipDataManager(dataManager)
	encodingConfig := cfg.Encoding
	contentType := encoding.ProvideContentType(encodingConfig)
	serverEncoderDecoder := encoding.ProvideServerEncoderDecoder(logger, tracerProvider, contentType)
	msgconfigConfig := &cfg.Events
	publisherProvider, err := msgconfig.ProvidePublisherProvider(ctx, logger, tracerProvider, msgconfigConfig)
	if err != nil {
		return nil, err
	}
	featureflagscfgConfig := &cfg.FeatureFlags
	client := tracing.BuildTracedHTTPClient()
	featureFlagManager, err := featureflagscfg.ProvideFeatureFlagManager(featureflagscfgConfig, logger, tracerProvider, client)
	if err != nil {
		return nil, err
	}
	analyticscfgConfig := &cfg.Analytics
	eventReporter, err := analyticscfg.ProvideEventReporter(analyticscfgConfig, logger, tracerProvider)
	if err != nil {
		return nil, err
	}
	routeParamManager, err := routingcfg.ProvideRouteParamManager(routingcfgConfig, logger, tracerProvider, provider)
	if err != nil {
		return nil, err
	}
	queuesConfig := &cfg.Queues
	authDataService, err := authentication2.ProvideService(ctx, logger, authenticationConfig, authenticator, dataManager, householdUserMembershipDataManager, serverEncoderDecoder, tracerProvider, publisherProvider, featureFlagManager, eventReporter, routeParamManager, provider, queuesConfig)
	if err != nil {
		return nil, err
	}
	userDataManager := database.ProvideUserDataManager(dataManager)
	householdInvitationDataManager := database.ProvideHouseholdInvitationDataManager(dataManager)
	generator := random.NewGenerator(logger, tracerProvider)
	passwordResetTokenDataManager := database.ProvidePasswordResetTokenDataManager(dataManager)
	userDataService, err := users.ProvideUsersService(authenticationConfig, logger, userDataManager, householdInvitationDataManager, householdUserMembershipDataManager, authenticator, serverEncoderDecoder, routeParamManager, tracerProvider, publisherProvider, generator, passwordResetTokenDataManager, featureFlagManager, eventReporter, queuesConfig)
	if err != nil {
		return nil, err
	}
	householdDataManager := database.ProvideHouseholdDataManager(dataManager)
	householdDataService, err := households.ProvideService(logger, householdDataManager, householdUserMembershipDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, generator, queuesConfig)
	if err != nil {
		return nil, err
	}
	householdInvitationDataService, err := householdinvitations.ProvideHouseholdInvitationsService(logger, userDataManager, householdInvitationDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, generator, queuesConfig)
	if err != nil {
		return nil, err
	}
	validinstrumentsConfig := &servicesConfig.ValidInstruments
	textsearchcfgConfig := &cfg.Search
	validInstrumentDataManager := database.ProvideValidInstrumentDataManager(dataManager)
	validInstrumentDataService, err := validinstruments.ProvideService(ctx, logger, validinstrumentsConfig, textsearchcfgConfig, validInstrumentDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validingredientsConfig := &servicesConfig.ValidIngredients
	validIngredientDataManager := database.ProvideValidIngredientDataManager(dataManager)
	validIngredientDataService, err := validingredients.ProvideService(ctx, logger, validingredientsConfig, textsearchcfgConfig, validIngredientDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validIngredientGroupDataManager := database.ProvideValidIngredientGroupDataManager(dataManager)
	validIngredientGroupDataService, err := validingredientgroups.ProvideService(logger, validIngredientGroupDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validpreparationsConfig := &servicesConfig.ValidPreparations
	validPreparationDataManager := database.ProvideValidPreparationDataManager(dataManager)
	validPreparationDataService, err := validpreparations.ProvideService(ctx, logger, validpreparationsConfig, textsearchcfgConfig, validPreparationDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validIngredientPreparationDataManager := database.ProvideValidIngredientPreparationDataManager(dataManager)
	validIngredientPreparationDataService, err := validingredientpreparations.ProvideService(logger, validIngredientPreparationDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealsConfig := &servicesConfig.Meals
	mealDataManager := database.ProvideMealDataManager(dataManager)
	mealDataService, err := meals.ProvideService(ctx, logger, mealsConfig, textsearchcfgConfig, mealDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipesConfig := &servicesConfig.Recipes
	recipeDataManager := database.ProvideRecipeDataManager(dataManager)
	recipeMediaDataManager := database.ProvideRecipeMediaDataManager(dataManager)
	recipeAnalyzer := recipeanalysis.NewRecipeAnalyzer(logger, tracerProvider)
	mediaUploadProcessor := images.NewImageUploadProcessor(logger, tracerProvider)
	recipeDataService, err := recipes.ProvideService(ctx, logger, recipesConfig, textsearchcfgConfig, recipeDataManager, recipeMediaDataManager, recipeAnalyzer, serverEncoderDecoder, routeParamManager, publisherProvider, mediaUploadProcessor, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipestepsConfig := &servicesConfig.RecipeSteps
	recipeStepDataManager := database.ProvideRecipeStepDataManager(dataManager)
	recipeStepDataService, err := recipesteps.ProvideService(ctx, logger, recipestepsConfig, recipeStepDataManager, recipeMediaDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, mediaUploadProcessor, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipeStepProductDataManager := database.ProvideRecipeStepProductDataManager(dataManager)
	recipeStepProductDataService, err := recipestepproducts.ProvideService(logger, recipeStepProductDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipeStepInstrumentDataManager := database.ProvideRecipeStepInstrumentDataManager(dataManager)
	recipeStepInstrumentDataService, err := recipestepinstruments.ProvideService(logger, recipeStepInstrumentDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipeStepIngredientDataManager := database.ProvideRecipeStepIngredientDataManager(dataManager)
	recipeStepIngredientDataService, err := recipestepingredients.ProvideService(logger, recipeStepIngredientDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealPlanDataManager := database.ProvideMealPlanDataManager(dataManager)
	mealPlanDataService, err := mealplans.ProvideService(logger, mealPlanDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealPlanOptionDataManager := database.ProvideMealPlanOptionDataManager(dataManager)
	mealPlanOptionDataService, err := mealplanoptions.ProvideService(logger, mealPlanOptionDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealPlanOptionVoteDataService, err := mealplanoptionvotes.ProvideService(logger, dataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validmeasurementunitsConfig := &servicesConfig.ValidMeasurementUnits
	validMeasurementUnitDataManager := database.ProvideValidMeasurementUnitDataManager(dataManager)
	validMeasurementUnitDataService, err := validmeasurementunits.ProvideService(ctx, logger, validmeasurementunitsConfig, textsearchcfgConfig, validMeasurementUnitDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validingredientstatesConfig := &servicesConfig.ValidIngredientStates
	validIngredientStateDataManager := database.ProvideValidIngredientStateDataManager(dataManager)
	validIngredientStateDataService, err := validingredientstates.ProvideService(ctx, logger, validingredientstatesConfig, textsearchcfgConfig, validIngredientStateDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validPreparationInstrumentDataManager := database.ProvideValidPreparationInstrumentDataManager(dataManager)
	validPreparationInstrumentDataService, err := validpreparationinstruments.ProvideService(logger, validPreparationInstrumentDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validIngredientMeasurementUnitDataManager := database.ProvideValidIngredientMeasurementUnitDataManager(dataManager)
	validIngredientMeasurementUnitDataService, err := validingredientmeasurementunits.ProvideService(logger, validIngredientMeasurementUnitDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealPlanEventDataManager := database.ProvideMealPlanEventDataManager(dataManager)
	mealPlanEventDataService, err := mealplanevents.ProvideService(logger, mealPlanEventDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealPlanTaskDataManager := database.ProvideMealPlanTaskDataManager(dataManager)
	mealPlanTaskDataService, err := mealplantasks.ProvideService(logger, mealPlanTaskDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipePrepTaskDataManager := database.ProvideRecipePrepTaskDataManager(dataManager)
	recipePrepTaskDataService, err := recipepreptasks.ProvideService(logger, recipePrepTaskDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	mealPlanGroceryListItemDataManager := database.ProvideMealPlanGroceryListItemDataManager(dataManager)
	mealPlanGroceryListItemDataService, err := mealplangrocerylistitems.ProvideService(logger, mealPlanGroceryListItemDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validMeasurementUnitConversionDataManager := database.ProvideValidMeasurementUnitConversionDataManager(dataManager)
	validMeasurementUnitConversionDataService, err := validmeasurementunitconversions.ProvideService(logger, validMeasurementUnitConversionDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipeStepCompletionConditionDataManager := database.ProvideRecipeStepCompletionConditionDataManager(dataManager)
	recipeStepCompletionConditionDataService, err := recipestepcompletionconditions.ProvideService(logger, recipeStepCompletionConditionDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validIngredientStateIngredientDataManager := database.ProvideValidIngredientStateIngredientDataManager(dataManager)
	validIngredientStateIngredientDataService, err := validingredientstateingredients.ProvideService(logger, validIngredientStateIngredientDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipeStepVesselDataManager := database.ProvideRecipeStepVesselDataManager(dataManager)
	recipeStepVesselDataService, err := recipestepvessels.ProvideService(logger, recipeStepVesselDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	webhookDataManager := database.ProvideWebhookDataManager(dataManager)
	webhookDataService, err := webhooks.ProvideWebhooksService(logger, webhookDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	adminUserDataManager := database.ProvideAdminUserDataManager(dataManager)
	adminDataService, err := admin.ProvideService(logger, adminUserDataManager, serverEncoderDecoder, tracerProvider, queuesConfig, publisherProvider)
	if err != nil {
		return nil, err
	}
	serviceSettingDataManager := database.ProvideServiceSettingDataManager(dataManager)
	serviceSettingDataService, err := servicesettings.ProvideService(logger, serviceSettingDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	serviceSettingConfigurationDataManager := database.ProvideServiceSettingConfigurationDataManager(dataManager)
	serviceSettingConfigurationDataService, err := servicesettingconfigurations.ProvideService(logger, serviceSettingConfigurationDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	userIngredientPreferenceDataManager := database.ProvideUserIngredientPreferenceDataManager(dataManager)
	userIngredientPreferenceDataService, err := useringredientpreferences.ProvideService(logger, userIngredientPreferenceDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	recipeRatingDataManager := database.ProvideRecipeRatingDataManager(dataManager)
	recipeRatingDataService, err := reciperatings.ProvideService(logger, recipeRatingDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	householdInstrumentOwnershipDataManager := database.ProvideHouseholdInstrumentOwnershipDataManager(dataManager)
	householdInstrumentOwnershipDataService, err := householdinstrumentownerships.ProvideService(logger, householdInstrumentOwnershipDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	oauth2clientsConfig := &servicesConfig.OAuth2Clients
	oAuth2ClientDataManager := database.ProvideOAuth2ClientDataManager(dataManager)
	oAuth2ClientDataService, err := oauth2clients.ProvideOAuth2ClientsService(logger, oauth2clientsConfig, oAuth2ClientDataManager, serverEncoderDecoder, routeParamManager, tracerProvider, generator, publisherProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validvesselsConfig := &servicesConfig.ValidVessels
	validVesselDataManager := database.ProvideValidVesselDataManager(dataManager)
	validVesselDataService, err := validvessels.ProvideService(ctx, logger, validvesselsConfig, textsearchcfgConfig, validVesselDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	validPreparationVesselDataManager := database.ProvideValidPreparationVesselDataManager(dataManager)
	validPreparationVesselDataService, err := validpreparationvessels.ProvideService(logger, validPreparationVesselDataManager, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, queuesConfig)
	if err != nil {
		return nil, err
	}
	workerService, err := workers.ProvideService(logger, dataManager, serverEncoderDecoder, publisherProvider, tracerProvider, recipeAnalyzer, queuesConfig)
	if err != nil {
		return nil, err
	}
	userNotificationDataManager := database.ProvideUserNotificationDataManager(dataManager)
	userNotificationDataService, err := usernotifications.ProvideService(logger, serverEncoderDecoder, routeParamManager, publisherProvider, tracerProvider, userNotificationDataManager, queuesConfig)
	if err != nil {
		return nil, err
	}
	auditLogEntryDataManager := database.ProvideAuditLogEntryDataManager(dataManager)
	auditLogEntryDataService, err := auditlogentries.ProvideService(logger, auditLogEntryDataManager, serverEncoderDecoder, routeParamManager, tracerProvider)
	if err != nil {
		return nil, err
	}
	dataprivacyConfig := &servicesConfig.DataPrivacy
	dataPrivacyDataManager := database.ProvideDataPrivacyDataManager(dataManager)
	dataPrivacyService, err := dataprivacy.ProvideService(ctx, logger, dataprivacyConfig, dataPrivacyDataManager, serverEncoderDecoder, publisherProvider, tracerProvider, routeParamManager, queuesConfig)
	if err != nil {
		return nil, err
	}
	server, err := http.ProvideHTTPServer(ctx, httpConfig, dataManager, logger, router, tracerProvider, authDataService, userDataService, householdDataService, householdInvitationDataService, validInstrumentDataService, validIngredientDataService, validIngredientGroupDataService, validPreparationDataService, validIngredientPreparationDataService, mealDataService, recipeDataService, recipeStepDataService, recipeStepProductDataService, recipeStepInstrumentDataService, recipeStepIngredientDataService, mealPlanDataService, mealPlanOptionDataService, mealPlanOptionVoteDataService, validMeasurementUnitDataService, validIngredientStateDataService, validPreparationInstrumentDataService, validIngredientMeasurementUnitDataService, mealPlanEventDataService, mealPlanTaskDataService, recipePrepTaskDataService, mealPlanGroceryListItemDataService, validMeasurementUnitConversionDataService, recipeStepCompletionConditionDataService, validIngredientStateIngredientDataService, recipeStepVesselDataService, webhookDataService, adminDataService, serviceSettingDataService, serviceSettingConfigurationDataService, userIngredientPreferenceDataService, recipeRatingDataService, householdInstrumentOwnershipDataService, oAuth2ClientDataService, validVesselDataService, validPreparationVesselDataService, workerService, userNotificationDataService, auditLogEntryDataService, dataPrivacyService, provider)
	if err != nil {
		return nil, err
	}
	return server, nil
}
