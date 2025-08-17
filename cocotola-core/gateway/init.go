package gateway

import (
	"go.opentelemetry.io/otel"
	// liblog "github.com/mocoarow/cocotola-1.24/lib/log"
)

const (
// loggerKey = liblog.CoreGatewayLoggerContextKey
)

var (
	tracer = otel.Tracer("github.com/mocoarow/cocotola-1.24/cocotola-core/gateway")
)
