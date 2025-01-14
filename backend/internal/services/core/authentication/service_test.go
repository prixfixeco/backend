package authentication

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/analytics"
	"github.com/dinnerdonebetter/backend/internal/authentication/mock"
	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	"github.com/dinnerdonebetter/backend/internal/featureflags"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/messagequeue/mock"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/metrics"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/pkg/testutils"
	"github.com/dinnerdonebetter/backend/internal/routing/mock"
	"github.com/dinnerdonebetter/backend/pkg/types/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric"
)

func buildTestService(t *testing.T) *service {
	t.Helper()

	logger := logging.NewNoopLogger()
	encoderDecoder := encoding.ProvideServerEncoderDecoder(logger, tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

	cfg := &Config{
		JWTSigningKey: base64.URLEncoding.EncodeToString([]byte(testutils.Example32ByteKey)),
	}
	queueCfg := &msgconfig.QueuesConfig{DataChangesTopicName: "data_changes"}

	pp := &mockpublishers.ProducerProvider{}
	pp.On("ProvidePublisher", queueCfg.DataChangesTopicName).Return(&mockpublishers.Publisher{}, nil)

	rpm := mockrouting.NewRouteParamManager()
	rpm.On(
		"BuildRouteParamStringIDFetcher",
		AuthProviderParamKey,
	).Return(func(*http.Request) string { return "" })

	mmp := &metrics.MockProvider{}
	mmp.On("NewInt64Counter", rejectedRequestCounterName, []metric.Int64CounterOption(nil)).Return(
		metrics.Int64CounterForTest(t.Name()), nil,
	)

	s, err := ProvideService(
		logger,
		cfg,
		&mockauthn.Authenticator{},
		database.NewMockDatabase(),
		&mocktypes.HouseholdUserMembershipDataManagerMock{},
		encoderDecoder,
		tracing.NewNoopTracerProvider(),
		pp,
		&featureflags.NoopFeatureFlagManager{},
		analytics.NewNoopEventReporter(),
		rpm,
		mmp,
		queueCfg,
	)
	require.NoError(t, err)

	mock.AssertExpectationsForObjects(t, pp, rpm, mmp)

	return s.(*service)
}

func TestProvideService(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()
		logger := logging.NewNoopLogger()
		encoderDecoder := encoding.ProvideServerEncoderDecoder(logger, tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		cfg := &Config{
			JWTSigningKey: base64.URLEncoding.EncodeToString([]byte(testutils.Example32ByteKey)),
		}
		queueCfg := &msgconfig.QueuesConfig{DataChangesTopicName: "data_changes"}

		pp := &mockpublishers.ProducerProvider{}
		pp.On("ProvidePublisher", queueCfg.DataChangesTopicName).Return(&mockpublishers.Publisher{}, nil)

		rpm := mockrouting.NewRouteParamManager()
		rpm.On(
			"BuildRouteParamStringIDFetcher",
			AuthProviderParamKey,
		).Return(func(*http.Request) string { return "" })

		s, err := ProvideService(
			logger,
			cfg,
			&mockauthn.Authenticator{},
			database.NewMockDatabase(),
			&mocktypes.HouseholdUserMembershipDataManagerMock{},
			encoderDecoder,
			tracing.NewNoopTracerProvider(),
			pp,
			&featureflags.NoopFeatureFlagManager{},
			analytics.NewNoopEventReporter(),
			rpm,
			metrics.NewNoopMetricsProvider(),
			queueCfg,
		)

		assert.NotNil(t, s)
		assert.NoError(t, err)
	})
}
