package controller

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type operator struct {
	appUserID      *mbuserdomain.AppUserID
	organizationID *mbuserdomain.OrganizationID
}

func (o *operator) AppUserID() *mbuserdomain.AppUserID {
	return o.appUserID
}
func (o *operator) OrganizationID() *mbuserdomain.OrganizationID {
	return o.organizationID
}

type SystemAdminInterface interface {
	AppUserID() *mbuserdomain.AppUserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type RBACUsecase interface {
	// Who can do what actions on which resources
	AddPolicyToUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, subject mbuserdomain.RBACSubject, listOfActionObjectEffect []mbuserdomain.RBACActionObjectEffect) error
	// Authorize(ctx context.Context, operator service.OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject) (bool, error)
	// Check whether the operator can do the action on the object
	CheckAuthorization(ctx context.Context, operator service.OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject) (bool, error)
}

type RBACHandler struct {
	rbacUsecase RBACUsecase
	logger      *slog.Logger
}

func NewRBACHandler(rbacUsecase RBACUsecase) *RBACHandler {
	return &RBACHandler{
		rbacUsecase: rbacUsecase,
		logger:      slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-RBACHandler")),
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
	appUserID, err := mbuserdomain.NewAppUserID(apiParam.AppUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	subject := mbuserdomain.NewRBACAppUser(appUserID)

	listofActionObjectEffect := make([]mbuserdomain.RBACActionObjectEffect, 0)
	for _, aoe := range apiParam.ListOfActionObjectEffect {
		actionObj := mbuserdomain.RBACActionObjectEffect{
			Action: mbuserdomain.NewRBACAction(aoe.Action),
			Object: mbuserdomain.NewRBACObject(aoe.Object),
			Effect: mbuserdomain.NewRBACEffect(aoe.Effect),
		}
		listofActionObjectEffect = append(listofActionObjectEffect, actionObj)
	}

	if err := h.rbacUsecase.AddPolicyToUser(ctx, organizationID, subject, listofActionObjectEffect); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
}

func (h *RBACHandler) AddPolicyToGroup(_ *gin.Context) {

}

func (h *RBACHandler) CheckAuthorization(c *gin.Context) {
	ctx := c.Request.Context()
	apiParam := libapi.AuthorizeRequest{}
	if err := c.ShouldBindJSON(&apiParam); err != nil {
		h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	organizationID, err := mbuserdomain.NewOrganizationID(apiParam.OrganizationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	operatorID, err := mbuserdomain.NewAppUserID(apiParam.AppUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	operator := &operator{
		appUserID:      operatorID,
		organizationID: organizationID,
	}

	ok, err := h.rbacUsecase.CheckAuthorization(ctx, operator, mbuserdomain.NewRBACAction(apiParam.Action), mbuserdomain.NewRBACObject(apiParam.Object))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, libapi.AuthorizeResponse{
		Authorized: ok,
	})
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
		// rbac.POST("authorize", rbacHandler.Authorize)
		rbac.POST("check-authorization", rbacHandler.CheckAuthorization) // TODO: implement
	}
}
