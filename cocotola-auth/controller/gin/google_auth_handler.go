package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type googleAuthParameter struct {
	OrganizationName string `json:"organizationName"`
	SessionState     string `json:"sessionState"`
	ParamState       string `json:"paramState"`
	Code             string `json:"code"`
}

type GoogleUserUsecase interface {
	GenerateState(context.Context) (string, error)
	Authorize(ctx context.Context, state, code, organizationName string) (*domain.AuthTokenSet, error)
}

type GoogleUserHandler struct {
	googleUserUsecase GoogleUserUsecase
	logger            *slog.Logger
}

func NewGoogleAuthHandler(googleUserUsecase GoogleUserUsecase) *GoogleUserHandler {
	return &GoogleUserHandler{
		googleUserUsecase: googleUserUsecase,
		logger:            slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-GoogleUserHandler")),
	}
}

func (h *GoogleUserHandler) GenerateState(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Info("GenerateState")

	state, err := h.googleUserUsecase.GenerateState(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": state})
}

func (h *GoogleUserHandler) Authorize(c *gin.Context) {
	ctx := c.Request.Context()

	h.logger.Info("Authorize")

	googleAuthParameter := googleAuthParameter{}
	if err := c.ShouldBindJSON(&googleAuthParameter); err != nil {
		h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter. err: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return
	}

	// // logger.Infof("RetrieveAccessToken. code: %s", googleAuthParameter)
	// googleAuthResponse, err := h.googleUserUsecase.RetrieveAccessToken(ctx, googleAuthParameter.Code)
	// if err != nil {
	// 	// logger.Warnf("failed to RetrieveAccessToken. err: %v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
	// 	return
	// }

	// // logger.Infof("RetrieveUserInfo. googleResponse: %+v", googleAuthResponse)
	// userInfo, err := h.googleUserUsecase.RetrieveUserInfo(ctx, googleAuthResponse)
	// if err != nil {
	// 	// logger.Warnf("failed to RetrieveUserInfo. error: %v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
	// 	return
	// }

	// logger.Info("RegisterAppUser")
	// authResult, err := h.googleUserUsecase.RegisterAppUser(ctx, userInfo, googleAuthResponse, googleAuthParameter.OrganizationName)
	// if err != nil {
	// 	// logger.Warnf("failed to RegisterStudent. err: %+v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusBadRequest)})
	// 	return
	// }

	// user, err := c.Cookie("auth_user")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "auth_user is empty"})
	// 	return
	// }
	if googleAuthParameter.SessionState == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sessionState is empty"})
		return
	}
	if googleAuthParameter.ParamState == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "paramState is empty"})
		return
	}
	if googleAuthParameter.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "code is empty"})
		return
	}
	if googleAuthParameter.OrganizationName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "organizationName is empty"})
		return
	}
	if googleAuthParameter.SessionState != googleAuthParameter.ParamState {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sessionState and paramState are not equal"})
		return
	}

	authResult, err := h.googleUserUsecase.Authorize(ctx, googleAuthParameter.ParamState, googleAuthParameter.Code, googleAuthParameter.OrganizationName)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthenticated) {
			h.logger.InfoContext(ctx, fmt.Sprintf("invalid parameter. err: %v", err))
			c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
			return
		}

		h.logger.ErrorContext(ctx, fmt.Sprintf("failed to RegisterStudent. err: %+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, libapi.AuthResponse{
		AccessToken:  &authResult.AccessToken,
		RefreshToken: &authResult.RefreshToken,
	})
}

func NewInitGoogleRouterFunc(googleUserUsecase GoogleUserUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		auth := parentRouterGroup.Group("google")
		for _, m := range middleware {
			auth.Use(m)
		}

		googleAuthHandler := NewGoogleAuthHandler(googleUserUsecase)
		auth.GET("state", googleAuthHandler.GenerateState)
		auth.POST("authorize", googleAuthHandler.Authorize)
	}
}
