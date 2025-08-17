//go:build medium

package gateway_test

import (
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/mocoarow/cocotola-1.24/moonbeam/sqls"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	testlibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/testlib/gateway"
)

var (
	loc = time.UTC
)
var (
	invalidOrgID     *domain.OrganizationID
	invalidAppUserID *domain.AppUserID
)

// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }

// func createMySQLContainer() string {
// 	ctx := context.Background()
// 	mysqlContainer, err := tc_mysql.Run(ctx,
// 		"mysql:8.0.36",
// 		// tc_mysql.WithConfigFile(filepath.Join("testdata", "my_8.cnf")),
// 		tc_mysql.WithUsername("user"),
// 		tc_mysql.WithPassword("password"),
// 		tc_mysql.WithDatabase("testdb"),
// 		// tc_mysql.WithScripts(filepath.Join("testdata", "schema.sql")),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	connectionString, err := mysqlContainer.ConnectionString(ctx, "collation=utf8mb4_bin", "multiStatements=true", "parseTime=true", "charset=utf8mb4")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return connectionString
// }

func init() {
	invalidOrgIDTmp, err := domain.NewOrganizationID(99999)
	if err != nil {
		panic(err)
	}
	invalidOrgID = invalidOrgIDTmp

	invalidAppUserIDTmp, err := domain.NewAppUserID(99999)
	if err != nil {
		panic(err)
	}
	invalidAppUserID = invalidAppUserIDTmp

	// ctx := context.Background()
	// mysqlHost := getEnv("MYSQL_HOST", "127.0.0.1")
	// mysqlPortS := getEnv("MYSQL_PORT", "3307")
	// mysqlPort, err := strconv.Atoi(mysqlPortS)
	// if err != nil {
	// 	panic(err)
	// }

	// postgresHost := getEnv("POSTGRES_HOST", "127.0.0.1")
	// postgresPortS := getEnv("POSTGRES_PORT", "5433")
	// postgresPort, err := strconv.Atoi(postgresPortS)
	// if err != nil {
	// 	panic(err)
	// }
	fns := []func() (*gorm.DB, error){
		// func() (*gorm.DB, error) {
		// 	return testlibgateway.InitMySQLWithDSN(sqls.SQL, connectionString)
		// },
		// func() (*gorm.DB, error) {
		// 	return testlibgateway.InitMySQL(sqls.SQL, mysqlHost, mysqlPort)
		// },
		// func() (*gorm.DB, error) {
		// 	return testlibgateway.InitPostgres(sqls.SQL, postgresHost, postgresPort)
		// },
		func() (*gorm.DB, error) {
			return testlibgateway.InitSQLiteInFile(sqls.SQL)
		},
	}

	for _, fn := range fns {
		db, err := fn()
		if err != nil {
			log.Fatalf("failed to initialize db: %+v", err)
			panic(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}

	// for dialect, db := range testlibgateway.ListDB() {
	// 	dialect := dialect
	// 	db := db
	// rf, err := gateway.NewRepositoryFactory(ctx, dialect, dialect.Name(), db, loc)
	// if err != nil {
	// 	panic(err)
	// }
	// authorizationManager, err := rf.NewAuthorizationManager(ctx)
	// if err := authorizationManager.Init(ctx); err != nil {
	// 	panic(err)
	// }
	// }
}
