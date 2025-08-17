package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
)

type SystemAdminInterface interface {
	AppUserID() *mbuserdomain.AppUserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type RBACUsecase interface {
	// Who can do what actions on which resources
	AddPolicyToUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, subject mbuserdomain.RBACSubject, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject, effect mbuserdomain.RBACEffect) error
}

type RBACHandler struct {
	rbacUsecase RBACUsecase
}

func NewRBACHandler(rbacUsecase RBACUsecase) *RBACHandler {
	return &RBACHandler{
		rbacUsecase: rbacUsecase,
	}
}

func (h *RBACHandler) AddPolicyToUser(c *gin.Context) {
	ctx := c.Request.Context()
	apiParam := libapi.AddPolicyToUserParameter{}
	if err := c.ShouldBindJSON(&apiParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	organizationID, err := mbuserdomain.NewOrganizationID(apiParam.OrganizationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	subject := mbuserdomain.NewRBACUser(apiParam.Subject)
	action := mbuserdomain.NewRBACAction(apiParam.Action)
	object := mbuserdomain.NewRBACObject(apiParam.Object)
	effect := mbuserdomain.NewRBACEffect(apiParam.Effect)

	if err := h.rbacUsecase.AddPolicyToUser(ctx, organizationID, subject, action, object, effect); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
}

func (h *RBACHandler) AddPolicyToGroup(c *gin.Context) {

}

func NewInitRBACRouterFunc(rbacUsecase RBACUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		rbac := parentRouterGroup.Group("rbac")
		for _, m := range middleware {
			rbac.Use(m)
		}

		rbacHandler := NewRBACHandler(rbacUsecase)
		rbac.PUT("policy/user", rbacHandler.AddPolicyToUser)
		rbac.PUT("policy/group", rbacHandler.AddPolicyToGroup)
	}
}
