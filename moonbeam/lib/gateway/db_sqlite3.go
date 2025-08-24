package gateway

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	gorm_sqlite "github.com/glebarez/sqlite"
	"github.com/golang-migrate/migrate/v4/database"

	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	migrate_sqlite3 "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway/sqlite"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
)

type DialectSQLite3 struct {
}

func (d *DialectSQLite3) Name() string {
	return "sqlite3"
}

func (d *DialectSQLite3) BoolDefaultValue() string {
	return "0"
}

type SQLite3Config struct {
	File string `yaml:"file" validate:"required"`
}

func OpenSQLite3(cfg *SQLite3Config, logLevel slog.Level, appName string) (*gorm.DB, error) {
	gormDialector := gorm_sqlite.Open(cfg.File)

	gormConfig := gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithTraceAll(), // trace all messages
			slog_gorm.WithContextFunc(liblog.LoggerNameKey, func(_ context.Context) (slog.Value, bool) {
				return slog.StringValue(appName + "-gorm"), true
			}),
			slog_gorm.SetLogLevel(slog_gorm.DefaultLogType, logLevel),
		),
	}

	return gorm.Open(gormDialector, &gormConfig)
}

func MigrateSQLite3DB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "sqlite3"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return liberrors.Errorf("iofs.New: %w", err)
	}

	var _ = sourceDriver

	return MigrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_sqlite3.WithInstance(sqlDB, &migrate_sqlite3.Config{})
	})
}

func InitSqlite3(ctx context.Context, cfg *SQLite3Config, migration bool, logLevel slog.Level, fs fs.FS, appName string) (DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := OpenSQLite3(cfg, logLevel, appName)
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("OpenSQLite file(%s): %w", cfg.File, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("DB. file(%s): %w", cfg.File, err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, nil, nil, liberrors.Errorf("Ping. file(%s): %w", cfg.File, err)
	}

	if migration {
		if err := MigrateSQLite3DB(db, fs); err != nil {
			return nil, nil, nil, liberrors.Errorf("migrate DB. file: %s err: %w", cfg.File, err)
		}
	}

	dialect := DialectSQLite3{}

	return &dialect, db, sqlDB, nil
}
