package helper

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	liblibcontroller "github.com/mocoarow/cocotola-1.24/lib/controller"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type roleUser struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	role           string
}

func (o *roleUser) GetUserID() *mbuserdomain.UserID {
	return o.userID
}
func (o *roleUser) GetOrganizationID() *mbuserdomain.OrganizationID {
	return o.organizationID
}
func (o *roleUser) GetRole() string {
	return o.role
}

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

func HandleRoleUserFunction(c *gin.Context, fn func(ctx context.Context, operator service.RoleUserInterface) error, errorHandle func(ctx context.Context, c *gin.Context, err error) bool) {
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

	operator := &roleUser{
		userID:         operatorID,
		organizationID: organizationID,
		role:           c.GetString("Role"),
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

func HandleUserFunction(c *gin.Context, fn func(ctx context.Context, operator mbuserdomain.UserInterface) error, errorHandle func(ctx context.Context, c *gin.Context, err error) bool) {
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

func HandleRBACFunction[P any](c *gin.Context, getParam func(ctx context.Context, operator mbuserdomain.UserInterface) (P, bool), checkAuthorization func(ctx context.Context, operator mbuserdomain.UserInterface, param P) bool, fn func(ctx context.Context, operator mbuserdomain.UserInterface, param P) error, errorHandle func(ctx context.Context, c *gin.Context, err error) bool) {
	ctx := c.Request.Context()
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-HandleUserFunction"))

	operator, ok := getOperator(c)
	if !ok {
		return
	}

	if newCtx, err := liblibcontroller.AddBaggageMembers(ctx, map[string]string{
		"operator_id":     strconv.Itoa(operator.userID.Int()),
		"organization_id": strconv.Itoa(operator.organizationID.Int()),
	}); err == nil {
		ctx = newCtx
	}

	logger.InfoContext(ctx, "HandleRBACFunction")

	p, ok := getParam(ctx, operator)
	if !ok {
		return
	}
	if ok := checkAuthorization(ctx, operator, p); !ok {
		return
	}
	if err := fn(ctx, operator, p); err != nil {
		if handled := errorHandle(ctx, c, err); !handled {
			c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		}
	}
}

func getOperator(c *gin.Context) (*operator, bool) {
	organizationIDInt := c.GetInt("OrganizationID")
	if organizationIDInt == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return nil, false
	}

	organizationID, err := mbuserdomain.NewOrganizationID(organizationIDInt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return nil, false
	}

	userID := c.GetInt("AuthorizedUser")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return nil, false
	}

	operatorID, err := mbuserdomain.NewUserID(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return nil, false
	}

	return &operator{
		userID:         operatorID,
		organizationID: organizationID,
	}, true
}
