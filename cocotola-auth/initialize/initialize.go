package initialize

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gorm.io/gorm"

	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func newCallbackOnAddUser(cocotolaAuthCallbackClient service.CocotolaAuthCallbackClient, logger *slog.Logger) func(ctx context.Context, obj any) {
	return func(ctx context.Context, obj any) {
		param, ok := obj.(map[string]int)
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid object type: %T", obj))
			return
		}

		organizationIDInt, ok := param["organizationId"]
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid organizationId type: %T", param["organizationId"]))
			return
		}

		organizationID, err := mbuserdomain.NewOrganizationID(organizationIDInt)
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid organizationId: %v", err))
			return
		}

		userIDInt, ok := param["userId"]
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid userId type: %T", param["userId"]))
			return
		}

		userID, err := mbuserdomain.NewUserID(userIDInt)
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid userId: %v", err))
			return
		}

		logger.InfoContext(ctx, fmt.Sprintf("OnAddUser: organizationID=%d, userID=%d", organizationID.Int(), userID.Int()))
		if err := cocotolaAuthCallbackClient.OnAddUser(ctx, organizationID, userID); err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("OnAddUser: %v", err))
			return
		}
	}
}

func newCallbackOnAddUserSpace(cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient, logger *slog.Logger) func(ctx context.Context, obj any) {
	return func(ctx context.Context, obj any) {
		param, ok := obj.(map[string]int)
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid object type: %T", obj))
			return
		}

		organizationIDInt, ok := param["organizationId"]
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid organizationId type: %T", param["organizationId"]))
			return
		}

		organizationID, err := mbuserdomain.NewOrganizationID(organizationIDInt)
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid organizationId: %v", err))
			return
		}

		userIDInt, ok := param["userId"]
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid userId type: %T", param["userId"]))
			return
		}

		userID, err := mbuserdomain.NewUserID(userIDInt)
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid userId: %v", err))
			return
		}

		spaceIDInt, ok := param["spaceId"]
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid spaceID type: %T", param["spaceID"]))
			return
		}

		spaceID, err := mbuserdomain.NewSpaceID(spaceIDInt)
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid spaceID: %v", err))
			return
		}

		logger.InfoContext(ctx, fmt.Sprintf("OnAddUserSpace: organizationID=%d, userID=%d, spaceID:%d", organizationID.Int(), userID.Int(), spaceID.Int()))
		if err := cocotolaCoreCallbackClient.OnAddUserSpace(ctx, organizationID, userID, spaceID); err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("OnAddUser: %v", err))
			return
		}
	}
}

func Initialize(ctx context.Context, systemToken libdomain.SystemToken, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, authConfig *config.AuthConfig) (*mbuserdomain.OrganizationID, *mbuserdomain.UserID, *mbuserdomain.SpaceID, error) {
	txManager, nonTxManager, err := initApp(ctx, systemToken, parent, dialect, driverName, db, logConfig, authConfig)
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("initApp: %w", err)
	}

	organizationID, publicDefaultSpaceID, err := initOrganization(ctx, systemToken, txManager, nonTxManager, "cocotola", authConfig.OwnerLoginID, authConfig.OwnerPassword)
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("initApp1: %w", err)
	}

	guestID, err := initApp2(ctx, systemToken, txManager, nonTxManager, "cocotola")
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("initApp2: %w", err)
	}

	return organizationID, guestID, publicDefaultSpaceID, nil
}

func initApp(ctx context.Context, systemToken libdomain.SystemToken, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, authConfig *config.AuthConfig) (service.TransactionManager, service.TransactionManager, error) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"initApp"))

	cocotolaAuthCallbackClient := initCocotolaAuthCallbackClient(authConfig)
	cocotolaCoreCallbackClient := initCocotolaCoreCallbackClient(authConfig.CoreAPIClient)

	userEventHandler := mblibservice.ResourceEventHandlerFuncs{ //nolint:exhaustruct
		AddFunc: newCallbackOnAddUser(cocotolaAuthCallbackClient, logger),
	}
	spaceEventHandler := mblibservice.ResourceEventHandlerFuncs{ //nolint:exhaustruct
		AddFunc: newCallbackOnAddUserSpace(cocotolaCoreCallbackClient, logger),
	}
	resouceEventHandlers := map[mbuserdomain.ResourceKey]mblibservice.ResourceEventHandler{
		mbuserdomain.ResourceUser:  userEventHandler,
		mbuserdomain.RecourceSpace: spaceEventHandler,
	}

	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, resouceEventHandlers)
	}
	rf, err := rff(ctx, db)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("rff: %w", err)
	}

	mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
	}

	// init transaction manager
	txManager := initTransactionManager(db, rff)

	// init non transaction manager
	nonTxManager := initNonTransactionManager(rf)

	// init auth token manager
	authTokenManager, err := controller.NewAuthTokenManager(ctx, authConfig)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("NewAuthTokenManager: %w", err)
	}

	// init auth middleware
	bearerTokenAuthMiddleware, err := controller.InitBearerTokenAuthMiddleware(systemToken, authTokenManager, nonTxManager)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("InitBearerTokenAuthMiddleware: %w", err)
	}
	basicAuthMiddleware := gin.BasicAuth(gin.Accounts{
		authConfig.AuthAPIServer.Username: authConfig.AuthAPIServer.Password,
	})

	// init public and private router group functions
	publicRouterGroupFuncs, err := controller.GetPublicRouterGroupFuncs(ctx, systemToken, authConfig, txManager, nonTxManager, authTokenManager)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("GetPublicRouterGroupFuncs: %w", err)
	}
	bearerTokenRouterGroupFuncs := controller.GetBearerTokenRouterGroupFuncs(ctx, systemToken, txManager, nonTxManager, authTokenManager, mbrf)
	basicPrivateRouterGroupFuncs := controller.GetBasicPrivateRouterGroupFuncs(ctx, systemToken, txManager, nonTxManager, cocotolaCoreCallbackClient)

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, parent, domain.AppName, logConfig)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, bearerTokenAuthMiddleware, bearerTokenRouterGroupFuncs)

	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, basicAuthMiddleware, basicPrivateRouterGroupFuncs)

	return txManager, nonTxManager, nil
}

func initCocotolaAuthCallbackClient(authConfig *config.AuthConfig) service.CocotolaAuthCallbackClient {
	httpClient := http.Client{ //nolint:exhaustruct
		Timeout:   time.Duration(authConfig.AuthAPIClient.TimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	authAPIEndpoint, err := url.Parse(authConfig.AuthAPIClient.Endpoint)
	if err != nil {
		libdomain.CheckError(err)
	}

	cocotolaAuthCallbackClient := gateway.NewCocotolaAuthCallbackClient(&httpClient, authAPIEndpoint, authConfig.AuthAPIClient.Username, authConfig.AuthAPIClient.Password)

	return cocotolaAuthCallbackClient
}

func initCocotolaCoreCallbackClient(coreAPIClientConfig *config.CoreAPIClientConfig) service.CocotolaCoreCallbackClient {
	httpClient := http.Client{ //nolint:exhaustruct
		Timeout:   time.Duration(coreAPIClientConfig.TimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	coreAPIEndpoint, err := url.Parse(coreAPIClientConfig.Endpoint)
	if err != nil {
		libdomain.CheckError(err)
	}

	cocotolaCoreCallbackClient := gateway.NewCocotolaCoreCallbackClient(&httpClient, coreAPIEndpoint, coreAPIClientConfig.Username, coreAPIClientConfig.Password)

	return cocotolaCoreCallbackClient
}

func initTransactionManager(db *gorm.DB, rff func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)) service.TransactionManager {
	txManager, err := mblibgateway.NewTransactionManagerT(db, rff)
	if err != nil {
		libdomain.CheckError(err)
	}
	return txManager
}

func initNonTransactionManager(rf service.RepositoryFactory) service.TransactionManager {
	nonTxManager, err := mblibgateway.NewNonTransactionManagerT(rf)
	if err != nil {
		libdomain.CheckError(err)
	}
	return nonTxManager
}
