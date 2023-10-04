package meals

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/dinnerdonebetter/backend/internal/identifiers"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/keys"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
)

const (
	// MealIDURIParamKey is a standard string that we'll use to refer to meal IDs with.
	MealIDURIParamKey = "mealID"
)

var _ types.MealDataService = (*service)(nil)

// CreateMealHandler is our meal creation route.
func (s *service) CreateMealHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	logger := s.logger.WithRequest(req)
	tracing.AttachRequestToSpan(span, req)

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	// read parsed input struct from request body.
	providedInput := new(types.MealCreationRequestInput)
	if err = s.encoderDecoder.DecodeRequest(ctx, req, providedInput); err != nil {
		observability.AcknowledgeError(err, logger, span, "decoding request")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "invalid request content", http.StatusBadRequest)
		return
	}

	if err = providedInput.ValidateWithContext(ctx); err != nil {
		logger.WithValue(keys.ValidationErrorKey, err).Debug("provided input was invalid")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, err.Error(), http.StatusBadRequest)
		return
	}

	input := converters.ConvertMealCreationRequestInputToMealDatabaseCreationInput(providedInput)
	input.ID = identifiers.New()

	input.CreatedByUser = sessionCtxData.Requester.UserID
	tracing.AttachToSpan(span, keys.MealIDKey, input.ID)

	meal, err := s.mealDataManager.CreateMeal(ctx, input)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "creating meal")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	dcm := &types.DataChangeMessage{
		EventType:   types.MealCreatedCustomerEventType,
		Meal:        meal,
		HouseholdID: sessionCtxData.ActiveHouseholdID,
		UserID:      sessionCtxData.Requester.UserID,
	}

	if err = s.dataChangesPublisher.Publish(ctx, dcm); err != nil {
		observability.AcknowledgeError(err, logger, span, "publishing to data changes topic")
	}

	s.encoderDecoder.EncodeResponseWithStatus(ctx, res, meal, http.StatusCreated)
}

// ReadMealHandler returns a GET handler that returns a meal.
func (s *service) ReadMealHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	logger := s.logger.WithRequest(req)
	tracing.AttachRequestToSpan(span, req)

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	// determine meal ID.
	mealID := s.mealIDFetcher(req)
	tracing.AttachToSpan(span, keys.MealIDKey, mealID)
	logger = logger.WithValue(keys.MealIDKey, mealID)

	// fetch meal from database.
	x, err := s.mealDataManager.GetMeal(ctx, mealID)
	if errors.Is(err, sql.ErrNoRows) {
		s.encoderDecoder.EncodeNotFoundResponse(ctx, res)
		return
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving meal")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, x)
}

// ListMealsHandler is our list route.
func (s *service) ListMealsHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	filter := types.ExtractQueryFilterFromRequest(req)
	logger := s.logger.WithRequest(req).
		WithValue(keys.FilterLimitKey, filter.Limit).
		WithValue(keys.FilterPageKey, filter.Page).
		WithValue(keys.FilterSortByKey, filter.SortBy)

	tracing.AttachRequestToSpan(span, req)
	tracing.AttachFilterDataToSpan(span, filter.Page, filter.Limit, filter.SortBy)

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	meals, err := s.mealDataManager.GetMeals(ctx, filter)
	if errors.Is(err, sql.ErrNoRows) {
		// in the event no rows exist, return an empty list.
		meals = &types.QueryFilteredResult[types.Meal]{Data: []*types.Meal{}}
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving meals")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, meals)
}

// SearchMealsHandler is our list route.
func (s *service) SearchMealsHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	tracing.AttachRequestToSpan(span, req)

	useDB := !s.cfg.UseSearchService || strings.TrimSpace(strings.ToLower(req.URL.Query().Get("useDB"))) == "true"

	query := req.URL.Query().Get(types.SearchQueryKey)
	tracing.AttachToSpan(span, keys.SearchQueryKey, query)

	filter := types.ExtractQueryFilterFromRequest(req)
	tracing.AttachFilterDataToSpan(span, filter.Page, filter.Limit, filter.SortBy)

	logger := s.logger.WithRequest(req).
		WithValue(keys.FilterLimitKey, filter.Limit).
		WithValue(keys.FilterPageKey, filter.Page).
		WithValue(keys.FilterSortByKey, filter.SortBy).
		WithValue(keys.SearchQueryKey, query).
		WithValue("using_database", useDB)

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	meals := &types.QueryFilteredResult[types.Meal]{
		Pagination: filter.ToPagination(),
	}
	if useDB {
		meals, err = s.mealDataManager.SearchForMeals(ctx, query, filter)
	} else {
		var mealSubsets []*types.MealSearchSubset
		mealSubsets, err = s.searchIndex.Search(ctx, query)
		if err != nil {
			observability.AcknowledgeError(err, logger, span, "searching for meals")
			s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
			return
		}

		ids := []string{}
		for _, mealSubset := range mealSubsets {
			ids = append(ids, mealSubset.ID)
		}

		meals.Data, err = s.mealDataManager.GetMealsWithIDs(ctx, ids)
	}

	if errors.Is(err, sql.ErrNoRows) {
		// in the event no rows exist, return an empty list.
		meals = &types.QueryFilteredResult[types.Meal]{
			Pagination: filter.ToPagination(),
			Data:       []*types.Meal{},
		}
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "searching for meals")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, meals)
}

// ArchiveMealHandler returns a handler that archives a meal.
func (s *service) ArchiveMealHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	logger := s.logger.WithRequest(req)
	tracing.AttachRequestToSpan(span, req)

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	// determine meal ID.
	mealID := s.mealIDFetcher(req)
	tracing.AttachToSpan(span, keys.MealIDKey, mealID)
	logger = logger.WithValue(keys.MealIDKey, mealID)

	exists, existenceCheckErr := s.mealDataManager.MealExists(ctx, mealID)
	if existenceCheckErr != nil && !errors.Is(existenceCheckErr, sql.ErrNoRows) {
		observability.AcknowledgeError(existenceCheckErr, logger, span, "checking meal existence")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	} else if !exists || errors.Is(existenceCheckErr, sql.ErrNoRows) {
		s.encoderDecoder.EncodeNotFoundResponse(ctx, res)
		return
	}

	if err = s.mealDataManager.ArchiveMeal(ctx, mealID, sessionCtxData.Requester.UserID); err != nil {
		observability.AcknowledgeError(err, logger, span, "archiving meal")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	dcm := &types.DataChangeMessage{
		EventType:   types.MealArchivedCustomerEventType,
		MealID:      mealID,
		HouseholdID: sessionCtxData.ActiveHouseholdID,
		UserID:      sessionCtxData.Requester.UserID,
	}

	if err = s.dataChangesPublisher.Publish(ctx, dcm); err != nil {
		observability.AcknowledgeError(err, logger, span, "publishing data change message")
	}

	// encode our response and peace.
	res.WriteHeader(http.StatusNoContent)
}
