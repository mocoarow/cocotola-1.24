package config

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

func initDBSQLite3(ctx context.Context, cfg *DBConfig, fs fs.FS, appName string) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	return libgateway.InitSqlite3(ctx, cfg.SQLite3, cfg.Migration, fs, appName)
}
