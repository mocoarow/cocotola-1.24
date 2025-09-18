package private

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
)

type CallbackUsecase interface {
	OnAddUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error
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

func (h *CallbackHandler) OnAddUser(c *gin.Context) {
	ctx := c.Request.Context()
	var apiReq libapiauth.CallbackOnAddUserRequest
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

	userID, err := mbuserdomain.NewUserID(apiReq.UserID)
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

		return
	}

	h.logger.Info("OnAddUser", slog.Int("userID", userID.Int()))
	if err := h.callbackUsecase.OnAddUser(ctx, organizationID, userID); err != nil {
		h.logger.ErrorContext(ctx, fmt.Sprintf("on add user: %+v", err))
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
		callback.POST("on-add-user", callbackHandler.OnAddUser)
	}
}
