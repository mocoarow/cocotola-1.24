package middleware

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

func NewAuthMiddleware(cocotolaAuthClient libapi.CocotolaAuthClient) gin.HandlerFunc {
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

		logger.InfoContext(ctx, fmt.Sprintf("length of groups: %d", len(appUserInfo.UserGroups)))

		for _, g := range appUserInfo.UserGroups {
			logger.InfoContext(ctx, fmt.Sprintf("group: %s", g))
		}

		logger.InfoContext(ctx, fmt.Sprintf("UserID: %d", appUserInfo.UserID))
		c.Set("AuthorizedUser", appUserInfo.UserID)
		c.Set("OrganizationID", appUserInfo.OrganizationID)
		if libdomain.IsGuestLoginID(appUserInfo.LoginID) {
			c.Set("Role", "guest")
		} else {
			c.Set("Role", "student")
		}

		// logger.WarnContext(ctx, "authenticated", slog.Int("app_user_id", appUserInfo.UserID), slog.Int("organization_id", appUserInfo.OrganizationID))
	}
}
