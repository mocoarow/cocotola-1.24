package gateway

import (
	"net/http"

	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway")
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
