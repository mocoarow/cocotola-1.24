package public

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type GuestUsecase interface {
	Authenticate(ctx context.Context, organizationName string) (*domain.AuthTokenSet, error)
}

type GuestAuthHandler struct {
	guestUsecase GuestUsecase
	logger       *slog.Logger
}

func NewGuestAuthHandler(guestUsecase GuestUsecase) *GuestAuthHandler {
	return &GuestAuthHandler{
		guestUsecase: guestUsecase,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-GuestAuthHandler")),
	}
}

func (h *GuestAuthHandler) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()

	var apiReq libapiauth.GuestAuthRequest
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

		return
	}

	authResult, err := h.guestUsecase.Authenticate(ctx, apiReq.OrganizationName)
	if err != nil {
		if errors.Is(err, mbuserservice.ErrSystemOwnerNotFound) {
			h.logger.InfoContext(ctx, fmt.Sprintf("system owner not found: %+v", err))
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})

			return
		}
		if errors.Is(err, domain.ErrUnauthenticated) {
			h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})

			return
		}

		h.logger.ErrorContext(ctx, fmt.Sprintf("guestUsecase.Authenticate: %+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})

		return
	}

	c.JSON(http.StatusOK, libapiauth.AuthResponse{
		AccessToken:  &authResult.AccessToken,
		RefreshToken: &authResult.RefreshToken,
	})
}

func NewInitGuestRouterFunc(guest GuestUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		auth := parentRouterGroup.Group("guest")
		for _, m := range middleware {
			auth.Use(m)
		}

		guestAuthHandler := NewGuestAuthHandler(guest)
		auth.POST("authenticate", guestAuthHandler.Authenticate)
	}
}
