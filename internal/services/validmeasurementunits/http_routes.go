package validmeasurementunits

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/segmentio/ksuid"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/internal/observability/keys"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
)

const (
	// ValidMeasurementUnitIDURIParamKey is a standard string that we'll use to refer to valid ingredient IDs with.
	ValidMeasurementUnitIDURIParamKey = "validMeasurementUnitID"
)

// CreateHandler is our valid ingredient creation route.
func (s *service) CreateHandler(res http.ResponseWriter, req *http.Request) {
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
	providedInput := new(types.ValidMeasurementUnitCreationRequestInput)
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

	input := types.ValidMeasurementUnitDatabaseCreationInputFromValidMeasurementUnitCreationInput(providedInput)
	input.ID = ksuid.New().String()

	tracing.AttachValidMeasurementUnitIDToSpan(span, input.ID)

	validMeasurementUnit, err := s.validMeasurementUnitDataManager.CreateValidMeasurementUnit(ctx, input)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "creating valid ingredient")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	if s.dataChangesPublisher != nil {
		dcm := &types.DataChangeMessage{
			DataType:             types.ValidMeasurementUnitDataType,
			EventType:            types.ValidMeasurementUnitCreatedCustomerEventType,
			ValidMeasurementUnit: validMeasurementUnit,
			AttributableToUserID: sessionCtxData.Requester.UserID,
		}

		if err = s.dataChangesPublisher.Publish(ctx, dcm); err != nil {
			observability.AcknowledgeError(err, logger, span, "publishing to data changes topic")
		}
	}

	s.encoderDecoder.EncodeResponseWithStatus(ctx, res, validMeasurementUnit, http.StatusCreated)
}

// ReadHandler returns a GET handler that returns a valid ingredient.
func (s *service) ReadHandler(res http.ResponseWriter, req *http.Request) {
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

	// determine valid ingredient ID.
	validMeasurementUnitID := s.validMeasurementUnitIDFetcher(req)
	tracing.AttachValidMeasurementUnitIDToSpan(span, validMeasurementUnitID)
	logger = logger.WithValue(keys.ValidMeasurementUnitIDKey, validMeasurementUnitID)

	// fetch valid ingredient from database.
	x, err := s.validMeasurementUnitDataManager.GetValidMeasurementUnit(ctx, validMeasurementUnitID)
	if errors.Is(err, sql.ErrNoRows) {
		s.encoderDecoder.EncodeNotFoundResponse(ctx, res)
		return
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving valid ingredient")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, x)
}

// ListHandler is our list route.
func (s *service) ListHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	filter := types.ExtractQueryFilter(req)
	logger := s.logger.WithRequest(req).
		WithValue(keys.FilterLimitKey, filter.Limit).
		WithValue(keys.FilterPageKey, filter.Page).
		WithValue(keys.FilterSortByKey, string(filter.SortBy))

	tracing.AttachRequestToSpan(span, req)
	tracing.AttachFilterToSpan(span, filter.Page, filter.Limit, string(filter.SortBy))

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	validMeasurementUnits, err := s.validMeasurementUnitDataManager.GetValidMeasurementUnits(ctx, filter)
	if errors.Is(err, sql.ErrNoRows) {
		// in the event no rows exist, return an empty list.
		validMeasurementUnits = &types.ValidMeasurementUnitList{ValidMeasurementUnits: []*types.ValidMeasurementUnit{}}
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving valid ingredients")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, validMeasurementUnits)
}

// SearchHandler is our search route.
func (s *service) SearchHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	query := req.URL.Query().Get(types.SearchQueryKey)
	filter := types.ExtractQueryFilter(req)
	logger := s.logger.WithRequest(req).
		WithValue(keys.FilterLimitKey, filter.Limit).
		WithValue(keys.FilterPageKey, filter.Page).
		WithValue(keys.FilterSortByKey, string(filter.SortBy)).
		WithValue(keys.SearchQueryKey, query)

	tracing.AttachRequestToSpan(span, req)
	tracing.AttachFilterToSpan(span, filter.Page, filter.Limit, string(filter.SortBy))

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		logger.Error(err, "retrieving session context data")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "unauthenticated", http.StatusUnauthorized)
		return
	}

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)

	validMeasurementUnits, err := s.validMeasurementUnitDataManager.SearchForValidMeasurementUnits(ctx, query)
	if errors.Is(err, sql.ErrNoRows) {
		// in the event no rows exist, return an empty list.
		validMeasurementUnits = []*types.ValidMeasurementUnit{}
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "searching valid ingredients")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, validMeasurementUnits)
}

// UpdateHandler returns a handler that updates a valid ingredient.
func (s *service) UpdateHandler(res http.ResponseWriter, req *http.Request) {
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

	// check for parsed input attached to session context data.
	input := new(types.ValidMeasurementUnitUpdateRequestInput)
	if err = s.encoderDecoder.DecodeRequest(ctx, req, input); err != nil {
		logger.Error(err, "error encountered decoding request body")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, "invalid request content", http.StatusBadRequest)
		return
	}

	if err = input.ValidateWithContext(ctx); err != nil {
		logger.Error(err, "provided input was invalid")
		s.encoderDecoder.EncodeErrorResponse(ctx, res, err.Error(), http.StatusBadRequest)
		return
	}

	// determine valid ingredient ID.
	validMeasurementUnitID := s.validMeasurementUnitIDFetcher(req)
	tracing.AttachValidMeasurementUnitIDToSpan(span, validMeasurementUnitID)
	logger = logger.WithValue(keys.ValidMeasurementUnitIDKey, validMeasurementUnitID)

	// fetch valid ingredient from database.
	validMeasurementUnit, err := s.validMeasurementUnitDataManager.GetValidMeasurementUnit(ctx, validMeasurementUnitID)
	if errors.Is(err, sql.ErrNoRows) {
		s.encoderDecoder.EncodeNotFoundResponse(ctx, res)
		return
	} else if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving valid ingredient for update")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	// update the valid ingredient.
	validMeasurementUnit.Update(input)

	if err = s.validMeasurementUnitDataManager.UpdateValidMeasurementUnit(ctx, validMeasurementUnit); err != nil {
		observability.AcknowledgeError(err, logger, span, "updating valid ingredient")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	if s.dataChangesPublisher != nil {
		dcm := &types.DataChangeMessage{
			DataType:             types.ValidMeasurementUnitDataType,
			EventType:            types.ValidMeasurementUnitUpdatedCustomerEventType,
			ValidMeasurementUnit: validMeasurementUnit,
			AttributableToUserID: sessionCtxData.Requester.UserID,
		}

		if err = s.dataChangesPublisher.Publish(ctx, dcm); err != nil {
			observability.AcknowledgeError(err, logger, span, "publishing data change message")
		}
	}

	// encode our response and peace.
	s.encoderDecoder.RespondWithData(ctx, res, validMeasurementUnit)
}

// ArchiveHandler returns a handler that archives a valid ingredient.
func (s *service) ArchiveHandler(res http.ResponseWriter, req *http.Request) {
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

	// determine valid ingredient ID.
	validMeasurementUnitID := s.validMeasurementUnitIDFetcher(req)
	tracing.AttachValidMeasurementUnitIDToSpan(span, validMeasurementUnitID)
	logger = logger.WithValue(keys.ValidMeasurementUnitIDKey, validMeasurementUnitID)

	exists, existenceCheckErr := s.validMeasurementUnitDataManager.ValidMeasurementUnitExists(ctx, validMeasurementUnitID)
	if existenceCheckErr != nil && !errors.Is(existenceCheckErr, sql.ErrNoRows) {
		observability.AcknowledgeError(existenceCheckErr, logger, span, "checking valid ingredient existence")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	} else if !exists || errors.Is(existenceCheckErr, sql.ErrNoRows) {
		s.encoderDecoder.EncodeNotFoundResponse(ctx, res)
		return
	}

	if err = s.validMeasurementUnitDataManager.ArchiveValidMeasurementUnit(ctx, validMeasurementUnitID); err != nil {
		observability.AcknowledgeError(err, logger, span, "updating valid ingredient")
		s.encoderDecoder.EncodeUnspecifiedInternalServerErrorResponse(ctx, res)
		return
	}

	if s.dataChangesPublisher != nil {
		dcm := &types.DataChangeMessage{
			DataType:             types.ValidMeasurementUnitDataType,
			EventType:            types.ValidMeasurementUnitArchivedCustomerEventType,
			AttributableToUserID: sessionCtxData.Requester.UserID,
		}

		if err = s.dataChangesPublisher.Publish(ctx, dcm); err != nil {
			observability.AcknowledgeError(err, logger, span, "publishing data change message")
		}
	}

	// encode our response and peace.
	res.WriteHeader(http.StatusNoContent)
}