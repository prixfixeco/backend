package recipes

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel/trace"

	"github.com/prixfixeco/api_server/internal/customerdata"
	mockencoding "github.com/prixfixeco/api_server/internal/encoding/mock"
	mockpublishers "github.com/prixfixeco/api_server/internal/messagequeue/mock"
	"github.com/prixfixeco/api_server/internal/observability/logging"
	"github.com/prixfixeco/api_server/internal/observability/tracing"
	mockrouting "github.com/prixfixeco/api_server/internal/routing/mock"
	mocktypes "github.com/prixfixeco/api_server/pkg/types/mock"
)

func buildTestService() *service {
	return &service{
		logger:            logging.NewNoopLogger(),
		recipeDataManager: &mocktypes.RecipeDataManager{},
		recipeIDFetcher:   func(req *http.Request) string { return "" },
		encoderDecoder:    mockencoding.NewMockEncoderDecoder(),
		tracer:            tracing.NewTracerForTest("test"),
	}
}

func TestProvideRecipesService(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		rpm := mockrouting.NewRouteParamManager()
		rpm.On(
			"BuildRouteParamStringIDFetcher",
			RecipeIDURIParamKey,
		).Return(func(*http.Request) string { return "" })

		cfg := Config{
			DataChangesTopicName: "data_changes",
		}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProviderPublisher", cfg.DataChangesTopicName).Return(&mockpublishers.Publisher{}, nil)

		s, err := ProvideService(
			ctx,
			logging.NewNoopLogger(),
			&cfg,
			&mocktypes.RecipeDataManager{},
			mockencoding.NewMockEncoderDecoder(),
			rpm,
			pp,
			&customerdata.MockCollector{},
			trace.NewNoopTracerProvider(),
		)

		assert.NotNil(t, s)
		assert.NoError(t, err)

		mock.AssertExpectationsForObjects(t, rpm, pp)
	})

	T.Run("with error providing data changes producer", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		cfg := Config{
			DataChangesTopicName: "data_changes",
		}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProviderPublisher", cfg.DataChangesTopicName).Return((*mockpublishers.Publisher)(nil), errors.New("blah"))

		s, err := ProvideService(
			ctx,
			logging.NewNoopLogger(),
			&cfg,
			&mocktypes.RecipeDataManager{},
			mockencoding.NewMockEncoderDecoder(),
			nil,
			pp,
			&customerdata.MockCollector{},
			trace.NewNoopTracerProvider(),
		)

		assert.Nil(t, s)
		assert.Error(t, err)

		mock.AssertExpectationsForObjects(t, pp)
	})
}
