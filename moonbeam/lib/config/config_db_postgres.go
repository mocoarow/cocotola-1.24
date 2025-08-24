package config

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

func initDBPostgres(ctx context.Context, cfg *DBConfig, logLevel slog.Level, fs fs.FS, appName string) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	return libgateway.InitPostgres(ctx, cfg.Postgres, cfg.Migration, logLevel, fs, appName) //nolint:wrapcheck
}
