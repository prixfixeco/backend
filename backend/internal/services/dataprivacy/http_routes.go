package workers

import (
	"net/http"

	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"

	servertiming "github.com/mitchellh/go-server-timing"
)

// DataDeletionHandler deletes a user, which consequently deletes all data in the system.
func (s *service) DataDeletionHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	timing := servertiming.FromContext(ctx)
	logger := s.logger.WithRequest(req).WithSpan(span)
	logger.Info("meal plan finalization worker invoked")

	responseDetails := types.ResponseDetails{
		TraceID: span.SpanContext().TraceID().String(),
	}

	// determine user ID.
	sessionContextTimer := timing.NewMetric("session").WithDesc("fetch session context").Start()
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		errRes := types.NewAPIErrorResponse("unauthenticated", types.ErrFetchingSessionContextData, responseDetails)
		s.encoderDecoder.EncodeResponseWithStatus(ctx, res, errRes, http.StatusUnauthorized)
		return
	}
	sessionContextTimer.Stop()

	tracing.AttachSessionContextDataToSpan(span, sessionCtxData)
	logger = sessionCtxData.AttachToLogger(logger)
	responseDetails.CurrentHouseholdID = sessionCtxData.ActiveHouseholdID

	destroyUserDataTimer := timing.NewMetric("database").WithDesc("destroy user data").Start()
	if err = s.userDataManager.DeleteUser(ctx, sessionCtxData.Requester.UserID); err != nil {
		observability.AcknowledgeError(err, logger, span, "deleting user")
		errRes := types.NewAPIErrorResponse("database error", types.ErrTalkingToDatabase, responseDetails)
		s.encoderDecoder.EncodeResponseWithStatus(ctx, res, errRes, http.StatusInternalServerError)
		return
	}
	destroyUserDataTimer.Stop()

	if err = s.dataChangesPublisher.Publish(ctx, &types.DataChangeMessage{
		Context:     nil,
		EventType:   types.UserDataDestroyedCustomerEventType,
		UserID:      sessionCtxData.Requester.UserID,
		HouseholdID: sessionCtxData.ActiveHouseholdID,
	}); err != nil {
		observability.AcknowledgeError(err, logger, span, "publishing data")
	}

	responseValue := &types.APIResponse[types.DataDeletionResponse]{
		Data:    types.DataDeletionResponse{Successful: true},
		Details: responseDetails,
	}

	logger.Info("user data deleted")

	s.encoderDecoder.EncodeResponseWithStatus(ctx, res, responseValue, http.StatusAccepted)
}
