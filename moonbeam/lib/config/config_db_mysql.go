package config

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

func initDBMySQL(ctx context.Context, cfg *DBConfig, level slog.Level, fs fs.FS, appName string) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	return libgateway.InitMySQL(ctx, cfg.MySQL, cfg.Migration, level, fs, appName)
}
