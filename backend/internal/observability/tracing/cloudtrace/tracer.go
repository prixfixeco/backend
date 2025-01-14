package cloudtrace

import (
	"context"
	"fmt"

	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

type errorHandler struct {
	logger logging.Logger
}

func (h errorHandler) Handle(err error) {
	h.logger.Error("tracer reported issue", err)
}

func init() {
	otel.SetErrorHandler(errorHandler{logger: logging.NewNoopLogger().WithName("otel_errors")})
}

// SetupCloudTrace creates a new trace provider instance and registers it as global trace provider.
func SetupCloudTrace(ctx context.Context, serviceName string, spanCollectionProbability float64, cfg *Config) (tracing.TracerProvider, error) {
	exporter, err := texporter.New(texporter.WithProjectID(cfg.ProjectID))
	if err != nil {
		return nil, fmt.Errorf("setting up trace exporter: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithOSType(),
		resource.WithAttributes(
			attribute.KeyValue{
				Key:   semconv.ServiceNameKey,
				Value: attribute.StringValue(serviceName),
			},
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(spanCollectionProbability)),
	)

	otel.SetTracerProvider(tp)

	return tp, nil
}
