package apiclients

import (
	"net/http"
	"testing"

	mockauthn "github.com/prixfixeco/backend/internal/authentication/mock"
	"github.com/prixfixeco/backend/internal/database"
	mockencoding "github.com/prixfixeco/backend/internal/encoding/mock"
	mockpublishers "github.com/prixfixeco/backend/internal/messagequeue/mock"
	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/internal/random"
	mockrandom "github.com/prixfixeco/backend/internal/random/mock"
	mockrouting "github.com/prixfixeco/backend/internal/routing/mock"
	authservice "github.com/prixfixeco/backend/internal/services/authentication"
	mocktypes "github.com/prixfixeco/backend/pkg/types/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildTestService(t *testing.T) *service {
	t.Helper()

	return &service{
		apiClientDataManager:      database.NewMockDatabase(),
		logger:                    logging.NewNoopLogger(),
		encoderDecoder:            mockencoding.NewMockEncoderDecoder(),
		authenticator:             &mockauthn.Authenticator{},
		sessionContextDataFetcher: authservice.FetchContextFromRequest,
		urlClientIDExtractor:      func(req *http.Request) string { return "" },
		secretGenerator:           &mockrandom.Generator{},
		tracer:                    tracing.NewTracerForTest(serviceName),
		cfg:                       &Config{},
	}
}

func TestProvideAPIClientsService(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()
		mockAPIClientDataManager := &mocktypes.APIClientDataManager{}

		rpm := mockrouting.NewRouteParamManager()
		rpm.On(
			"BuildRouteParamStringIDFetcher",
			APIClientIDURIParamKey,
		).Return(func(*http.Request) string { return "" })

		cfg := &Config{
			dataChangesTopicName: t.Name(),
		}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProvidePublisher", cfg.dataChangesTopicName).Return(&mockpublishers.Publisher{}, nil)

		s, err := ProvideAPIClientsService(
			logging.NewNoopLogger(),
			mockAPIClientDataManager,
			&mocktypes.UserDataManager{},
			&mockauthn.Authenticator{},
			mockencoding.NewMockEncoderDecoder(),
			rpm,
			cfg,
			tracing.NewNoopTracerProvider(),
			random.NewGenerator(logging.NewNoopLogger(), tracing.NewNoopTracerProvider()),
			pp,
		)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		mock.AssertExpectationsForObjects(t, mockAPIClientDataManager, rpm)
	})
}
