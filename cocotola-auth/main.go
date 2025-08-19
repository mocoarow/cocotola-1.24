package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbsql "github.com/mocoarow/cocotola-1.24/moonbeam/sqls"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	libgateway "github.com/mocoarow/cocotola-1.24/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/initialize"
)

func main() {
	ctx := context.Background()

	mbliberrors.UseXerrorsErrorf()

	// load config
	cfg, err := config.LoadConfig()
	libdomain.CheckError(err)

	// init log
	mblibconfig.InitLog(cfg.Log)
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, "main"))

	// init tracer
	tp, err := mblibconfig.InitTracerProvider(ctx, initialize.AppName, cfg.Trace)
	libdomain.CheckError(err)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// init db
	dialect, db, sqlDB, err := mblibconfig.InitDB(ctx, cfg.DB, mbsql.SQL)
	libdomain.CheckError(err)
	defer sqlDB.Close()
	defer tp.ForceFlush(ctx) // flushes any pending spans

	// init gin
	router := libcontroller.InitRootRouterGroup(ctx, cfg.CORS, cfg.Debug)

	if err := initialize.Initialize(ctx, router, dialect, cfg.DB.DriverName, db, cfg.App); err != nil {
		libdomain.CheckError(err)
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

// func initGin(ctx context.Context, cfg *config.Config, txManager, nonTxManager service.TransactionManager) *gin.Engine {
// 	if !cfg.Debug.Gin {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs(cfg.Auth, txManager, nonTxManager)
// 	// privateRouterGroupFuncs := controller.InitPublicRouterGroupFuncs(cfg.Auth, txManager, nonTxManager)
// 	router := gin.New()
// 	initialize.InitAppServer(ctx, router, cfg.CORS, cfg.Debug, cfg.App.Name, publicRouterGroupFuncs)
// 	return router
// }
