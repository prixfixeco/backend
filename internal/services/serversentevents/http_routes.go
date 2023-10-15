package serversentevents

import (
	"net/http"

	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
)

// StreamSubscriptionHandler subscribes to a stream.
func (s *service) StreamSubscriptionHandler(res http.ResponseWriter, req *http.Request) {
	ctx, span := s.tracer.StartSpan(req.Context())
	defer span.End()

	logger := s.logger.WithRequest(req)
	tracing.AttachRequestToSpan(span, req)

	responseDetails := types.ResponseDetails{
		TraceID: span.SpanContext().TraceID().String(),
	}

	// determine user ID.
	sessionCtxData, err := s.sessionContextDataFetcher(req)
	if err != nil {
		observability.AcknowledgeError(err, logger, span, "retrieving session context data")
		errRes := types.NewAPIErrorResponse("unauthenticated", types.ErrFetchingSessionContextData, responseDetails)
		s.encoderDecoder.EncodeResponseWithStatus(ctx, res, errRes, http.StatusUnauthorized)
		return
	}

	s.eventsServer.CreateStream(sessionCtxData.Requester.UserID)
	q := req.URL.Query()
	q.Set("stream", sessionCtxData.Requester.UserID)
	req.URL.RawQuery = q.Encode()

	go func() {
		<-ctx.Done()
		s.eventsServer.RemoveStream(sessionCtxData.Requester.UserID)
	}()

	s.eventsServer.HTTPHandler(res, req)
}