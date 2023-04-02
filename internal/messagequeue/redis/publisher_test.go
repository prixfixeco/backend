package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/logging/zerolog"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	testutils "github.com/prixfixeco/backend/tests/utils"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockMessagePublisher struct {
	mock.Mock
}

func (m *mockMessagePublisher) Publish(ctx context.Context, channel string, message any) *redis.IntCmd {
	return m.Called(ctx, channel, message).Get(0).(*redis.IntCmd)
}

func Test_redisPublisher_Publish(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		logger := zerolog.NewZerologLogger(logging.DebugLevel)

		cfg := Config{
			QueueAddresses: []string{t.Name()},
		}
		provider := ProvideRedisPublisherProvider(logger, tracing.NewNoopTracerProvider(), cfg)
		require.NotNil(t, provider)

		a, err := provider.ProviderPublisher(t.Name())
		assert.NotNil(t, a)
		assert.NoError(t, err)

		actual, ok := a.(*redisPublisher)
		require.True(t, ok)

		ctx := context.Background()
		inputData := &struct {
			Name string `json:"name"`
		}{
			Name: t.Name(),
		}

		mmp := &mockMessagePublisher{}
		mmp.On(
			"Publish",
			testutils.ContextMatcher,
			actual.topic,
			[]byte(fmt.Sprintf(`{"name":%q}%s`, t.Name(), string(byte(10)))),
		).Return(&redis.IntCmd{})

		actual.publisher = mmp

		err = actual.Publish(ctx, inputData)
		assert.NoError(t, err)

		mock.AssertExpectationsForObjects(t, mmp)
	})

	T.Run("with error encoding value", func(t *testing.T) {
		t.Parallel()

		logger := zerolog.NewZerologLogger(logging.DebugLevel)

		cfg := Config{
			QueueAddresses: []string{t.Name()},
		}
		provider := ProvideRedisPublisherProvider(logger, tracing.NewNoopTracerProvider(), cfg)
		require.NotNil(t, provider)

		a, err := provider.ProviderPublisher(t.Name())
		assert.NotNil(t, a)
		assert.NoError(t, err)

		actual, ok := a.(*redisPublisher)
		require.True(t, ok)

		ctx := context.Background()
		inputData := &struct {
			Name json.Number `json:"name"`
		}{
			Name: json.Number(t.Name()),
		}

		err = actual.Publish(ctx, inputData)
		assert.Error(t, err)
	})
}

func TestProvideRedisPublisherProvider(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		logger := zerolog.NewZerologLogger(logging.DebugLevel)

		cfg := Config{
			QueueAddresses: []string{t.Name()},
		}
		actual := ProvideRedisPublisherProvider(logger, tracing.NewNoopTracerProvider(), cfg)
		assert.NotNil(t, actual)
	})
}

func Test_publisherProvider_ProviderPublisher(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		logger := zerolog.NewZerologLogger(logging.DebugLevel)

		cfg := Config{
			QueueAddresses: []string{t.Name()},
		}
		provider := ProvideRedisPublisherProvider(logger, tracing.NewNoopTracerProvider(), cfg)
		require.NotNil(t, provider)

		actual, err := provider.ProviderPublisher(t.Name())
		assert.NotNil(t, actual)
		assert.NoError(t, err)
	})

	T.Run("with cache hit", func(t *testing.T) {
		t.Parallel()

		logger := zerolog.NewZerologLogger(logging.DebugLevel)

		cfg := Config{
			QueueAddresses: []string{t.Name()},
		}
		provider := ProvideRedisPublisherProvider(logger, tracing.NewNoopTracerProvider(), cfg)
		require.NotNil(t, provider)

		actual, err := provider.ProviderPublisher(t.Name())
		assert.NotNil(t, actual)
		assert.NoError(t, err)

		actual, err = provider.ProviderPublisher(t.Name())
		assert.NotNil(t, actual)
		assert.NoError(t, err)
	})
}
