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
)

type SpaceQueryUsecase interface {
	FindPublicSpaces(ctx context.Context, operator mbuserdomain.UserInterface) ([]*mbuserdomain.Space, error)
}

type SpaceHandler struct {
	spaceQueryUsecase SpaceQueryUsecase
	logger            *slog.Logger
}

func (h *SpaceHandler) FindSpaces(c *gin.Context) {
	helper.HandleUserFunction(c, func(ctx context.Context, operator mbuserdomain.UserInterface) error {
		result, err := h.spaceQueryUsecase.FindPublicSpaces(ctx, operator)
		if err != nil {
			return mbliberrors.Errorf("FindPublicSpaces: %w", err)
		}

		spaces := make([]libapiauth.FindSpacesResponseSpace, 0, len(result))
		for _, s := range result {
			spaces = append(spaces, libapiauth.FindSpacesResponseSpace{
				ID:   s.SpaceID.Int(),
				Key:  s.KeyName,
				Name: s.Name,
			})
		}
		apiResp := libapiauth.FindSpacesResponse{
			Results: spaces,
		}
		c.JSON(http.StatusOK, apiResp)
		return nil
	}, h.errorHandle)
}

func NewSpaceHandler(spaceQueryUsecase SpaceQueryUsecase) *SpaceHandler {
	return &SpaceHandler{
		spaceQueryUsecase: spaceQueryUsecase,
		logger:            slog.Default().With(slog.String(mbliblog.LoggerNameKey, "SpaceHandler")),
	}
}

func (h *SpaceHandler) errorHandle(ctx context.Context, _ *gin.Context, err error) bool {
	h.logger.ErrorContext(ctx, fmt.Sprintf("SpaceHandler. error: %+v", err))

	return false
}

func NewInitSpaceRouterFunc(spaceQueryUsecase SpaceQueryUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		space := parentRouterGroup.Group("space")
		for _, m := range middleware {
			space.Use(m)
		}
		spaceHandler := NewSpaceHandler(spaceQueryUsecase)
		space.GET("", spaceHandler.FindSpaces)
	}
}
