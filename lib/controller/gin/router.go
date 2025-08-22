package controller

import (
	"context"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"

	"github.com/mocoarow/cocotola-1.24/lib/config"
	"github.com/mocoarow/cocotola-1.24/lib/controller/gin/middleware"
)

type InitRouterGroupFunc func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc)

func InitRootRouterGroup(ctx context.Context, corsConfig *mblibconfig.CORSConfig, logConfig *mblibconfig.LogConfig, debugConfig *config.DebugConfig) *gin.Engine {
	if !debugConfig.Gin {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// cors
	ginCorsConfig := mblibconfig.InitCORS(corsConfig)

	router.Use(gin.Recovery())
	router.Use(cors.New(ginCorsConfig))
	if value, ok := logConfig.Enabled["accessLog"]; ok && value {
		router.Use(sloggin.New(slog.Default()))
	}

	if debugConfig.Wait {
		router.Use(middleware.NewWaitMiddleware())
	}

	return router
}

func InitAPIRouterGroup(ctx context.Context, parentRouterGroup gin.IRouter, appName string, logConfig *mblibconfig.LogConfig) *gin.RouterGroup {
	api := parentRouterGroup.Group("api")
	api.Use(otelgin.Middleware(appName))
	if value, ok := logConfig.Enabled["traceLog"]; ok && value {
		api.Use(middleware.NewTraceLogMiddleware(appName, true))
	} else {
		api.Use(middleware.NewTraceLogMiddleware(appName, false))
	}
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
