package validmeasurementunits

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/encoding"
	"github.com/dinnerdonebetter/backend/internal/encoding/mock"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	mockpublishers "github.com/dinnerdonebetter/backend/internal/messagequeue/mock"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	mockrouting "github.com/dinnerdonebetter/backend/internal/routing/mock"
	textsearchcfg "github.com/dinnerdonebetter/backend/internal/search/text/config"
	mocktypes "github.com/dinnerdonebetter/backend/pkg/types/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildTestService() *service {
	return &service{
		logger:                          logging.NewNoopLogger(),
		validMeasurementUnitDataManager: &mocktypes.ValidMeasurementUnitDataManagerMock{},
		validMeasurementUnitIDFetcher:   func(req *http.Request) string { return "" },
		encoderDecoder:                  encoding.ProvideServerEncoderDecoder(nil, nil, encoding.ContentTypeJSON),
		tracer:                          tracing.NewTracerForTest("test"),
	}
}

func TestProvideValidMeasurementUnitsService(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := logging.NewNoopLogger()

		rpm := mockrouting.NewRouteParamManager()
		rpm.On(
			"BuildRouteParamStringIDFetcher",
			ValidMeasurementUnitIDURIParamKey,
		).Return(func(*http.Request) string { return "" })
		rpm.On(
			"BuildRouteParamStringIDFetcher",
			ValidIngredientIDURIParamKey,
		).Return(func(*http.Request) string { return "" })

		cfg := &Config{}
		msgCfg := &msgconfig.QueuesConfig{
			DataChangesTopicName: "data_changes",
		}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProvidePublisher", msgCfg.DataChangesTopicName).Return((*mockpublishers.Publisher)(nil), nil)
		s, err := ProvideService(
			ctx,
			logger,
			cfg,
			&textsearchcfg.Config{},
			&mocktypes.ValidMeasurementUnitDataManagerMock{},
			mockencoding.NewMockEncoderDecoder(),
			rpm,
			pp,
			tracing.NewNoopTracerProvider(),
			msgCfg,
		)

		assert.NotNil(t, s)
		assert.NoError(t, err)

		mock.AssertExpectationsForObjects(t, rpm, pp)
	})

	T.Run("with error providing data changes producer", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := logging.NewNoopLogger()

		cfg := &Config{}
		msgCfg := &msgconfig.QueuesConfig{
			DataChangesTopicName: "data_changes",
		}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProvidePublisher", msgCfg.DataChangesTopicName).Return((*mockpublishers.Publisher)(nil), errors.New("blah"))

		s, err := ProvideService(
			ctx,
			logger,
			cfg,
			&textsearchcfg.Config{},
			&mocktypes.ValidMeasurementUnitDataManagerMock{},
			mockencoding.NewMockEncoderDecoder(),
			nil,
			pp,
			tracing.NewNoopTracerProvider(),
			msgCfg,
		)

		assert.Nil(t, s)
		assert.Error(t, err)

		mock.AssertExpectationsForObjects(t, pp)
	})
}