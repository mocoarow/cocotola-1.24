package config

import (
	"context"
	"io"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracerExporterNone(_ context.Context, _ *TraceConfig) (sdktrace.SpanExporter, error) {
	return stdouttrace.New( //nolint:wrapcheck
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(io.Discard),
	)
}
