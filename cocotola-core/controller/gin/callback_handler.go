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
)

type CallbackUsecase interface {
	OnAddAppUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.AppUserID) error
}

type CallbackHandler struct {
	callbackUsecase CallbackUsecase
	logger          *slog.Logger
}

func NewCallbackHandler(callbackUsecase CallbackUsecase) *CallbackHandler {
	return &CallbackHandler{
		callbackUsecase: callbackUsecase,
		logger:          slog.Default().With(slog.String(mbliblog.LoggerNameKey, "DeckHandler")),
	}
}

func (h *CallbackHandler) OnAddAppUser(c *gin.Context) {
	ctx := c.Request.Context()
	apiReq := libapi.CallbackOnAddAppUserRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	organizationID, err := mbuserdomain.NewOrganizationID(apiReq.OrganizationID)
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	appUserID, err := mbuserdomain.NewAppUserID(apiReq.AppUserID)
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	h.logger.Info("OnAddAppUser", slog.Int("appUserID", appUserID.Int()))
	if err := h.callbackUsecase.OnAddAppUser(ctx, organizationID, appUserID); err != nil {
		h.logger.ErrorContext(ctx, fmt.Sprintf("on add app user: %+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}
}

func NewInitCallbackRouterFunc(callbackUsecase CallbackUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		callback := parentRouterGroup.Group("callback")
		callbackHandler := NewCallbackHandler(callbackUsecase)
		for _, m := range middleware {
			callback.Use(m)
		}
		callback.POST("on-add-user", callbackHandler.OnAddAppUser)
	}
}
