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
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin/private"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin/public"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/usecase"
)

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

func GetPublicRouterGroupFuncs(_ context.Context, systemToken libdomain.SystemToken, authConfig *config.AuthConfig, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) ([]libcontroller.InitRouterGroupFunc, error) {
	// - google
	httpClient := http.Client{ //nolint:exhaustruct
		Timeout:   time.Duration(authConfig.GoogleAPITimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	googleAuthClient := gateway.NewGoogleAuthClient(&httpClient, authConfig.GoogleClientID, authConfig.GoogleClientSecret, authConfig.GoogleCallbackURL)
	googleUserUsecase := usecase.NewGoogleUser(systemToken, mbTxManager, mbNonTxManager, txManager, nonTxManager, authTokenManager, googleAuthClient)
	// - authentication
	authenticationUsecase := usecase.NewAuthentication(systemToken, mbTxManager, authTokenManager)
	// &systemOwnerByOrganizationName{})
	// - password
	passwordUsecase := usecase.NewPassword(systemToken, mbTxManager, mbNonTxManager, authTokenManager)
	// - guest
	guestUsecase := usecase.NewGuest(systemToken, mbTxManager, mbNonTxManager, authTokenManager)

	// public router
	return []libcontroller.InitRouterGroupFunc{
		NewInitTestRouterFunc(),
		public.NewInitAuthRouterFunc(authenticationUsecase),
		public.NewInitGoogleRouterFunc(googleUserUsecase),
		public.NewInitPasswordRouterFunc(passwordUsecase),
		public.NewInitGuestRouterFunc(guestUsecase),
	}, nil
}

func GetBasicPrivateRouterGroupFuncs(_ context.Context, systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient) []libcontroller.InitRouterGroupFunc {
	// - rbac
	rbacUsecase := usecase.NewRBACUsecase(systemToken, mbTxManager, mbNonTxManager)
	// - callback
	callbackUsecase := usecase.NewCallbackUsecase(systemToken, mbTxManager, mbNonTxManager, cocotolaCoreCallbackClient)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		private.NewInitRBACRouterFunc(rbacUsecase),
		private.NewInitCallbackRouterFunc(callbackUsecase),
	}
}

func GetBearerTokenRouterGroupFuncs(_ context.Context, systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager, mbrf mbuserservice.RepositoryFactory) []libcontroller.InitRouterGroupFunc {
	// - user
	userUsecase := usecase.NewUserUsecase(systemToken, mbTxManager, mbNonTxManager, authTokenManager)
	spaceUsecase := gateway.NewSpaceQueryUsecase(mbrf)
	profileUsecase := usecase.NewProfileUsecase(mbNonTxManager)
	return []libcontroller.InitRouterGroupFunc{
		public.NewInitUserRouterFunc(userUsecase),
		private.NewInitSpaceRouterFunc(spaceUsecase),
		private.NewInitProfileRouterFunc(profileUsecase),
		// NewInitRBACRouterFunc(rbacUsecase),
	}
}

func InitBearerTokenAuthMiddleware(systemToken libdomain.SystemToken, authTokenManager service.AuthTokenManager, mbNonTxManager mbuserservice.TransactionManager, mbrf mbuserservice.RepositoryFactory) (gin.HandlerFunc, error) {
	return middleware.NewAuthMiddleware(systemToken, authTokenManager, mbNonTxManager, mbrf), nil
}
