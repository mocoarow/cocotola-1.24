package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	libconfig "github.com/mocoarow/cocotola-1.24/lib/config"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	libgateway "github.com/mocoarow/cocotola-1.24/lib/gateway"
)

type ServerConfig struct {
	HTTPPort             int `yaml:"httpPort" validate:"required"`
	MetricsPort          int `yaml:"metricsPort" validate:"required"`
	ReadHeaderTimeoutSec int `yaml:"readHeaderTimeoutSec" validate:"gte=1"`
}

type Config struct {
	Server   *ServerConfig             `yaml:"server" validate:"required"`
	Trace    *mblibconfig.TraceConfig  `yaml:"trace" validate:"required"`
	CORS     *mblibconfig.CORSConfig   `yaml:"cors" validate:"required"`
	Shutdown *libconfig.ShutdownConfig `yaml:"shutdown" validate:"required"`
	Log      *mblibconfig.LogConfig    `yaml:"log" validate:"required"`
	Debug    *libconfig.DebugConfig    `yaml:"debug"`
}

const AppName = "cocotola-empty"

func main() {
	ctx := context.Background()
	env := flag.String("env", "", "environment")
	flag.Parse()
	appEnv := libdomain.GetNonEmptyValue(*env, os.Getenv("APP_ENV"), "local")
	slog.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	cfg := &Config{
		Server: &ServerConfig{
			HTTPPort:             8080,
			MetricsPort:          8081,
			ReadHeaderTimeoutSec: 10,
		},
		Trace: &mblibconfig.TraceConfig{
			Exporter: "gcp",
			Google: &mblibconfig.GoogleTraceConfig{
				ProjectID: "mocoarow-25-08",
			},
		},
		CORS: &mblibconfig.CORSConfig{
			AllowOrigins: []string{"*"},
		},
		Shutdown: &libconfig.ShutdownConfig{
			TimeSec1: 10,
			TimeSec2: 10,
		},
		Log: &mblibconfig.LogConfig{
			Level:    "info",
			Platform: "gcp",
		},
		Debug: &libconfig.DebugConfig{
			Gin:  false,
			Wait: false,
		},
	}

	// init log
	mblibconfig.InitLog(cfg.Log)
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, "main"))
	logger.InfoContext(ctx, fmt.Sprintf("env: %s", appEnv))

	// init tracer
	tp, err := mblibconfig.InitTracerProvider(ctx, AppName, cfg.Trace)
	libdomain.CheckError(err)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	router := libcontroller.InitRootRouterGroup(ctx, cfg.CORS, cfg.Debug)

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, router, AppName)
	// v1
	v1 := api.Group("v1")
	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, []libcontroller.InitRouterGroupFunc{
		func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
			test := parentRouterGroup.Group("test")
			for _, m := range middleware {
				test.Use(m)
			}
			test.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "pong"})
			})
			test.POST("/200", func(c *gin.Context) {
				logger.InfoContext(ctx, "POST /200")
				params := gin.H{}
				if err := c.BindJSON(&params); err != nil {
					logger.InfoContext(ctx, fmt.Sprintf("err: %+v", err))
					c.Status(http.StatusBadRequest)
					return
				}

				logger.InfoContext(ctx, fmt.Sprintf("params: %+v", params))
				c.Status(http.StatusOK)
			})
		},
	})

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
