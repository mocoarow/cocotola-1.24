package config

import (
	"context"
	"database/sql"
	"io/fs"

	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

func initDBMySQL(ctx context.Context, cfg *DBConfig, fs fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	return libgateway.InitMySQL(ctx, cfg.MySQL, cfg.Migration, fs)
}
