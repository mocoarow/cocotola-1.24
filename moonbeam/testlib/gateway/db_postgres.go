package gateway

import (
	"embed"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

var testPostgresHost string
var testPostgresPort int

func openPostgresForTest() (*gorm.DB, error) {
	return libgateway.OpenPostgres(&libgateway.PostgresConfig{
		Username: "user",
		Password: "password",
		Host:     testPostgresHost,
		Port:     testPostgresPort,
		Database: "postgres",
	})
}

// func setupPostgres(sqlFS embed.FS, db *gorm.DB) error {
// 	driverName := "postgres"
// 	sourceDriver, err := iofs.New(sqlFS, driverName)
// 	if err != nil {
// 		return err
// 	}

// 	return setupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
// 		return migrate_postgres.WithInstance(sqlDB, &migrate_postgres.Config{})
// 	})
// }

func InitPostgres(fs embed.FS, dbHost string, dbPort int) (*gorm.DB, error) {
	testPostgresHost = dbHost
	testPostgresPort = dbPort
	db, err := openPostgresForTest()
	if err != nil {
		return nil, err
	}

	if err := libgateway.MigratePostgresDB(db, fs); err != nil {
		return nil, err
	}

	return db, nil
}
