package middleware

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func NewAuthMiddleware(systemToken libdomain.SystemToken, authTokenManager service.AuthTokenManager, transactionManager service.TransactionManager) gin.HandlerFunc {
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
		appUserModel, err := service.GetUserInfo(ctx, systemToken, authTokenManager, transactionManager, bearerToken)
		if err != nil {
			logger.WarnContext(ctx, fmt.Sprintf("getUserInfo: %v", err))

			return
		}

		c.Set("AuthorizedUser", appUserModel.AppUserID.Int())
		c.Set("OrganizationID", appUserModel.OrganizationID.Int())

		// logger.WarnContext(ctx, "authenticated", slog.Int("app_user_id", appUserInfo.AppUserID), slog.Int("organization_id", appUserInfo.OrganizationID))
	}
}
