package initialize

import (
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/mocoarow/cocotola-auth/initialize")
)
