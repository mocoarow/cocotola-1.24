package config

import (
	"context"

	gcpexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracerExporterGCP(ctx context.Context, traceConfig *TraceConfig) (sdktrace.SpanExporter, error) {
	return gcpexporter.New(gcpexporter.WithProjectID(traceConfig.Google.ProjectID))
}
