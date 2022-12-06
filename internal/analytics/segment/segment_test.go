package segment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/prixfixeco/backend/internal/observability/logging"
	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/pkg/types"
	"github.com/prixfixeco/backend/pkg/types/fakes"
)

func TestNewSegmentEventReporter(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		logger := logging.NewNoopLogger()

		collector, err := NewSegmentEventReporter(logger, tracing.NewNoopTracerProvider(), t.Name())
		require.NoError(t, err)
		require.NotNil(t, collector)
	})

	T.Run("with empty API key", func(t *testing.T) {
		t.Parallel()

		logger := logging.NewNoopLogger()

		collector, err := NewSegmentEventReporter(logger, tracing.NewNoopTracerProvider(), "")
		require.Error(t, err)
		require.Nil(t, collector)
	})
}

func TestSegmentEventReporter_Close(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		logger := logging.NewNoopLogger()

		collector, err := NewSegmentEventReporter(logger, tracing.NewNoopTracerProvider(), t.Name())
		require.NoError(t, err)
		require.NotNil(t, collector)

		collector.Close()
	})
}

func TestSegmentEventReporter_AddUser(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := logging.NewNoopLogger()
		exampleUserID := fakes.BuildFakeID()
		properties := map[string]any{
			"test.name": t.Name(),
		}

		collector, err := NewSegmentEventReporter(logger, tracing.NewNoopTracerProvider(), t.Name())
		require.NoError(t, err)
		require.NotNil(t, collector)

		require.NoError(t, collector.AddUser(ctx, exampleUserID, properties))
	})
}

func TestSegmentEventReporter_EventOccurred(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := logging.NewNoopLogger()
		exampleUserID := fakes.BuildFakeID()
		properties := map[string]any{
			"test.name": t.Name(),
		}

		collector, err := NewSegmentEventReporter(logger, tracing.NewNoopTracerProvider(), t.Name())
		require.NoError(t, err)
		require.NotNil(t, collector)

		require.NoError(t, collector.EventOccurred(ctx, types.CustomerEventType(t.Name()), exampleUserID, properties))
	})
}