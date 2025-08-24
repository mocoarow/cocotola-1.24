package gateway

import (
	"log/slog"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
)

func (c *GoogleAuthClient) SetLogger() {
	c.logger = slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-GoogleAuthClient"))
}
