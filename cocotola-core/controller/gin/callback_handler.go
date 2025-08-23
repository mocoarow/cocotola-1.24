package controller

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CallbackOnAddAppUserParameter struct {
	AppUserID int `json:"appUserId" binding:"required,gte=1"`
}
type CallbackUsecase interface {
	OnAddAppUser(ctx context.Context, appUserID *mbuserdomain.AppUserID) error
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
	apiParam := CallbackOnAddAppUserParameter{}
	if err := c.ShouldBindJSON(&apiParam); err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	appUserID, err := mbuserdomain.NewAppUserID(apiParam.AppUserID)
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	h.logger.Info("OnAddAppUser", slog.Int("appUserID", appUserID.Int()))
	if err := h.callbackUsecase.OnAddAppUser(ctx, appUserID); err != nil {
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
