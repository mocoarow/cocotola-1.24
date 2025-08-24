package config

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracerExporterStdout(_ context.Context, _ *TraceConfig) (sdktrace.SpanExporter, error) {
	return stdouttrace.New( //nolint:wrapcheck
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(os.Stderr),
	)
}
