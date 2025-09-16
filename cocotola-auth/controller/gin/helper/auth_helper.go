package helper

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	liblibcontroller "github.com/mocoarow/cocotola-1.24/lib/controller"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type operator struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
}

func (o *operator) GetUserID() *mbuserdomain.UserID {
	return o.userID
}
func (o *operator) GetOrganizationID() *mbuserdomain.OrganizationID {
	return o.organizationID
}

func HandleUserFunction(c *gin.Context, fn func(ctx context.Context, operator mbuserservice.OperatorInterface) error, errorHandle func(ctx context.Context, c *gin.Context, err error) bool) {
	ctx := c.Request.Context()
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-HandleUserFunction"))

	organizationIDInt := c.GetInt("OrganizationID")
	if organizationIDInt == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})

		return
	}

	organizationID, err := mbuserdomain.NewOrganizationID(organizationIDInt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})

		return
	}

	userID := c.GetInt("AuthorizedUser")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})

		return
	}

	operatorID, err := mbuserdomain.NewUserID(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})

		return
	}

	// logger.InfoContext(ctx, "", slog.Int("organization_id", organizationID.Int()), slog.Int("operator_id", operatorID.Int()))

	operator := &operator{
		userID:         operatorID,
		organizationID: organizationID,
	}

	if newCtx, err := liblibcontroller.AddBaggageMembers(ctx, map[string]string{
		"operator_id":     strconv.Itoa(operatorID.Int()),
		"organization_id": strconv.Itoa(organizationID.Int()),
	}); err == nil {
		ctx = newCtx
	}

	logger.InfoContext(ctx, "xxxxxxxx")

	if err := fn(ctx, operator); err != nil {
		if handled := errorHandle(ctx, c, err); !handled {
			c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		}
	}
}
