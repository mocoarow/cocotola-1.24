package config

import (
	"context"
	"database/sql"
	"io/fs"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

type InitTracerExporterFunc func(ctx context.Context, traceConfig *TraceConfig) (sdktrace.SpanExporter, error)

var initTracerExporters map[string]InitTracerExporterFunc

type InitDBFunc func(context.Context, *DBConfig, fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error)

var initDBs map[string]InitDBFunc

func init() {
	initTracerExporters = map[string]InitTracerExporterFunc{
		"otlp":   initTracerExporterOTLP,
		"gcp":    initTracerExporterGCP,
		"none":   initTracerExporterNone,
		"stdout": initTracerExporterStdout,
	}
	initDBs = map[string]InitDBFunc{
		"mysql":    initDBMySQL,
		"postgres": initDBPostgres,
		"sqlite3":  initDBSQLite3,
	}
}
