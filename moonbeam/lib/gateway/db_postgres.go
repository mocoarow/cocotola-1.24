package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4/database"
	migrate_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
)

type DialectPostgres struct {
}

func (d *DialectPostgres) Name() string {
	return "postgres"
}

func (d *DialectPostgres) BoolDefaultValue() string {
	return "false"
}

type PostgresConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

func OpenPostgres(cfg *PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, "disable", time.UTC.String())

	gormDialector := gorm_postgres.Open(dsn)

	gormConfig := gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithTraceAll(), // trace all messages
			slog_gorm.WithContextFunc(liblog.LoggerNameKey, func(ctx context.Context) (slog.Value, bool) {
				return slog.StringValue("gorm"), true
			}),
		),
	}

	return gorm.Open(gormDialector, &gormConfig)
}

func MigratePostgresDB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "postgres"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return MigrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_postgres.WithInstance(sqlDB, &migrate_postgres.Config{})
	})
}

func InitPostgres(ctx context.Context, cfg *PostgresConfig, migration bool, fs fs.FS) (DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := OpenPostgres(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, nil, err
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, nil, nil, err
	}

	if migration {
		if err := MigratePostgresDB(db, fs); err != nil {
			return nil, nil, nil, liberrors.Errorf("failed to MigrateMySQLDB. err: %w", err)
		}
	}

	dialect := DialectPostgres{}
	return &dialect, db, sqlDB, nil
}
