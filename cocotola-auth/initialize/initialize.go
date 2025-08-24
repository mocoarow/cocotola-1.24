package initialize

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func newCallbackOnAddAppUser(cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient, logger *slog.Logger) func(ctx context.Context, obj any) {
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

		appUserIDInt, ok := param["appUserId"]
		if !ok {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid appuserId type: %T", param["appuserId"]))

			return
		}

		appUserID, err := mbuserdomain.NewAppUserID(appUserIDInt)
		if err != nil {
			logger.ErrorContext(ctx, fmt.Sprintf("invalid appuserId: %v", err))

			return
		}

		go func(ctx context.Context) {
			logger.InfoContext(ctx, fmt.Sprintf("OnAddAppUser: organizationID=%d, appUserID=%d", organizationID.Int(), appUserID.Int()))
			if err := cocotolaCoreCallbackClient.OnAddAppUser(ctx, organizationID, appUserID); err != nil {
				logger.ErrorContext(ctx, fmt.Sprintf("OnAddAppUser: %v", err))
			}
		}(context.Background())
	}
}

func Initialize(ctx context.Context, systemToken libdomain.SystemToken, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, authConfig *config.AuthConfig) error {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-Initialize"))
	httpClient := &http.Client{
		Timeout: time.Duration(authConfig.CoreAPIClient.TimeoutSec) * time.Second,
	}
	coreAPIEndpoint, err := url.Parse(authConfig.CoreAPIClient.Endpoint)
	if err != nil {
		return mbliberrors.Errorf("invalid core api endpoint: %w", err)
	}
	cocotolaCoreCallbackClient := gateway.NewCocotolaCoreCallbackClient(httpClient, coreAPIEndpoint, authConfig.CoreAPIClient.Username, authConfig.CoreAPIClient.Password)
	appUserEventHandler := mblibservice.ResourceEventHandlerFuncs{
		AddFunc: newCallbackOnAddAppUser(cocotolaCoreCallbackClient, logger),
	}
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, appUserEventHandler)
	}
	rf, err := rff(ctx, db)
	if err != nil {
		return mbliberrors.Errorf("rff: %w", err)
	}

	// init transaction manager
	txManager, err := mblibgateway.NewTransactionManagerT(db, rff)
	if err != nil {
		return mbliberrors.Errorf("NewTransactionManagerT: %w", err)
	}

	// init non transaction manager
	nonTxManager, err := mblibgateway.NewNonTransactionManagerT(rf)
	if err != nil {
		return mbliberrors.Errorf("NewNonTransactionManagerT: %w", err)
	}

	// init auth token manager
	authTokenManager, err := controller.NewAuthTokenManager(ctx, authConfig)
	if err != nil {
		return mbliberrors.Errorf("NewAuthTokenManager: %w", err)
	}

	// init auth middleware
	bearerTokenAuthMiddleware, err := controller.InitBearerTokenAuthMiddleware(systemToken, authTokenManager, nonTxManager)
	if err != nil {
		return mbliberrors.Errorf("InitBearerTokenAuthMiddleware: %w", err)
	}
	basicAuthMiddleware := gin.BasicAuth(gin.Accounts{
		authConfig.AuthAPIServer.Username: authConfig.AuthAPIServer.Password,
	})

	// init public and private router group functions
	publicRouterGroupFuncs, err := controller.GetPublicRouterGroupFuncs(ctx, systemToken, authConfig, txManager, nonTxManager, authTokenManager)
	if err != nil {
		return mbliberrors.Errorf("GetPublicRouterGroupFuncs: %w", err)
	}
	bearerTokenPrivateRouterGroupFuncs := controller.GetBearerTokenPrivateRouterGroupFuncs(ctx, systemToken, txManager, nonTxManager, authTokenManager)
	basicPrivateRouterGroupFuncs := controller.GetBasicPrivateRouterGroupFuncs(ctx, txManager, nonTxManager)

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, parent, domain.AppName, logConfig)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, bearerTokenAuthMiddleware, bearerTokenPrivateRouterGroupFuncs)

	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, basicAuthMiddleware, basicPrivateRouterGroupFuncs)

	if err := initApp1(ctx, systemToken, txManager, nonTxManager, "cocotola", authConfig.OwnerLoginID, authConfig.OwnerPassword); err != nil {
		return mbliberrors.Errorf("initApp1: %w", err)
	}

	return nil
}

func addOrganization(ctx context.Context, systemAdminAction *service.SystemAdminAction, organizationName, loginID, password string) (*mbuserdomain.OrganizationID, error) {
	firstOwnerAddParam, err := mbuserservice.NewAppUserAddParameter(loginID, "Owner(cocotola)", password, "", "", "", "")
	if err != nil {
		return nil, mbliberrors.Errorf("new AppUserAddParameter: %w", err)
	}

	organizationAddParameter, err := mbuserservice.NewOrganizationAddParameter(organizationName, firstOwnerAddParam)
	if err != nil {
		return nil, mbliberrors.Errorf("new OrganizationAddParameter: %w", err)
	}

	organizationID, err := systemAdminAction.SystemAdmin.AddOrganization(ctx, organizationAddParameter)
	if err != nil {
		return nil, mbliberrors.Errorf("add organization: %w", err)
	}

	return organizationID, nil
}

func initApp1(ctx context.Context, systemToken libdomain.SystemToken, _, nonTxManager service.TransactionManager, organizationName, loginID, password string) error {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp1"))

	if err := nonTxManager.Do(ctx, func(rf service.RepositoryFactory) error {
		// 1. check whether the organization already exists
		systemAdminAction, err := service.NewSystemAdminAction(ctx, systemToken, rf)
		if err != nil {
			return mbliberrors.Errorf("new organization action: %w", err)
		}

		organization, err := systemAdminAction.SystemAdmin.FindOrganizationByName(ctx, organizationName)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("organization: %d", organization.OrganizationID().Int()))

			return nil
		} else if !errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
			return mbliberrors.Errorf("find organization by name(%s): %w", organizationName, err)
		}

		// 2. add organization
		organizationID, err := addOrganization(ctx, systemAdminAction, organizationName, loginID, password)
		if err != nil {
			return mbliberrors.Errorf("add organization: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", organizationID.Int()))

		// 3. add policy to "first-owner" user
		systemOwnerAction, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
			service.WithOrganizationByName(organizationName),
			service.WithAuthorizationManager(),
		)
		if err != nil {
			return mbliberrors.Errorf("new system owner action: %w", err)
		}

		firstOwner, err := systemOwnerAction.SystemOwner.FindAppUserByLoginID(ctx, loginID)
		if err != nil {
			return mbliberrors.Errorf("FindAppUserByLoginID: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("firstOwner: %d", firstOwner.AppUserID().Int()))

		// first owner can create app users
		subject := firstOwner.AppUserID().GetRBACSubject()
		action := mbuserdomain.NewRBACAction("CreateAppUser")
		object := mbuserdomain.NewRBACObject("*")
		effect := mbuserservice.RBACAllowEffect

		if err := systemOwnerAction.AuthorizationManager.AddPolicyToUserBySystemOwner(ctx, systemOwnerAction.SystemOwner, subject, action, object, effect); err != nil {
			return mbliberrors.Errorf("AddPolicyToUserBySystemOwner: %w", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", organizationID.Int()))

		return nil
	}); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
