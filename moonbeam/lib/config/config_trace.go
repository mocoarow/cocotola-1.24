package config

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type OTLPConfig struct {
	Endpoint string `yaml:"endpoint" validate:"required"`
	Insecure bool   `yaml:"insecure"`
}
type GoogleTraceConfig struct {
	ProjectID string `yaml:"projectID" validate:"required"`
}

type TraceConfig struct {
	Exporter string             `yaml:"exporter" validate:"required"`
	OTLP     *OTLPConfig        `yaml:"otlp"`
	Google   *GoogleTraceConfig `yaml:"google"`
}

func initTracerExporter(ctx context.Context, traceConfig *TraceConfig) (sdktrace.SpanExporter, error) {
	initTracerExporter, ok := initTracerExporters[traceConfig.Exporter]
	if !ok {
		return nil, liberrors.Errorf("invalid exporter: %s. err: %w", traceConfig.Exporter, libdomain.ErrInvalidArgument)
	}

	return initTracerExporter(ctx, traceConfig)
}

func InitTracerProvider(ctx context.Context, appName string, traceConfig *TraceConfig) (*sdktrace.TracerProvider, error) {
	exp, err := initTracerExporter(ctx, traceConfig)
	if err != nil {
		return nil, liberrors.Errorf("initTracerExporter. err: %w", err)
		// return nil, errors.Wrap(err, "initTracerExporter")
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// Record information about this application in a Resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
		)),
	)

	return tp, nil
}
