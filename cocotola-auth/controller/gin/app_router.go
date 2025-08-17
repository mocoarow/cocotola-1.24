package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/usecase"
)

type systemOwnerByOrganizationName struct {
}

func (s systemOwnerByOrganizationName) Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*mbuserservice.SystemOwner, error) {
	rsrf, err := rf.NewmoonbeamRepositoryFactory(ctx)
	if err != nil {
		return nil, err
	}
	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, rsrf)
	if err != nil {
		return nil, err
	}

	systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
	if err != nil {
		return nil, err
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
func GetPublicRouterGroupFuncs(ctx context.Context, authConfig *config.AuthConfig, txManager, nonTxManager service.TransactionManager) ([]libcontroller.InitRouterGroupFunc, error) {
	// - google
	httpClient := http.Client{
		Timeout:   time.Duration(authConfig.APITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	signingKey := []byte(authConfig.SigningKey)
	signingMethod := jwt.SigningMethodHS256
	fireabseAuthClient, err := gateway.NewFirebaseClient(ctx, authConfig.GoogleProjectID)
	if err != nil {
		return nil, err
	}
	authTokenManager := gateway.NewAuthTokenManager(ctx, fireabseAuthClient, signingKey, signingMethod, time.Duration(authConfig.AccessTokenTTLMin)*time.Minute, time.Duration(authConfig.RefreshTokenTTLHour)*time.Hour)

	googleAuthClient := gateway.NewGoogleAuthClient(&httpClient, authConfig.GoogleClientID, authConfig.GoogleClientSecret, authConfig.GoogleCallbackURL)
	googleUserUsecase := usecase.NewGoogleUser(txManager, nonTxManager, authTokenManager, googleAuthClient)
	// - authentication
	authenticationUsecase := usecase.NewAuthentication(txManager, authTokenManager, &systemOwnerByOrganizationName{})
	// - password
	passwordUsecase := usecase.NewPassword(txManager, nonTxManager, authTokenManager)

	// public router
	return []libcontroller.InitRouterGroupFunc{
		NewInitTestRouterFunc(),
		NewInitAuthRouterFunc(authenticationUsecase),
		NewInitGoogleRouterFunc(googleUserUsecase),
		NewInitPasswordRouterFunc(passwordUsecase),
	}, nil
}

func GetPrivateRouterGroupFuncs(txManager, nonTxManager service.TransactionManager) []libcontroller.InitRouterGroupFunc {
	// - rbac
	rbacUsecase := usecase.NewRBACUsecase(txManager, nonTxManager)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitRBACRouterFunc(rbacUsecase),
	}
}

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
