package controller

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapicore "github.com/mocoarow/cocotola-1.24/lib/api/core"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
)

type CallbackUsecase interface {
	OnAddUserSpace(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.UserID, spaceID *mbuserdomain.SpaceID) error
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

// func (h *CallbackHandler) OnAddUser(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	var apiReq libapicore.CallbackOnAddUserRequest
// 	if err := c.ShouldBindJSON(&apiReq); err != nil {
// 		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
// 		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

// 		return
// 	}

// 	organizationID, err := mbuserdomain.NewOrganizationID(apiReq.OrganizationID)
// 	if err != nil {
// 		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
// 		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

// 		return
// 	}

// 	appUserID, err := mbuserdomain.NewUserID(apiReq.UserID)
// 	if err != nil {
// 		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
// 		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

// 		return
// 	}

// 	h.logger.Info("OnAddUser", slog.Int("appUserID", appUserID.Int()))
// 	if err := h.callbackUsecase.OnAddUser(ctx, organizationID, appUserID); err != nil {
// 		h.logger.ErrorContext(ctx, fmt.Sprintf("on add app user: %+v", err))
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusBadRequest)})

//			return
//		}
//	}
func (h *CallbackHandler) OnAddUserSpace(c *gin.Context) {
	ctx := c.Request.Context()
	var apiReq libapicore.CallbackOnAddUserSpaceRequest
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

	appUserID, err := mbuserdomain.NewUserID(apiReq.UserID)
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	spaceID, err := mbuserdomain.NewSpaceID(apiReq.SpaceID)
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	h.logger.Info("OnAddUserSpace", slog.Int("appUserID", appUserID.Int()))
	if err := h.callbackUsecase.OnAddUserSpace(ctx, organizationID, appUserID, spaceID); err != nil {
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
		// callback.POST("on-add-user", callbackHandler.OnAddUser)
		callback.POST("on-add-user-space", callbackHandler.OnAddUserSpace)
	}
}
