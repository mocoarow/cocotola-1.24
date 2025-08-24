package gateway

import (
	"embed"
	"log/slog"
	"os"

	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

var testDBFile string

func openSQLiteForTest() (*gorm.DB, error) {
	return libgateway.OpenSQLite3(&libgateway.SQLite3Config{
		File: testDBFile,
	}, slog.LevelInfo, "test")
}

// func OpenSQLiteInMemory(sqlFS embed.FS) (*gorm.DB, error) {
// 	logger := slog.Default()
// 	db, err := gorm.Open(gormSQLite.Open("file:memdb1?mode=memory&cache=shared"), &gorm.Config{
// 		Logger: slog_gorm.New(
// 			slog_gorm.WithLogger(logger), // Optional, use slog.Default() by default
// 			slog_gorm.WithTraceAll(),     // trace all messages
// 		),
// 	})
// 	if err != nil {
// 		return nil, liberrors.Errorf("gorm.Open. err: %w", err)
// 	}
// 	if err := setupSQLite(sqlFS, db); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// func setupSQLite(sqlFS embed.FS, db *gorm.DB) error {
// 	driverName := "sqlite3"
// 	sourceDriver, err := iofs.New(sqlFS, driverName)
// 	if err != nil {
// 		return err
// 	}
// 	return setupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
// 		return migrate_sqlite3.WithInstance(sqlDB, &migrate_sqlite3.Config{})
// 	})
// }

func InitSQLiteInFile(fs embed.FS) (*gorm.DB, error) {
	testDBFile = "./test.db"
	os.Remove(testDBFile)
	db, err := openSQLiteForTest()
	if err != nil {
		return nil, liberrors.Errorf("openSQLiteForTest. err: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, liberrors.Errorf("db.DB. err: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, liberrors.Errorf("sqlDB.ping. err: %w", err)
	}

	if err := libgateway.MigrateSQLite3DB(db, fs); err != nil {
		return nil, liberrors.Errorf("migrate. err: %w", err)
	}

	return db, nil
}
