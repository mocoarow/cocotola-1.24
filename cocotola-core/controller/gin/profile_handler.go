package controller

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/usecase"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
)

type ProfileHandler struct {
	profileUsecase usecase.ProfileUsecase
	logger         *slog.Logger
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		result, err := h.profileUsecase.GetMyProfile(ctx, operator)
		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, result)
		return nil
	}, h.errorHandle)
}

func NewProfileHandler(profileUsecase usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase: profileUsecase,
		logger:         slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileHandler")),
	}
}

func (h *ProfileHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	h.logger.ErrorContext(ctx, fmt.Sprintf("ProfileHandler. error: %+v", err))
	return false
}

func NewInitProfileRouterFunc(profileUsecase usecase.ProfileUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		profile := parentRouterGroup.Group("profile")
		for _, m := range middleware {
			profile.Use(m)
		}
		profileHandler := NewProfileHandler(profileUsecase)
		parentRouterGroup.GET("me", profileHandler.GetMyProfile)
	}
}
