package mealplans

import (
	"errors"
	"net/http"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/encoding"
	mockencoding "github.com/dinnerdonebetter/backend/internal/encoding/mock"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	mockpublishers "github.com/dinnerdonebetter/backend/internal/messagequeue/mock"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	mockrouting "github.com/dinnerdonebetter/backend/internal/routing/mock"
	mocktypes "github.com/dinnerdonebetter/backend/pkg/types/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildTestService() *service {
	return &service{
		logger:              logging.NewNoopLogger(),
		mealPlanDataManager: &mocktypes.MealPlanDataManagerMock{},
		mealPlanIDFetcher:   func(req *http.Request) string { return "" },
		encoderDecoder:      encoding.ProvideServerEncoderDecoder(nil, nil, encoding.ContentTypeJSON),
		tracer:              tracing.NewTracerForTest("test"),
	}
}

func TestProvideMealPlansService(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		rpm := mockrouting.NewRouteParamManager()
		rpm.On(
			"BuildRouteParamStringIDFetcher",
			MealPlanIDURIParamKey,
		).Return(func(*http.Request) string { return "" })

		msgCfg := &msgconfig.QueuesConfig{DataChangesTopicName: "data_changes"}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProvidePublisher", msgCfg.DataChangesTopicName).Return(&mockpublishers.Publisher{}, nil)

		s, err := ProvideService(
			logging.NewNoopLogger(),
			&mocktypes.MealPlanDataManagerMock{},
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

	T.Run("with error providing pre-writes producer", func(t *testing.T) {
		t.Parallel()

		msgCfg := &msgconfig.QueuesConfig{DataChangesTopicName: "data_changes"}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProvidePublisher", msgCfg.DataChangesTopicName).Return((*mockpublishers.Publisher)(nil), errors.New("blah"))

		s, err := ProvideService(
			logging.NewNoopLogger(),
			&mocktypes.MealPlanDataManagerMock{},
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
