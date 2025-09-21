package controller

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapicore "github.com/mocoarow/cocotola-1.24/lib/api/core"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type ProfileUsecase interface {
	GetMyProfile(ctx context.Context, operator mbuserdomain.UserInterface, bearerToken string) (*domain.ProfileModel, error)
}

type ProfileHandler struct {
	profileUsecase ProfileUsecase
	logger         *slog.Logger
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	helper.HandleRoleUserFunction(c, func(ctx context.Context, operator service.RoleUserInterface) error {
		result, err := h.profileUsecase.GetMyProfile(ctx, operator, operator.GetBearerToken())
		if err != nil {
			return mbliberrors.Errorf("GetMyProfile: %w", err)
		}

		apiResp := libapicore.GetMyProfileResponse{
			PrivateSpace: libapicore.GetMyProfileResponseSpace{
				SpaceID:      result.PrivateSpaceID.Int(),
				RootFolderID: result.RootFolderID.Int(),
			},
		}
		c.JSON(http.StatusOK, apiResp)

		return nil
	}, h.errorHandle)
}

func NewProfileHandler(profileUsecase ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase: profileUsecase,
		logger:         slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileHandler")),
	}
}

func (h *ProfileHandler) errorHandle(ctx context.Context, _ *gin.Context, err error) bool {
	h.logger.ErrorContext(ctx, fmt.Sprintf("ProfileHandler. error: %+v", err))

	return false
}

func NewInitProfileRouterFunc(profileUsecase ProfileUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		profile := parentRouterGroup.Group("profile")
		for _, m := range middleware {
			profile.Use(m)
		}
		profileHandler := NewProfileHandler(profileUsecase)
		profile.GET("me", profileHandler.GetMyProfile)
	}
}
