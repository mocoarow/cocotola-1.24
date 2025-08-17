package config

import (
	"context"
	"io"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracerExporterNone(ctx context.Context, traceConfig *TraceConfig) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(io.Discard),
	)
}
