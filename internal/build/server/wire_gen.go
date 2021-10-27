// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"context"
	"gitlab.com/prixfixe/prixfixe/internal/authentication"
	"gitlab.com/prixfixe/prixfixe/internal/config"
	"gitlab.com/prixfixe/prixfixe/internal/database"
	config2 "gitlab.com/prixfixe/prixfixe/internal/database/config"
	"gitlab.com/prixfixe/prixfixe/internal/encoding"
	config3 "gitlab.com/prixfixe/prixfixe/internal/messagequeue/config"
	"gitlab.com/prixfixe/prixfixe/internal/observability/logging"
	"gitlab.com/prixfixe/prixfixe/internal/observability/metrics"
	"gitlab.com/prixfixe/prixfixe/internal/routing/chi"
	"gitlab.com/prixfixe/prixfixe/internal/search/elasticsearch"
	"gitlab.com/prixfixe/prixfixe/internal/server"
	"gitlab.com/prixfixe/prixfixe/internal/services/admin"
	"gitlab.com/prixfixe/prixfixe/internal/services/apiclients"
	authentication2 "gitlab.com/prixfixe/prixfixe/internal/services/authentication"
	"gitlab.com/prixfixe/prixfixe/internal/services/households"
	"gitlab.com/prixfixe/prixfixe/internal/services/mealplanoptions"
	"gitlab.com/prixfixe/prixfixe/internal/services/mealplanoptionvotes"
	"gitlab.com/prixfixe/prixfixe/internal/services/mealplans"
	"gitlab.com/prixfixe/prixfixe/internal/services/recipes"
	"gitlab.com/prixfixe/prixfixe/internal/services/recipestepingredients"
	"gitlab.com/prixfixe/prixfixe/internal/services/recipestepinstruments"
	"gitlab.com/prixfixe/prixfixe/internal/services/recipestepproducts"
	"gitlab.com/prixfixe/prixfixe/internal/services/recipesteps"
	"gitlab.com/prixfixe/prixfixe/internal/services/users"
	"gitlab.com/prixfixe/prixfixe/internal/services/validingredientpreparations"
	"gitlab.com/prixfixe/prixfixe/internal/services/validingredients"
	"gitlab.com/prixfixe/prixfixe/internal/services/validinstruments"
	"gitlab.com/prixfixe/prixfixe/internal/services/validpreparations"
	"gitlab.com/prixfixe/prixfixe/internal/services/webhooks"
	"gitlab.com/prixfixe/prixfixe/internal/services/websockets"
	"gitlab.com/prixfixe/prixfixe/internal/storage"
	"gitlab.com/prixfixe/prixfixe/internal/uploads"
	"gitlab.com/prixfixe/prixfixe/internal/uploads/images"
)

// Injectors from build.go:

// Build builds a server.
func Build(ctx context.Context, logger logging.Logger, cfg *config.InstanceConfig) (*server.HTTPServer, error) {
	serverConfig := cfg.Server
	observabilityConfig := &cfg.Observability
	metricsConfig := &observabilityConfig.Metrics
	instrumentationHandler, err := metrics.ProvideMetricsInstrumentationHandlerForServer(metricsConfig, logger)
	if err != nil {
		return nil, err
	}
	servicesConfigurations := &cfg.Services
	authenticationConfig := &servicesConfigurations.Auth
	authenticator := authentication.ProvideArgon2Authenticator(logger)
	dataManager, err := config.ProvideDatabaseClient(ctx, logger, cfg)
	if err != nil {
		return nil, err
	}
	userDataManager := database.ProvideUserDataManager(dataManager)
	apiClientDataManager := database.ProvideAPIClientDataManager(dataManager)
	householdUserMembershipDataManager := database.ProvideHouseholdUserMembershipDataManager(dataManager)
	cookieConfig := authenticationConfig.Cookies
	sessionManager, err := config2.ProvideSessionManager(cookieConfig, dataManager)
	if err != nil {
		return nil, err
	}
	encodingConfig := cfg.Encoding
	contentType := encoding.ProvideContentType(encodingConfig)
	serverEncoderDecoder := encoding.ProvideServerEncoderDecoder(logger, contentType)
	authService, err := authentication2.ProvideService(logger, authenticationConfig, authenticator, userDataManager, apiClientDataManager, householdUserMembershipDataManager, sessionManager, serverEncoderDecoder)
	if err != nil {
		return nil, err
	}
	householdDataManager := database.ProvideHouseholdDataManager(dataManager)
	unitCounterProvider, err := metrics.ProvideUnitCounterProvider(metricsConfig, logger)
	if err != nil {
		return nil, err
	}
	imageUploadProcessor := images.NewImageUploadProcessor(logger)
	uploadsConfig := &cfg.Uploads
	storageConfig := &uploadsConfig.Storage
	routeParamManager := chi.NewRouteParamManager()
	uploader, err := storage.NewUploadManager(ctx, logger, storageConfig, routeParamManager)
	if err != nil {
		return nil, err
	}
	uploadManager := uploads.ProvideUploadManager(uploader)
	userDataService := users.ProvideUsersService(authenticationConfig, logger, userDataManager, householdDataManager, authenticator, serverEncoderDecoder, unitCounterProvider, imageUploadProcessor, uploadManager, routeParamManager)
	householdsConfig := servicesConfigurations.Households
	configConfig := &cfg.Events
	publisherProvider, err := config3.ProvidePublisherProvider(logger, configConfig)
	if err != nil {
		return nil, err
	}
	householdDataService, err := households.ProvideService(logger, householdsConfig, householdDataManager, householdUserMembershipDataManager, serverEncoderDecoder, unitCounterProvider, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	apiclientsConfig := apiclients.ProvideConfig(authenticationConfig)
	apiClientDataService := apiclients.ProvideAPIClientsService(logger, apiClientDataManager, userDataManager, authenticator, serverEncoderDecoder, unitCounterProvider, routeParamManager, apiclientsConfig)
	consumerProvider, err := config3.ProvideConsumerProvider(logger, configConfig)
	if err != nil {
		return nil, err
	}
	websocketDataService, err := websockets.ProvideService(ctx, authenticationConfig, logger, serverEncoderDecoder, consumerProvider)
	if err != nil {
		return nil, err
	}
	validinstrumentsConfig := &servicesConfigurations.ValidInstruments
	validInstrumentDataManager := database.ProvideValidInstrumentDataManager(dataManager)
	indexManagerProvider := elasticsearch.ProvideIndexManagerProvider()
	validInstrumentDataService, err := validinstruments.ProvideService(ctx, logger, validinstrumentsConfig, validInstrumentDataManager, serverEncoderDecoder, indexManagerProvider, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	validingredientsConfig := &servicesConfigurations.ValidIngredients
	validIngredientDataManager := database.ProvideValidIngredientDataManager(dataManager)
	validIngredientDataService, err := validingredients.ProvideService(ctx, logger, validingredientsConfig, validIngredientDataManager, serverEncoderDecoder, indexManagerProvider, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	validpreparationsConfig := &servicesConfigurations.ValidPreparations
	validPreparationDataManager := database.ProvideValidPreparationDataManager(dataManager)
	validPreparationDataService, err := validpreparations.ProvideService(ctx, logger, validpreparationsConfig, validPreparationDataManager, serverEncoderDecoder, indexManagerProvider, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	validingredientpreparationsConfig := &servicesConfigurations.ValidIngredientPreparations
	validIngredientPreparationDataManager := database.ProvideValidIngredientPreparationDataManager(dataManager)
	validIngredientPreparationDataService, err := validingredientpreparations.ProvideService(ctx, logger, validingredientpreparationsConfig, validIngredientPreparationDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	recipesConfig := &servicesConfigurations.Recipes
	recipeDataManager := database.ProvideRecipeDataManager(dataManager)
	recipeDataService, err := recipes.ProvideService(ctx, logger, recipesConfig, recipeDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	recipestepsConfig := &servicesConfigurations.RecipeSteps
	recipeStepDataManager := database.ProvideRecipeStepDataManager(dataManager)
	recipeStepDataService, err := recipesteps.ProvideService(ctx, logger, recipestepsConfig, recipeStepDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	recipestepinstrumentsConfig := &servicesConfigurations.RecipeStepInstruments
	recipeStepInstrumentDataManager := database.ProvideRecipeStepInstrumentDataManager(dataManager)
	recipeStepInstrumentDataService, err := recipestepinstruments.ProvideService(ctx, logger, recipestepinstrumentsConfig, recipeStepInstrumentDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	recipestepingredientsConfig := &servicesConfigurations.RecipeStepIngredients
	recipeStepIngredientDataManager := database.ProvideRecipeStepIngredientDataManager(dataManager)
	recipeStepIngredientDataService, err := recipestepingredients.ProvideService(ctx, logger, recipestepingredientsConfig, recipeStepIngredientDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	recipestepproductsConfig := &servicesConfigurations.RecipeStepProducts
	recipeStepProductDataManager := database.ProvideRecipeStepProductDataManager(dataManager)
	recipeStepProductDataService, err := recipestepproducts.ProvideService(ctx, logger, recipestepproductsConfig, recipeStepProductDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	mealplansConfig := &servicesConfigurations.MealPlans
	mealPlanDataManager := database.ProvideMealPlanDataManager(dataManager)
	mealPlanDataService, err := mealplans.ProvideService(ctx, logger, mealplansConfig, mealPlanDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	mealplanoptionsConfig := &servicesConfigurations.MealPlanOptions
	mealPlanOptionDataManager := database.ProvideMealPlanOptionDataManager(dataManager)
	mealPlanOptionDataService, err := mealplanoptions.ProvideService(ctx, logger, mealplanoptionsConfig, mealPlanOptionDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	mealplanoptionvotesConfig := &servicesConfigurations.MealPlanOptionVotes
	mealPlanOptionVoteDataManager := database.ProvideMealPlanOptionVoteDataManager(dataManager)
	mealPlanOptionVoteDataService, err := mealplanoptionvotes.ProvideService(ctx, logger, mealplanoptionvotesConfig, mealPlanOptionVoteDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	webhooksConfig := &servicesConfigurations.Webhooks
	webhookDataManager := database.ProvideWebhookDataManager(dataManager)
	webhookDataService, err := webhooks.ProvideWebhooksService(logger, webhooksConfig, webhookDataManager, serverEncoderDecoder, routeParamManager, publisherProvider)
	if err != nil {
		return nil, err
	}
	adminUserDataManager := database.ProvideAdminUserDataManager(dataManager)
	adminService := admin.ProvideService(logger, authenticationConfig, authenticator, adminUserDataManager, sessionManager, serverEncoderDecoder, routeParamManager)
	routingConfig := &cfg.Routing
	router := chi.NewRouter(logger, routingConfig)
	httpServer, err := server.ProvideHTTPServer(ctx, serverConfig, instrumentationHandler, authService, userDataService, householdDataService, apiClientDataService, websocketDataService, validInstrumentDataService, validIngredientDataService, validPreparationDataService, validIngredientPreparationDataService, recipeDataService, recipeStepDataService, recipeStepInstrumentDataService, recipeStepIngredientDataService, recipeStepProductDataService, mealPlanDataService, mealPlanOptionDataService, mealPlanOptionVoteDataService, webhookDataService, adminService, logger, serverEncoderDecoder, router)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}
