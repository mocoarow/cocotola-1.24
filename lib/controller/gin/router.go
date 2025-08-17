package controller

import (
	"context"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/mocoarow/cocotola-1.24/lib/config"
	"github.com/mocoarow/cocotola-1.24/lib/controller/gin/middleware"
	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
)

type InitRouterGroupFunc func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc)

func InitRootRouterGroup(ctx context.Context, corsConfig *mblibconfig.CORSConfig, debugConfig *config.DebugConfig) *gin.Engine {
	if !debugConfig.Gin {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// cors
	ginCorsConfig := mblibconfig.InitCORS(corsConfig)

	router.Use(gin.Recovery())
	router.Use(cors.New(ginCorsConfig))
	router.Use(sloggin.New(slog.Default()))

	if debugConfig.Wait {
		router.Use(middleware.NewWaitMiddleware())
	}

	return router
}

func InitAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, appName string) *gin.RouterGroup {
	api := parentRouterGroup.Group("api")
	api.Use(otelgin.Middleware(appName))
	api.Use(middleware.NewTraceLogMiddleware(appName))
	return api
}

func InitPublicAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, initPublicRouterFunc []InitRouterGroupFunc) {
	for _, fn := range initPublicRouterFunc {
		fn(parentRouterGroup)
	}
}

func InitPrivateAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, authMiddleware gin.HandlerFunc, initPrivateRouterFunc []InitRouterGroupFunc) {
	for _, fn := range initPrivateRouterFunc {
		fn(parentRouterGroup, authMiddleware)
	}
}
