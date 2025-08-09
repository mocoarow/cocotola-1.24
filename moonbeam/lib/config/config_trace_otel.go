package config

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracerExporterOTLP(ctx context.Context, traceConfig *TraceConfig) (sdktrace.SpanExporter, error) {
	options := make([]otlptracehttp.Option, 0)
	options = append(options, otlptracehttp.WithEndpoint(traceConfig.OTLP.Endpoint))
	if traceConfig.OTLP.Insecure {
		options = append(options, otlptracehttp.WithInsecure())
	}
	return otlptracehttp.New(ctx, options...)
}
