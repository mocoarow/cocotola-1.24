package controller

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libgateway "github.com/mocoarow/cocotola-1.24/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/config"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/middleware"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/usecase"

	guestgatgeway "github.com/mocoarow/cocotola-1.24/cocotola-core/gateway/guest"
	resourcemanagergateway "github.com/mocoarow/cocotola-1.24/cocotola-core/gateway/resource_manager"
	resourcemanager "github.com/mocoarow/cocotola-1.24/cocotola-core/usecase/resource_manager"
)

// type NewIteratorFunc func(ctx context.Context, workbookID appD.WorkbookID, problemType appD.ProblemTypeName, reader io.Reader) (appS.ProblemAddParameterIterator, error)

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

func GetPublicRouterGroupFuncs(_ context.Context, db *gorm.DB) []libcontroller.InitRouterGroupFunc {
	cardQueryUsecase := guestgatgeway.NewCardQueryUsecase(db)
	// public router
	return []libcontroller.InitRouterGroupFunc{
		// controller.NewInitTestRouterFunc(),
		NewInitCardRouterFunc(cardQueryUsecase),
	}
}

func GetBearerTokenPrivateRouterGroupFuncs(_ context.Context, db *gorm.DB, txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) ([]libcontroller.InitRouterGroupFunc, error) {
	// - workbookQueryUsecase
	guestDeckQueryUsecase := guestgatgeway.NewDeckQueryUsecase(db, rbacClient)
	studentDeckQueryUsecase := resourcemanagergateway.NewDeckQueryUsecase(db)
	// - workbookCommandUsecase
	deckCommandUsecase := resourcemanager.NewDeckCommandUsecase(txManager, nonTxManager, rbacClient)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitDeckRouterFunc(guestDeckQueryUsecase, studentDeckQueryUsecase, deckCommandUsecase),
		// NewInitProfileRouterFunc(profileUsecase),
	}, nil
}

func GetBasicPrivateRouterGroupFuncs(_ context.Context, txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) ([]libcontroller.InitRouterGroupFunc, error) {
	callbackUsecase := usecase.NewCallback(txManager, nonTxManager, rbacClient)
	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitCallbackRouterFunc(callbackUsecase),
	}, nil
}

func InitBearerTokenAuthMiddleware(authClientConfig *config.AuthAPIClientConfig) (gin.HandlerFunc, error) {
	// middleware
	httpClient := http.Client{ //nolint:exhaustruct
		Timeout:   time.Duration(authClientConfig.TimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	authEndpoint, err := url.Parse(authClientConfig.Endpoint)
	if err != nil {
		return nil, mbliberrors.Errorf("Parse: %w", err)
	}
	cocotolaAuthClient := libgateway.NewCocotolaAuthClient(&httpClient, authEndpoint)

	return middleware.NewAuthMiddleware(cocotolaAuthClient), nil
}
