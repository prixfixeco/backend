package capitalism

import (
	"context"
	"testing"

	capitalismmock "github.com/dinnerdonebetter/backend/internal/capitalism/mock"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	mockrouting "github.com/dinnerdonebetter/backend/internal/routing/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildTestService() *service {
	return &service{
		logger:         logging.NewNoopLogger(),
		tracer:         tracing.NewTracerForTest("test"),
		paymentManager: capitalismmock.NewMockPaymentManager(),
	}
}

func TestProvideValidInstrumentsService(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := logging.NewNoopLogger()

		rpm := mockrouting.NewRouteParamManager()
		mpm := capitalismmock.NewMockPaymentManager()

		s := ProvideService(
			ctx,
			logger,
			tracing.NewNoopTracerProvider(),
			mpm,
		)
		assert.NotNil(t, s)

		mock.AssertExpectationsForObjects(t, rpm)
	})
}