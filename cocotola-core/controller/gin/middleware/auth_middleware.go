package middleware

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func NewAuthMiddleware(cocotolaAuthClient service.CocotolaAuthClient) gin.HandlerFunc {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-AuthMiddleware"))

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, "AuthMiddleware")
		defer span.End()

		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			logger.InfoContext(ctx, "invalid header. Bearer not found")
			return
		}

		bearerToken := authorization[len("Bearer "):]
		appUserInfo, err := cocotolaAuthClient.RetrieveUserInfo(ctx, bearerToken)
		if err != nil {
			logger.WarnContext(ctx, fmt.Sprintf("getUserInfo: %v", err))
			return
		}

		c.Set("AuthorizedUser", appUserInfo.AppUserID)
		c.Set("OrganizationID", appUserInfo.OrganizationID)

		// logger.WarnContext(ctx, "authenticated", slog.Int("app_user_id", appUserInfo.AppUserID), slog.Int("organization_id", appUserInfo.OrganizationID))
	}
}
