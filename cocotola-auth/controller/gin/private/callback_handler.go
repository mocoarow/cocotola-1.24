package private

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	liblibcontroller "github.com/mocoarow/cocotola-1.24/lib/controller"
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
	if newCtx, err := liblibcontroller.AddBaggageMembers(ctx, map[string]string{
		"operator_id":     strconv.Itoa(apiReq.UserID.Value.Int()),
		"organization_id": strconv.Itoa(apiReq.OrganizationID.Value.Int()),
	}); err == nil {
		ctx = newCtx
		// Add baggage members as span attributes
		liblibcontroller.AddBaggageToCurrentSpan(ctx)
	}

	if err := h.callbackUsecase.OnAddUser(ctx, apiReq.OrganizationID.Value, apiReq.UserID.Value); err != nil {
		h.logger.ErrorContext(ctx, fmt.Sprintf("on add user: %+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	c.Status(http.StatusOK)
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
