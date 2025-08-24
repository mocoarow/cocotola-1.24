package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin/middleware"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/usecase"
)

type systemOwnerByOrganizationName struct {
}

func (s systemOwnerByOrganizationName) Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*mbuserservice.SystemOwner, error) {
	mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
	if err != nil {
		return nil, mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
	}
	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, mbrf)
	if err != nil {
		return nil, mbliberrors.Errorf("NewSystemAdmin: %w", err)
	}

	systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("GetFindSystemOwnerByOrganizationNameUser: %w", err)
	}

	return systemOwner, nil
}

func NewInitTestRouterFunc() libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		test := parentRouterGroup.Group("test")
		for _, m := range middleware {
			test.Use(m)
		}
		test.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
}

func NewAuthTokenManager(ctx context.Context, authConfig *config.AuthConfig) (service.AuthTokenManager, error) {
	signingKey := []byte(authConfig.SigningKey)
	signingMethod := jwt.SigningMethodHS256
	fireabseAuthClient, err := gateway.NewFirebaseClient(ctx, authConfig.GoogleProjectID)
	if err != nil {
		return nil, mbliberrors.Errorf("NewFirebaseClient: %w", err)
	}
	authTokenManager := gateway.NewAuthTokenManager(ctx, fireabseAuthClient, signingKey, signingMethod, time.Duration(authConfig.AccessTokenTTLMin)*time.Minute, time.Duration(authConfig.RefreshTokenTTLHour)*time.Hour)

	return authTokenManager, nil
}

func GetPublicRouterGroupFuncs(_ context.Context, systemToken libdomain.SystemToken, authConfig *config.AuthConfig, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) ([]libcontroller.InitRouterGroupFunc, error) {
	// - google
	httpClient := http.Client{
		Timeout:   time.Duration(authConfig.GoogleAPITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	googleAuthClient := gateway.NewGoogleAuthClient(&httpClient, authConfig.GoogleClientID, authConfig.GoogleClientSecret, authConfig.GoogleCallbackURL)
	googleUserUsecase := usecase.NewGoogleUser(systemToken, txManager, nonTxManager, authTokenManager, googleAuthClient)
	// - authentication
	authenticationUsecase := usecase.NewAuthentication(systemToken, txManager, authTokenManager, &systemOwnerByOrganizationName{})
	// - password
	passwordUsecase := usecase.NewPassword(systemToken, txManager, nonTxManager, authTokenManager)

	// public router
	return []libcontroller.InitRouterGroupFunc{
		NewInitTestRouterFunc(),
		NewInitAuthRouterFunc(authenticationUsecase),
		NewInitGoogleRouterFunc(googleUserUsecase),
		NewInitPasswordRouterFunc(passwordUsecase),
	}, nil
}

func GetBasicPrivateRouterGroupFuncs(_ context.Context, txManager, nonTxManager service.TransactionManager) []libcontroller.InitRouterGroupFunc {
	// - rbac
	rbacUsecase := usecase.NewRBACUsecase(txManager, nonTxManager)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitRBACRouterFunc(rbacUsecase),
	}
}
func GetBearerTokenPrivateRouterGroupFuncs(_ context.Context, systemToken libdomain.SystemToken, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) []libcontroller.InitRouterGroupFunc {
	// - rbac
	// rbacUsecase := usecase.NewRBACUsecase(txManager, nonTxManager)
	// - user
	userUsecase := usecase.NewUserUsecase(systemToken, txManager, nonTxManager, authTokenManager)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitUserRouterFunc(userUsecase),
		// NewInitRBACRouterFunc(rbacUsecase),
	}
}

func InitBearerTokenAuthMiddleware(systemToken libdomain.SystemToken, authTokenManager service.AuthTokenManager, nonTxManager service.TransactionManager) (gin.HandlerFunc, error) {
	return middleware.NewAuthMiddleware(systemToken, authTokenManager, nonTxManager), nil
}

// func InitAuthMiddleware(authConfig *config.AuthConfig) (gin.HandlerFunc, error) {
// 	authMiddleware := gin.BasicAuth(gin.Accounts{
// 		authConfig.AuthAPIServer.Username: authConfig.AuthAPIServer.Password,
// 	})
// 	return authMiddleware, nil
// }

// func InitRootRouterGroup(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig cors.Config, debugConfig *libconfig.DebugConfig) {
// 	rootRouterGroup.Use(cors.New(corsConfig))
// 	rootRouterGroup.Use(sloggin.New(slog.Default()))

// 	if debugConfig.Wait {
// 		rootRouterGroup.Use(libmiddleware.NewWaitMiddleware())
// 	}
// }

// func InitAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, initPublicRouterFunc []libcontroller.InitRouterGroupFunc, initPrivateRouterFunc []libcontroller.InitRouterGroupFunc, appName string) error {
// 	v1 := parentRouterGroup.Group("v1")
// 	{
// 		v1.Use(otelgin.Middleware(appName))
// 		v1.Use(libmiddleware.NewTraceLogMiddleware(appName))

// 		for _, fn := range initPublicRouterFunc {
// 			if err := fn(v1); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
