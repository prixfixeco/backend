package websockets

import (
	"context"
	"encoding/json"
	"time"

	"github.com/prixfixeco/backend/internal/observability"
	"github.com/prixfixeco/backend/internal/observability/keys"
	"github.com/prixfixeco/backend/pkg/types"

	"github.com/gorilla/websocket"
)

func (s *service) handleDataChange(ctx context.Context, payload []byte) error {
	_, span := s.tracer.StartSpan(ctx)
	defer span.End()

	var msg *types.DataChangeMessage
	if err := json.Unmarshal(payload, &msg); err != nil {
		observability.AcknowledgeError(err, s.logger, span, "unmarshalling data change message")
		return err
	}

	s.logger.WithValue("msg", msg).Debug("handling data change")

	s.connectionsHat.RLock()
	defer s.connectionsHat.RUnlock()

	for userID, connections := range s.connections {
		if msg.HouseholdID != userID {
			continue
		}
		logger := s.logger.WithValue(keys.UserIDKey, userID).WithValue("connection_count", len(connections))
		for i, conn := range connections {
			logger = logger.WithValue("connection_index", i)

			if err := conn.SetWriteDeadline(time.Now().Add(s.websocketDeadline)); err != nil {
				observability.AcknowledgeError(err, logger, span, "setting write deadline")
				continue
			}

			if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
				logger.Error(err, "writing message to websocket")
				s.logger.Error(err, "writing message to websocket")
			}
		}
	}

	return nil
}

func removeConnection(s []websocketConnection, index int) []websocketConnection {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *service) pingConnections() {
	ticker := time.NewTicker(s.pollDuration)

	s.logger.Debug("pinging websocket connections")

	for range ticker.C {
		s.connectionsHat.Lock()
		for userID, connections := range s.connections {
			l := s.logger.WithValue(keys.UserIDKey, userID).WithValue("connection_count", len(connections))
			for i, conn := range connections {
				l = l.WithValue("connection_index", i)
				if err := conn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(s.pollDuration/2)); err != nil {
					if closeErr := s.connections[userID][i].Close(); closeErr != nil {
						l.WithError(closeErr).Debug("error occurred closing troubled websocket connection")
					}

					l.Debug("removing websocket connection")

					s.connections[userID] = removeConnection(s.connections[userID], i)
					if len(s.connections[userID]) == 0 {
						delete(s.connections, userID)
					}
				}
			}
		}
		s.connectionsHat.Unlock()
	}
}
