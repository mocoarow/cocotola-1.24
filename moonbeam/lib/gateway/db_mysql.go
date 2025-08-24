package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	slog_gorm "github.com/orandin/slog-gorm"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
)

type DialectMySQL struct {
}

func (d *DialectMySQL) Name() string {
	return "mysql"
}

func (d *DialectMySQL) BoolDefaultValue() string {
	return "0"
}

type MySQLConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

func OpenMySQLWithDSN(dsn string, logLevel slog.Level, appName string) (*gorm.DB, error) {
	gormDialector := gorm_mysql.Open(dsn)

	gormConfig := gorm.Config{
		Logger: slog_gorm.New(
			slog_gorm.WithTraceAll(), // trace all messages
			slog_gorm.WithContextFunc(liblog.LoggerNameKey, func(ctx context.Context) (slog.Value, bool) {
				return slog.StringValue(appName + "-gorm"), true
			}),
			slog_gorm.SetLogLevel(slog_gorm.DefaultLogType, logLevel),
		),
	}

	return gorm.Open(gormDialector, &gormConfig)
}

func OpenMySQL(cfg *MySQLConfig, logLevel slog.Level, appName string) (*gorm.DB, error) {
	c := mysql.Config{
		DBName:               cfg.Database,
		User:                 cfg.Username,
		Passwd:               cfg.Password,
		Addr:                 fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Net:                  "tcp",
		ParseTime:            true,
		MultiStatements:      true,
		Params:               map[string]string{"charset": "utf8mb4"},
		Collation:            "utf8mb4_bin",
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		MaxAllowedPacket:     64 << 20, // 64 MiB.
		Loc:                  time.UTC,
	}

	return OpenMySQLWithDSN(c.FormatDSN(), logLevel, appName)
}

func MigrateMySQLDB(db *gorm.DB, sqlFS fs.FS) error {
	driverName := "mysql"
	sourceDriver, err := iofs.New(sqlFS, driverName)
	if err != nil {
		return err
	}

	return MigrateDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
		return migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
	})
}

func InitMySQL(ctx context.Context, cfg *MySQLConfig, migration bool, logLevel slog.Level, fs fs.FS, appName string) (DialectRDBMS, *gorm.DB, *sql.DB, error) {
	db, err := OpenMySQL(cfg, logLevel, appName)
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
		if err := MigrateMySQLDB(db, fs); err != nil {
			return nil, nil, nil, liberrors.Errorf("failed to MigrateMySQLDB. err: %w", err)
		}
	}

	dialect := DialectMySQL{}
	return &dialect, db, sqlDB, nil
}
