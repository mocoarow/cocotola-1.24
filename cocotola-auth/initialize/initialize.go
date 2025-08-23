package initialize

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func Initialize(ctx context.Context, systemToken libdomain.SystemToken, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, authConfig *config.AuthConfig) error {
	appUserEventHandler := mblibservice.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {

		},
	}
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, appUserEventHandler) // nolint:wrapcheck
	}
	rf, err := rff(ctx, db)
	if err != nil {
		return err
	}

	// init transaction manager
	txManager, err := mblibgateway.NewTransactionManagerT(db, rff)
	if err != nil {
		return err
	}

	// init non transaction manager
	nonTxManager, err := mblibgateway.NewNonTransactionManagerT(rf)
	if err != nil {
		return err
	}

	// init auth token manager
	authTokenManager, err := controller.NewAuthTokenManager(ctx, authConfig)
	if err != nil {
		return err
	}

	// init auth middleware
	bearerTokenAuthMiddleware, err := controller.InitBearerTokenAuthMiddleware(systemToken, authTokenManager, nonTxManager)
	if err != nil {
		return err
	}
	basicAuthMiddleware := gin.BasicAuth(gin.Accounts{
		authConfig.AuthAPIServer.Username: authConfig.AuthAPIServer.Password,
	})

	// init public and private router group functions
	publicRouterGroupFuncs, err := controller.GetPublicRouterGroupFuncs(ctx, systemToken, authConfig, txManager, nonTxManager, authTokenManager)
	if err != nil {
		return err
	}

	bearerTokenPrivateRouterGroupFuncs := controller.GetBasicPrivateRouterGroupFuncs(ctx, txManager, nonTxManager)
	basicPrivateRouterGroupFuncs := controller.GetBearerTokenPrivateRouterGroupFuncs(ctx, txManager, nonTxManager)

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, parent, domain.AppName, logConfig)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, bearerTokenAuthMiddleware, bearerTokenPrivateRouterGroupFuncs)

	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, basicAuthMiddleware, basicPrivateRouterGroupFuncs)

	initApp1(ctx, txManager, nonTxManager, "cocotola", authConfig.OwnerLoginID, authConfig.OwnerPassword)

	return nil
}

// func initAPIServer(ctx context.Context, root gin.IRouter, appName string, logConfig *mblibconfig.LogConfig, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
// 	// api
// 	api := libcontroller.InitAPIRouterGroup(ctx, root, appName, logConfig)

// 	// v1
// 	v1 := api.Group("v1")

// 	// public router
// 	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

// 	// private router
// 	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, privateRouterGroupFuncs)
// }

func initApp1(ctx context.Context, txManager, nonTxManager service.TransactionManager, organizationName, loginID, password string) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp1"))

	addOrganizationFunc := func(ctx context.Context, systemAdmin *mbuserservice.SystemAdmin) error {
		organization, err := systemAdmin.FindOrganizationByName(ctx, organizationName)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("organization: %d", organization.OrganizationID().Int()))
			return nil
		} else if !errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
			return mbliberrors.Errorf("failed to FindOrganizationByName. err: %w", err)
		}

		firstOwnerAddParam, err := mbuserservice.NewAppUserAddParameter(loginID, "Owner(cocotola)", password, "", "", "", "")
		if err != nil {
			return mbliberrors.Errorf("NewFirstOwnerAddParameter. err: %w", err)
		}

		organizationAddParameter, err := mbuserservice.NewOrganizationAddParameter(organizationName, firstOwnerAddParam)
		if err != nil {
			return mbliberrors.Errorf("NewOrganizationAddParameter. err: %w", err)
		}

		organizationID, err := systemAdmin.AddOrganization(ctx, organizationAddParameter)
		if err != nil {
			return mbliberrors.Errorf("AddOrganization. err: %w", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", organizationID.Int()))
		return nil
	}

	if err := systemAdminAction(ctx, txManager, addOrganizationFunc); err != nil {
		panic(err)
	}
}

func systemAdminAction(ctx context.Context, transactionManager service.TransactionManager, fn func(context.Context, *mbuserservice.SystemAdmin) error) error {
	return transactionManager.Do(ctx, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		if err != nil {
			return mbliberrors.Errorf(". err: %w", err)
		}

		systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, rsrf)
		if err != nil {
			return mbliberrors.Errorf(". err: %w", err)
		}

		return fn(ctx, systemAdmin)
	})
}

// func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig *mblibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, publicRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
// 	// cors
// 	ginCorsConfig := mblibconfig.InitCORS(corsConfig)

// 	// root
// 	libcontroller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

// 	InitApiServer(ctx, rootRouterGroup, appName, publicRouterGroupFuncs)
// }
