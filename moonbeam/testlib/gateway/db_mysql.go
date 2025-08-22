package gateway

import (
	"embed"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

var testDSN string
var testDBHost string
var testDBPort int

func openMySQLForTest() (*gorm.DB, error) {
	if testDSN != "" {
		return libgateway.OpenMySQLWithDSN(testDSN, "test")
	}

	return libgateway.OpenMySQL(&libgateway.MySQLConfig{
		Username: "user",
		Password: "password",
		Database: "testdb",
		Host:     testDBHost,
		Port:     testDBPort,
	}, "test")
}

// func setupMySQL(sqlFS embed.FS, db *gorm.DB) error {
// 	driverName := "mysql"
// 	sourceDriver, err := iofs.New(sqlFS, driverName)
// 	if err != nil {
// 		return err
// 	}

// 	return setupDB(db, driverName, sourceDriver, func(sqlDB *sql.DB) (database.Driver, error) {
// 		return migrate_mysql.WithInstance(sqlDB, &migrate_mysql.Config{})
// 	})
// }

func InitMySQL(fs embed.FS, dbHost string, dbPort int) (*gorm.DB, error) {
	testDBHost = dbHost
	testDBPort = dbPort
	db, err := openMySQLForTest()
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if err := libgateway.MigrateMySQLDB(db, fs); err != nil {
		return nil, err
	}

	return db, nil
}

func InitMySQLWithDSN(fs embed.FS, dsn string) (*gorm.DB, error) {
	testDSN = dsn
	db, err := openMySQLForTest()
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if err := libgateway.MigrateMySQLDB(db, fs); err != nil {
		return nil, err
	}

	return db, nil
}
