package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	"github.com/mocoarow/cocotola-1.24/moonbeam/sqls"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	libgateway "github.com/mocoarow/cocotola-1.24/lib/gateway"

	authinit "github.com/mocoarow/cocotola-1.24/cocotola-auth/initialize"
	coreinit "github.com/mocoarow/cocotola-1.24/cocotola-core/initialize"

	"github.com/mocoarow/cocotola-1.24/cocotola-app/config"
	web "github.com/mocoarow/cocotola-1.24/cocotola-app/web_dist"
)

const AppName = "cocotola-app"

func main() {
	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := libdomain.GetNonEmptyValue(*env, os.Getenv("APP_ENV"), "local")
	slog.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	mbliberrors.UseXerrorsErrorf()

	cfg, err := config.LoadConfig(appEnv)
	libdomain.CheckError(err)

	// init log
	mblibconfig.InitLog(cfg.Log)
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, "main"))
	logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// init tracer
	tp, err := mblibconfig.InitTracerProvider(ctx, AppName, cfg.Trace)
	libdomain.CheckError(err)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// init db
	dialect, db, sqlDB, err := mblibconfig.InitDB(ctx, cfg.DB, sqls.SQL)
	libdomain.CheckError(err)
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	router := libcontroller.InitRootRouterGroup(ctx, cfg.CORS, cfg.Debug)

	// web
	{
		viteStaticFS, err := fs.Sub(web.Web, "flutter")
		libdomain.CheckError(err)
		initGinWeb(ctx, router, viteStaticFS, "flutter")
	}
	// auth
	{
		auth := router.Group("auth")
		if err := authinit.Initialize(ctx, auth, dialect, cfg.DB.DriverName, db, cfg.Auth); err != nil {
			libdomain.CheckError(err)
		}
	}
	// core
	{
		core := router.Group("core")
		if err := coreinit.Initialize(ctx, core, dialect, cfg.DB.DriverName, db, cfg.Core); err != nil {
			libdomain.CheckError(err)
		}
	}

	// run
	readHeaderTimeout := time.Duration(cfg.Server.ReadHeaderTimeoutSec) * time.Second
	shutdownTime := time.Duration(cfg.Shutdown.TimeSec1) * time.Second
	result := libgateway.Run(ctx,
		libgateway.WithAppServerProcess(router, cfg.Server.HTTPPort, readHeaderTimeout, shutdownTime),
		libgateway.WithSignalWatchProcess(),
		libgateway.WithMetricsServerProcess(cfg.Server.MetricsPort, cfg.Shutdown.TimeSec1),
	)

	gracefulShutdownTime2 := time.Duration(cfg.Shutdown.TimeSec2) * time.Second
	time.Sleep(gracefulShutdownTime2)
	logger.InfoContext(ctx, "exited")
	os.Exit(result)
}

func initGinWeb(ctx context.Context, router *gin.Engine, viteStaticFS fs.FS, webType string) {
	router.NoRoute(func(c *gin.Context) {
		logger := slog.Default()
		logger.InfoContext(c.Request.Context(), c.Request.URL.Path)
		if webType == "flutter" {
			for _, prefix := range web.GetFlutterResources() {
				if strings.HasPrefix(c.Request.RequestURI, prefix) {
					c.FileFromFS(c.Request.URL.Path, http.FS(viteStaticFS))
					return
				}
			}
		} else if webType == "react" {
			for _, prefix := range web.GetReactResources() {
				if strings.HasPrefix(c.Request.RequestURI, prefix) {
					c.FileFromFS(c.Request.URL.Path, http.FS(viteStaticFS))
					return
				}
			}
		}

		if !strings.HasPrefix(c.Request.URL.Path, "/auth") &&
			!strings.HasPrefix(c.Request.URL.Path, "/core") &&
			!strings.HasPrefix(c.Request.URL.Path, "/synthesizer") &&
			!strings.HasPrefix(c.Request.URL.Path, "/tatoeba") {
			c.FileFromFS("", http.FS(viteStaticFS))
			return
		}
	})
}
