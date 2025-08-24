package gateway

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller"
)

type Process func() error

type ProcessFunc func(ctx context.Context) Process

func WithAppServerProcess(router http.Handler, port int, readHeaderTimeout, shutdownTime time.Duration) ProcessFunc {
	return func(ctx context.Context) Process {
		return func() error {
			return libcontroller.AppServerProcess(ctx, router, port, readHeaderTimeout, shutdownTime)
		}
	}
}
func WithMetricsServerProcess(port int, shutdownTime int) ProcessFunc {
	return func(ctx context.Context) Process {
		return func() error {
			return mblibgateway.MetricsServerProcess(ctx, port, shutdownTime)
		}
	}
}

func WithSignalWatchProcess() ProcessFunc {
	return func(ctx context.Context) Process {
		return func() error {
			return mblibgateway.SignalWatchProcess(ctx)
		}
	}
}

func Run(ctx context.Context, processFuncs ...ProcessFunc) int {
	var eg *errgroup.Group
	eg, ctx = errgroup.WithContext(ctx)

	for _, pf := range processFuncs {
		eg.Go(pf(ctx))
	}
	eg.Go(func() error {
		<-ctx.Done()

		return ctx.Err()
	})

	if err := eg.Wait(); err != nil {
		return 1
	}

	return 0
}
