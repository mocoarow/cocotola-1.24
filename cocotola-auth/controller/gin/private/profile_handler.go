package private

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type ProfileQueryUsecase interface {
	GetMyProfile(ctx context.Context, operator mbuserdomain.UserInterface) (*domain.ProfileModel, error)
}

type ProfileHandler struct {
	profileQueryUsecase ProfileQueryUsecase
	logger              *slog.Logger
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	helper.HandleUserFunction(c, func(ctx context.Context, operator mbuserdomain.UserInterface) error {
		result, err := h.profileQueryUsecase.GetMyProfile(ctx, operator)
		if err != nil {
			return mbliberrors.Errorf("GetMyProfile: %w", err)
		}

		apiResp := libapiauth.ProfileResponse{
			PrivateSpaceID: result.PrivateSpaceID.Int(),
		}
		c.JSON(http.StatusOK, apiResp)

		return nil
	}, h.errorHandle)
}

func NewProfileHandler(profileQueryUsecase ProfileQueryUsecase) *ProfileHandler {
	return &ProfileHandler{
		profileQueryUsecase: profileQueryUsecase,
		logger:              slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileHandler")),
	}
}

func (h *ProfileHandler) errorHandle(ctx context.Context, _ *gin.Context, err error) bool {
	h.logger.ErrorContext(ctx, fmt.Sprintf("ProfileHandler. error: %+v", err))

	return false
}

func NewInitProfileRouterFunc(profileQueryUsecase ProfileQueryUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		profile := parentRouterGroup.Group("profile")
		for _, m := range middleware {
			profile.Use(m)
		}
		profileHandler := NewProfileHandler(profileQueryUsecase)
		profile.GET("me", profileHandler.GetMyProfile)
	}
}
