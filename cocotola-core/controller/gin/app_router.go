package controller

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"gorm.io/gorm"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/config"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/middleware"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/usecase"

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

func GetPublicRouterGroupFuncs() []libcontroller.InitRouterGroupFunc {
	// public router
	return []libcontroller.InitRouterGroupFunc{
		// controller.NewInitTestRouterFunc(),
	}
}

func GetBearerTokenPrivateRouterGroupFuncs(ctx context.Context, coreConfig *config.CoreConfig, db *gorm.DB, txManager, nonTxManager service.TransactionManager, rbacClient service.CocotolaRBACClient) ([]libcontroller.InitRouterGroupFunc, error) {
	// - workbookQueryUsecase
	deckQueryUsecase := resourcemanagergateway.NewDeckQueryUsecase(db)
	// - workbookCommandUsecase
	deckCommandUsecase := resourcemanager.NewDeckCommandUsecase(txManager, nonTxManager, rbacClient)

	// - profileUsecase
	profileUsecase := usecase.NewProfileUsecase(nonTxManager)

	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitDeckRouterFunc(deckQueryUsecase, deckCommandUsecase),
		NewInitProfileRouterFunc(profileUsecase),
	}, nil
}

func GetBasicPrivateRouterGroupFuncs(ctx context.Context, coreConfig *config.CoreConfig, db *gorm.DB, txManager, nonTxManager service.TransactionManager, rbacClient service.CocotolaRBACClient) ([]libcontroller.InitRouterGroupFunc, error) {
	callbackUsecase := usecase.NewCallback(txManager, nonTxManager, rbacClient)
	// private router
	return []libcontroller.InitRouterGroupFunc{
		NewInitCallbackRouterFunc(callbackUsecase),
	}, nil
}

func InitBearerTokenAuthMiddleware(authClientConfig *config.AuthAPIClientConfig) (gin.HandlerFunc, error) {
	// middleware
	httpClient := http.Client{
		Timeout:   time.Duration(authClientConfig.TimeoutSec) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	authEndpoint, err := url.Parse(authClientConfig.Endpoint)
	if err != nil {
		return nil, err
	}
	cocotolaAuthClient := gateway.NewCocotolaAuthClient(&httpClient, authEndpoint)
	return middleware.NewAuthMiddleware(cocotolaAuthClient), nil
}
