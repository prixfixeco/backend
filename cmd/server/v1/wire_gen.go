// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"gitlab.com/prixfixe/prixfixe/database/v1"
	"gitlab.com/prixfixe/prixfixe/internal/v1/auth"
	"gitlab.com/prixfixe/prixfixe/internal/v1/config"
	"gitlab.com/prixfixe/prixfixe/internal/v1/encoding"
	"gitlab.com/prixfixe/prixfixe/internal/v1/metrics"
	"gitlab.com/prixfixe/prixfixe/server/v1"
	"gitlab.com/prixfixe/prixfixe/server/v1/http"
	auth2 "gitlab.com/prixfixe/prixfixe/services/v1/auth"
	"gitlab.com/prixfixe/prixfixe/services/v1/frontend"
	"gitlab.com/prixfixe/prixfixe/services/v1/ingredienttagmappings"
	"gitlab.com/prixfixe/prixfixe/services/v1/invitations"
	"gitlab.com/prixfixe/prixfixe/services/v1/iterationmedias"
	"gitlab.com/prixfixe/prixfixe/services/v1/oauth2clients"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipeiterations"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipeiterationsteps"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipes"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipestepingredients"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipesteppreparations"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipesteps"
	"gitlab.com/prixfixe/prixfixe/services/v1/recipetags"
	"gitlab.com/prixfixe/prixfixe/services/v1/reports"
	"gitlab.com/prixfixe/prixfixe/services/v1/requiredpreparationinstruments"
	"gitlab.com/prixfixe/prixfixe/services/v1/users"
	"gitlab.com/prixfixe/prixfixe/services/v1/validingredientpreparations"
	"gitlab.com/prixfixe/prixfixe/services/v1/validingredients"
	"gitlab.com/prixfixe/prixfixe/services/v1/validingredienttags"
	"gitlab.com/prixfixe/prixfixe/services/v1/validinstruments"
	"gitlab.com/prixfixe/prixfixe/services/v1/validpreparations"
	"gitlab.com/prixfixe/prixfixe/services/v1/webhooks"
	"gitlab.com/verygoodsoftwarenotvirus/logging/v1"
	"gitlab.com/verygoodsoftwarenotvirus/newsman"
)

// Injectors from wire.go:

func BuildServer(ctx context.Context, cfg *config.ServerConfig, logger logging.Logger, database2 database.Database) (*server.Server, error) {
	bcryptHashCost := auth.ProvideBcryptHashCost()
	authenticator := auth.ProvideBcryptAuthenticator(bcryptHashCost, logger)
	userDataManager := users.ProvideUserDataManager(database2)
	clientIDFetcher := httpserver.ProvideOAuth2ClientsServiceClientIDFetcher(logger)
	encoderDecoder := encoding.ProvideResponseEncoder()
	unitCounterProvider := metrics.ProvideUnitCounterProvider()
	service, err := oauth2clients.ProvideOAuth2ClientsService(logger, database2, authenticator, clientIDFetcher, encoderDecoder, unitCounterProvider)
	if err != nil {
		return nil, err
	}
	oAuth2ClientValidator := auth2.ProvideOAuth2ClientValidator(service)
	authSettings := config.ProvideConfigAuthSettings(cfg)
	databaseSettings := config.ProvideConfigDatabaseSettings(cfg)
	db, err := config.ProvideDatabaseConnection(logger, databaseSettings)
	if err != nil {
		return nil, err
	}
	sessionManager := config.ProvideSessionManager(authSettings, db)
	authService, err := auth2.ProvideAuthService(logger, cfg, authenticator, userDataManager, oAuth2ClientValidator, sessionManager, encoderDecoder)
	if err != nil {
		return nil, err
	}
	frontendSettings := config.ProvideConfigFrontendSettings(cfg)
	frontendService := frontend.ProvideFrontendService(logger, frontendSettings)
	validInstrumentDataManager := validinstruments.ProvideValidInstrumentDataManager(database2)
	validInstrumentIDFetcher := httpserver.ProvideValidInstrumentsServiceValidInstrumentIDFetcher(logger)
	websocketAuthFunc := auth2.ProvideWebsocketAuthFunc(authService)
	typeNameManipulationFunc := httpserver.ProvideNewsmanTypeNameManipulationFunc()
	newsmanNewsman := newsman.NewNewsman(websocketAuthFunc, typeNameManipulationFunc)
	reporter := ProvideReporter(newsmanNewsman)
	validinstrumentsService, err := validinstruments.ProvideValidInstrumentsService(logger, validInstrumentDataManager, validInstrumentIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	validInstrumentDataServer := validinstruments.ProvideValidInstrumentDataServer(validinstrumentsService)
	validIngredientDataManager := validingredients.ProvideValidIngredientDataManager(database2)
	validIngredientIDFetcher := httpserver.ProvideValidIngredientsServiceValidIngredientIDFetcher(logger)
	validingredientsService, err := validingredients.ProvideValidIngredientsService(logger, validIngredientDataManager, validIngredientIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	validIngredientDataServer := validingredients.ProvideValidIngredientDataServer(validingredientsService)
	validIngredientTagDataManager := validingredienttags.ProvideValidIngredientTagDataManager(database2)
	validIngredientTagIDFetcher := httpserver.ProvideValidIngredientTagsServiceValidIngredientTagIDFetcher(logger)
	validingredienttagsService, err := validingredienttags.ProvideValidIngredientTagsService(logger, validIngredientTagDataManager, validIngredientTagIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	validIngredientTagDataServer := validingredienttags.ProvideValidIngredientTagDataServer(validingredienttagsService)
	ingredientTagMappingDataManager := ingredienttagmappings.ProvideIngredientTagMappingDataManager(database2)
	ingredienttagmappingsValidIngredientIDFetcher := httpserver.ProvideIngredientTagMappingsServiceValidIngredientIDFetcher(logger)
	ingredientTagMappingIDFetcher := httpserver.ProvideIngredientTagMappingsServiceIngredientTagMappingIDFetcher(logger)
	ingredienttagmappingsService, err := ingredienttagmappings.ProvideIngredientTagMappingsService(logger, validIngredientDataManager, ingredientTagMappingDataManager, ingredienttagmappingsValidIngredientIDFetcher, ingredientTagMappingIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	ingredientTagMappingDataServer := ingredienttagmappings.ProvideIngredientTagMappingDataServer(ingredienttagmappingsService)
	validPreparationDataManager := validpreparations.ProvideValidPreparationDataManager(database2)
	validPreparationIDFetcher := httpserver.ProvideValidPreparationsServiceValidPreparationIDFetcher(logger)
	validpreparationsService, err := validpreparations.ProvideValidPreparationsService(logger, validPreparationDataManager, validPreparationIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	validPreparationDataServer := validpreparations.ProvideValidPreparationDataServer(validpreparationsService)
	requiredPreparationInstrumentDataManager := requiredpreparationinstruments.ProvideRequiredPreparationInstrumentDataManager(database2)
	requiredpreparationinstrumentsValidPreparationIDFetcher := httpserver.ProvideRequiredPreparationInstrumentsServiceValidPreparationIDFetcher(logger)
	requiredPreparationInstrumentIDFetcher := httpserver.ProvideRequiredPreparationInstrumentsServiceRequiredPreparationInstrumentIDFetcher(logger)
	requiredpreparationinstrumentsService, err := requiredpreparationinstruments.ProvideRequiredPreparationInstrumentsService(logger, validPreparationDataManager, requiredPreparationInstrumentDataManager, requiredpreparationinstrumentsValidPreparationIDFetcher, requiredPreparationInstrumentIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	requiredPreparationInstrumentDataServer := requiredpreparationinstruments.ProvideRequiredPreparationInstrumentDataServer(requiredpreparationinstrumentsService)
	validIngredientPreparationDataManager := validingredientpreparations.ProvideValidIngredientPreparationDataManager(database2)
	validingredientpreparationsValidIngredientIDFetcher := httpserver.ProvideValidIngredientPreparationsServiceValidIngredientIDFetcher(logger)
	validIngredientPreparationIDFetcher := httpserver.ProvideValidIngredientPreparationsServiceValidIngredientPreparationIDFetcher(logger)
	validingredientpreparationsService, err := validingredientpreparations.ProvideValidIngredientPreparationsService(logger, validIngredientDataManager, validIngredientPreparationDataManager, validingredientpreparationsValidIngredientIDFetcher, validIngredientPreparationIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	validIngredientPreparationDataServer := validingredientpreparations.ProvideValidIngredientPreparationDataServer(validingredientpreparationsService)
	recipeDataManager := recipes.ProvideRecipeDataManager(database2)
	recipeIDFetcher := httpserver.ProvideRecipesServiceRecipeIDFetcher(logger)
	userIDFetcher := httpserver.ProvideRecipesServiceUserIDFetcher()
	recipesService, err := recipes.ProvideRecipesService(logger, recipeDataManager, recipeIDFetcher, userIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeDataServer := recipes.ProvideRecipeDataServer(recipesService)
	recipeTagDataManager := recipetags.ProvideRecipeTagDataManager(database2)
	recipetagsRecipeIDFetcher := httpserver.ProvideRecipeTagsServiceRecipeIDFetcher(logger)
	recipeTagIDFetcher := httpserver.ProvideRecipeTagsServiceRecipeTagIDFetcher(logger)
	recipetagsUserIDFetcher := httpserver.ProvideRecipeTagsServiceUserIDFetcher()
	recipetagsService, err := recipetags.ProvideRecipeTagsService(logger, recipeDataManager, recipeTagDataManager, recipetagsRecipeIDFetcher, recipeTagIDFetcher, recipetagsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeTagDataServer := recipetags.ProvideRecipeTagDataServer(recipetagsService)
	recipeStepDataManager := recipesteps.ProvideRecipeStepDataManager(database2)
	recipestepsRecipeIDFetcher := httpserver.ProvideRecipeStepsServiceRecipeIDFetcher(logger)
	recipeStepIDFetcher := httpserver.ProvideRecipeStepsServiceRecipeStepIDFetcher(logger)
	recipestepsUserIDFetcher := httpserver.ProvideRecipeStepsServiceUserIDFetcher()
	recipestepsService, err := recipesteps.ProvideRecipeStepsService(logger, recipeDataManager, recipeStepDataManager, recipestepsRecipeIDFetcher, recipeStepIDFetcher, recipestepsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeStepDataServer := recipesteps.ProvideRecipeStepDataServer(recipestepsService)
	recipeStepPreparationDataManager := recipesteppreparations.ProvideRecipeStepPreparationDataManager(database2)
	recipesteppreparationsRecipeIDFetcher := httpserver.ProvideRecipeStepPreparationsServiceRecipeIDFetcher(logger)
	recipesteppreparationsRecipeStepIDFetcher := httpserver.ProvideRecipeStepPreparationsServiceRecipeStepIDFetcher(logger)
	recipeStepPreparationIDFetcher := httpserver.ProvideRecipeStepPreparationsServiceRecipeStepPreparationIDFetcher(logger)
	recipesteppreparationsUserIDFetcher := httpserver.ProvideRecipeStepPreparationsServiceUserIDFetcher()
	recipesteppreparationsService, err := recipesteppreparations.ProvideRecipeStepPreparationsService(logger, recipeDataManager, recipeStepDataManager, recipeStepPreparationDataManager, recipesteppreparationsRecipeIDFetcher, recipesteppreparationsRecipeStepIDFetcher, recipeStepPreparationIDFetcher, recipesteppreparationsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeStepPreparationDataServer := recipesteppreparations.ProvideRecipeStepPreparationDataServer(recipesteppreparationsService)
	recipeStepIngredientDataManager := recipestepingredients.ProvideRecipeStepIngredientDataManager(database2)
	recipestepingredientsRecipeIDFetcher := httpserver.ProvideRecipeStepIngredientsServiceRecipeIDFetcher(logger)
	recipestepingredientsRecipeStepIDFetcher := httpserver.ProvideRecipeStepIngredientsServiceRecipeStepIDFetcher(logger)
	recipeStepIngredientIDFetcher := httpserver.ProvideRecipeStepIngredientsServiceRecipeStepIngredientIDFetcher(logger)
	recipestepingredientsUserIDFetcher := httpserver.ProvideRecipeStepIngredientsServiceUserIDFetcher()
	recipestepingredientsService, err := recipestepingredients.ProvideRecipeStepIngredientsService(logger, recipeDataManager, recipeStepDataManager, recipeStepIngredientDataManager, recipestepingredientsRecipeIDFetcher, recipestepingredientsRecipeStepIDFetcher, recipeStepIngredientIDFetcher, recipestepingredientsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeStepIngredientDataServer := recipestepingredients.ProvideRecipeStepIngredientDataServer(recipestepingredientsService)
	recipeIterationDataManager := recipeiterations.ProvideRecipeIterationDataManager(database2)
	recipeiterationsRecipeIDFetcher := httpserver.ProvideRecipeIterationsServiceRecipeIDFetcher(logger)
	recipeIterationIDFetcher := httpserver.ProvideRecipeIterationsServiceRecipeIterationIDFetcher(logger)
	recipeiterationsUserIDFetcher := httpserver.ProvideRecipeIterationsServiceUserIDFetcher()
	recipeiterationsService, err := recipeiterations.ProvideRecipeIterationsService(logger, recipeDataManager, recipeIterationDataManager, recipeiterationsRecipeIDFetcher, recipeIterationIDFetcher, recipeiterationsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeIterationDataServer := recipeiterations.ProvideRecipeIterationDataServer(recipeiterationsService)
	recipeIterationStepDataManager := recipeiterationsteps.ProvideRecipeIterationStepDataManager(database2)
	recipeiterationstepsRecipeIDFetcher := httpserver.ProvideRecipeIterationStepsServiceRecipeIDFetcher(logger)
	recipeIterationStepIDFetcher := httpserver.ProvideRecipeIterationStepsServiceRecipeIterationStepIDFetcher(logger)
	recipeiterationstepsUserIDFetcher := httpserver.ProvideRecipeIterationStepsServiceUserIDFetcher()
	recipeiterationstepsService, err := recipeiterationsteps.ProvideRecipeIterationStepsService(logger, recipeDataManager, recipeIterationStepDataManager, recipeiterationstepsRecipeIDFetcher, recipeIterationStepIDFetcher, recipeiterationstepsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	recipeIterationStepDataServer := recipeiterationsteps.ProvideRecipeIterationStepDataServer(recipeiterationstepsService)
	iterationMediaDataManager := iterationmedias.ProvideIterationMediaDataManager(database2)
	iterationmediasRecipeIDFetcher := httpserver.ProvideIterationMediasServiceRecipeIDFetcher(logger)
	iterationmediasRecipeIterationIDFetcher := httpserver.ProvideIterationMediasServiceRecipeIterationIDFetcher(logger)
	iterationMediaIDFetcher := httpserver.ProvideIterationMediasServiceIterationMediaIDFetcher(logger)
	iterationmediasUserIDFetcher := httpserver.ProvideIterationMediasServiceUserIDFetcher()
	iterationmediasService, err := iterationmedias.ProvideIterationMediasService(logger, recipeDataManager, recipeIterationDataManager, iterationMediaDataManager, iterationmediasRecipeIDFetcher, iterationmediasRecipeIterationIDFetcher, iterationMediaIDFetcher, iterationmediasUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	iterationMediaDataServer := iterationmedias.ProvideIterationMediaDataServer(iterationmediasService)
	invitationDataManager := invitations.ProvideInvitationDataManager(database2)
	invitationIDFetcher := httpserver.ProvideInvitationsServiceInvitationIDFetcher(logger)
	invitationsUserIDFetcher := httpserver.ProvideInvitationsServiceUserIDFetcher()
	invitationsService, err := invitations.ProvideInvitationsService(logger, invitationDataManager, invitationIDFetcher, invitationsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	invitationDataServer := invitations.ProvideInvitationDataServer(invitationsService)
	reportDataManager := reports.ProvideReportDataManager(database2)
	reportIDFetcher := httpserver.ProvideReportsServiceReportIDFetcher(logger)
	reportsUserIDFetcher := httpserver.ProvideReportsServiceUserIDFetcher()
	reportsService, err := reports.ProvideReportsService(logger, reportDataManager, reportIDFetcher, reportsUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	reportDataServer := reports.ProvideReportDataServer(reportsService)
	usersUserIDFetcher := httpserver.ProvideUsersServiceUserIDFetcher(logger)
	usersService, err := users.ProvideUsersService(authSettings, logger, userDataManager, authenticator, usersUserIDFetcher, encoderDecoder, unitCounterProvider, reporter)
	if err != nil {
		return nil, err
	}
	userDataServer := users.ProvideUserDataServer(usersService)
	oAuth2ClientDataServer := oauth2clients.ProvideOAuth2ClientDataServer(service)
	webhookDataManager := webhooks.ProvideWebhookDataManager(database2)
	webhooksUserIDFetcher := httpserver.ProvideWebhooksServiceUserIDFetcher()
	webhookIDFetcher := httpserver.ProvideWebhooksServiceWebhookIDFetcher(logger)
	webhooksService, err := webhooks.ProvideWebhooksService(logger, webhookDataManager, webhooksUserIDFetcher, webhookIDFetcher, encoderDecoder, unitCounterProvider, newsmanNewsman)
	if err != nil {
		return nil, err
	}
	webhookDataServer := webhooks.ProvideWebhookDataServer(webhooksService)
	httpserverServer, err := httpserver.ProvideServer(ctx, cfg, authService, frontendService, validInstrumentDataServer, validIngredientDataServer, validIngredientTagDataServer, ingredientTagMappingDataServer, validPreparationDataServer, requiredPreparationInstrumentDataServer, validIngredientPreparationDataServer, recipeDataServer, recipeTagDataServer, recipeStepDataServer, recipeStepPreparationDataServer, recipeStepIngredientDataServer, recipeIterationDataServer, recipeIterationStepDataServer, iterationMediaDataServer, invitationDataServer, reportDataServer, userDataServer, oAuth2ClientDataServer, webhookDataServer, database2, logger, encoderDecoder, newsmanNewsman)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.ProvideServer(cfg, httpserverServer)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}

// wire.go:

// ProvideReporter is an obligatory function that hopefully wire will eliminate for me one day.
func ProvideReporter(n *newsman.Newsman) newsman.Reporter {
	return n
}
