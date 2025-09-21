package middleware

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/usecase"
)

func NewAuthMiddleware(systemToken libdomain.SystemToken, authTokenManager service.AuthTokenManager, mbNonTxManager mbuserservice.TransactionManager, _ mbuserservice.RepositoryFactory) gin.HandlerFunc {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-AuthMiddleware"))

	sysAdmin := service.NewSystemAdmin(systemToken)
	getUserInfoQuery := usecase.NewGetUserInfoQuery(mbNonTxManager, authTokenManager)

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
		userModel, err := getUserInfoQuery.Execute(ctx, sysAdmin, bearerToken)
		if err != nil {
			logger.WarnContext(ctx, fmt.Sprintf("getUserInfo: %v", err))
			return
		}

		c.Set("AuthorizedUser", userModel.UserID.Int())
		c.Set("OrganizationID", userModel.OrganizationID.Int())
	}
}
