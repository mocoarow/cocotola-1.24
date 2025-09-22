package initialize

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gorm.io/gorm"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	libgateway "github.com/mocoarow/cocotola-1.24/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/middleware"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type operator struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
}

func (o *operator) GetUserID() *mbuserdomain.UserID {
	return o.userID
}
func (o *operator) GetOrganizationID() *mbuserdomain.OrganizationID {
	return o.organizationID
}

type AuthInitParameter struct {
	OrganizationID       *mbuserdomain.OrganizationID
	SystemOwnerID        *mbuserdomain.UserID
	GuestID              *mbuserdomain.UserID
	PublicDefaultSpaceID *mbuserdomain.SpaceID
}

func Initialize(ctx context.Context, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, coreConfig *config.CoreConfig, authInitParam *AuthInitParameter) (*domain.FolderID, []*domain.DeckID, error) {
	ctx, span := tracer.Start(ctx, "Initialize")
	defer span.End()

	txManager, err := initApp(ctx, parent, dialect, driverName, db, logConfig, coreConfig)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("initApp: %w", err)
	}

	rootFolderID, err := initRootFolder(ctx, txManager, authInitParam.OrganizationID, authInitParam.SystemOwnerID, authInitParam.PublicDefaultSpaceID)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("initApp2: %w", err)
	}

	deckIDs, err := initEnglishWord(ctx, txManager, authInitParam.OrganizationID, authInitParam.SystemOwnerID, authInitParam.PublicDefaultSpaceID, rootFolderID)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("initEnglishWord: %w", err)
	}

	return rootFolderID, deckIDs, nil
}

func initApp(ctx context.Context, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, coreConfig *config.CoreConfig) (service.TransactionManager, error) {
	ctx, span := tracer.Start(ctx, "initApp")
	defer span.End()

	// - rbacClient
	rbacClient := initCocotolaRBACClient(coreConfig.AuthAPIClient)
	authClient := initCocotolaAuthClient(coreConfig.AuthAPIClient)

	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, rbacClient)
	}
	rf, err := rff(ctx, db)
	if err != nil {
		return nil, mbliberrors.Errorf("rff: %w", err)
	}

	// init transaction manager
	txManager := initTransactionManager(db, rff)

	// init non transaction manager
	nonTxManager := initNonTransactionManager(rf)

	// init auth middleware
	bearerTokenAuthMiddleware, err := controller.InitBearerTokenAuthMiddleware(coreConfig.AuthAPIClient)
	if err != nil {
		return nil, mbliberrors.Errorf("InitBearerTokenAuthMiddleware: %w", err)
	}

	// init guest middleware
	// TODO:
	guestMiddleware := middleware.NewGuestMiddleware(1, 1)

	basicAuthMiddleware := gin.BasicAuth(gin.Accounts{
		coreConfig.CoreAPIServer.Username: coreConfig.CoreAPIServer.Password,
	})

	// init public and private router group functions
	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs(ctx, db)

	bearerTokenPrivateRouterGroupFuncs, err := controller.GetBearerTokenPrivateRouterGroupFuncs(ctx, db, txManager, nonTxManager, rbacClient, authClient)
	if err != nil {
		return nil, mbliberrors.Errorf("GetBearerTokenPrivateRouterGroupFuncs: %w", err)
	}

	basicPrivateRouterGroupFuncs, err := controller.GetBasicPrivateRouterGroupFuncs(ctx, txManager, nonTxManager, rbacClient)
	if err != nil {
		return nil, mbliberrors.Errorf("GetBasicPrivateRouterGroupFuncs: %w", err)
	}

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, parent, domain.AppName, logConfig)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs, guestMiddleware)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, bearerTokenAuthMiddleware, bearerTokenPrivateRouterGroupFuncs)

	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, basicAuthMiddleware, basicPrivateRouterGroupFuncs)

	return txManager, nil
}

func initCocotolaRBACClient(authAPIClientConfig *config.AuthAPIClientConfig) libapi.CocotolaRBACClient {
	httpClient := http.Client{ //nolint:exhaustruct
		Timeout:   time.Duration(authAPIClientConfig.TimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	authAPIEndpoint, err := url.Parse(authAPIClientConfig.Endpoint)
	if err != nil {
		libdomain.CheckError(err)
	}

	rbacClient := libgateway.NewCocotolaRBACClient(&httpClient, authAPIEndpoint, authAPIClientConfig.Username, authAPIClientConfig.Password)
	return rbacClient
}

func initCocotolaAuthClient(authAPIClientConfig *config.AuthAPIClientConfig) libapi.CocotolaAuthClient {
	httpClient := http.Client{ //nolint:exhaustruct
		Timeout:   time.Duration(authAPIClientConfig.TimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	authAPIEndpoint, err := url.Parse(authAPIClientConfig.Endpoint)
	if err != nil {
		libdomain.CheckError(err)
	}

	authClient := libgateway.NewCocotolaAuthClient(&httpClient, authAPIEndpoint)
	return authClient
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
