package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type UserUsecase interface {
	RegisterAppUser(ctx context.Context, operator service.OperatorInterface, param *mbuserservice.AppUserAddParameter) (*domain.AuthTokenSet, error)
}

type UserHandler struct {
	userUsecase UserUsecase
	logger      *slog.Logger
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		logger:      slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-UserHandler")),
	}
}

func (h *UserHandler) RegisterAppUser(c *gin.Context) {
	helper.HandleAppUserFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		apiParam := libapi.AppUserAddRequest{}
		if err := c.ShouldBindJSON(&apiParam); err != nil {
			h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter: %v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return nil
		}

		param, err := mbuserservice.NewAppUserAddParameter(apiParam.LoginID, apiParam.Username, apiParam.Password, "", "", "", "")
		if err != nil {
			h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter: %v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return nil
		}

		authResult, err := h.userUsecase.RegisterAppUser(c.Request.Context(), operator, param)
		if err != nil {
			if errors.Is(err, domain.ErrUnauthenticated) {
				h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter. err: %v", err))
				c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
				return nil
			}

			h.logger.ErrorContext(ctx, fmt.Sprintf("failed to RegisterStudent. err: %+v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
			return nil
		}

		c.JSON(http.StatusOK, libapi.AuthResponse{
			AccessToken:  &authResult.AccessToken,
			RefreshToken: &authResult.RefreshToken,
		})
		return nil
	}, h.errorHandle)
}

func (h *UserHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	if errors.Is(err, mblibdomain.ErrInvalidArgument) {
		h.logger.WarnContext(ctx, fmt.Sprintf("PrivateDeckHandler err: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return true
	}
	if errors.Is(err, service.ErrAppUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		return true
	}
	h.logger.ErrorContext(ctx, fmt.Sprintf("DeckHandler. error: %+v", err))
	return false
}

func NewInitUserRouterFunc(userUsecase UserUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		user := parentRouterGroup.Group("user")
		userHandler := NewUserHandler(userUsecase)
		for _, m := range middleware {
			user.Use(m)
		}
		user.POST("", userHandler.RegisterAppUser)
	}
}
