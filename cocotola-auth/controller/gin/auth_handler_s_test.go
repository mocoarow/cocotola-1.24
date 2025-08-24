//go:build small

package controller_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libconfig "github.com/mocoarow/cocotola-1.24/lib/config"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin"
	controllermock "github.com/mocoarow/cocotola-1.24/cocotola-auth/controller/gin/mocks"
)

var (
	anyOfCtx = mock.MatchedBy(func(_ context.Context) bool { return true })
	// corsConfig   cors.Config
	corsConfig   *mblibconfig.CORSConfig
	logConfig    *mblibconfig.LogConfig
	serverConfig *config.ServerConfig
	authConfig   *config.AuthConfig
	debugConfig  *libconfig.DebugConfig
	// authTokenManager  auth.AuthTokenManager
)

func init() {
	corsConfig = &mblibconfig.CORSConfig{
		AllowOrigins: []string{"*"},
	}
	logConfig = &mblibconfig.LogConfig{
		Enabled: map[string]bool{
			"accessLog": false,
			"traceLog":  false,
		},
	}
	serverConfig = &config.ServerConfig{
		HTTPPort:    8080,
		MetricsPort: 8081,
	}
	authConfig = &config.AuthConfig{
		SigningKey:          "ah5T9Y9V2JPU74fhCtHQfDqLp3Zg8ZNc",
		AccessTokenTTLMin:   1,
		RefreshTokenTTLHour: 1,
	}
	debugConfig = &libconfig.DebugConfig{
		Gin:  false,
		Wait: false,
	}
}

func initAuthRouter(t *testing.T, ctx context.Context, authentication controller.AuthenticationUsecase) *gin.Engine {
	t.Helper()
	fn := controller.NewInitAuthRouterFunc(authentication)

	initPublicRouterFuncs := []libcontroller.InitRouterGroupFunc{fn}
	// initPrivateRouterFuncs := []libcontroller.InitRouterGroupFunc{}

	router := libcontroller.InitRootRouterGroup(ctx, corsConfig, logConfig, debugConfig)
	api := router.Group("api")
	v1 := api.Group("v1")

	libcontroller.InitPublicAPIRouterGroup(ctx, v1, initPublicRouterFuncs)
	// if err := libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, initPrivateRouterFuncs); err != nil {
	// 	require.NoError(t, err)
	// }

	return router
}

func TestAuthHandler_GetUserInfo_shouldReturn401_whenAuthorizationHeaderIsEmpty(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authenticationUsecase := new(controllermock.MockAuthenticationUsecase)

	// given
	r := initAuthRouter(t, ctx, authenticationUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/auth/userinfo", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code, "status code should be 401")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 1, "response should have one message")
	assert.Equal(t, "Unauthorized", message[0], "message should be 'Unauthorized'")
}

func TestAuthHandler_GetUserInfo_shouldReturn401_whenAuthorizationHeaderIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authenticationUsecase := new(controllermock.MockAuthenticationUsecase)
	authenticationUsecase.On("GetUserInfo", anyOfCtx, "INVALID_TOKEN").Return(nil, errors.New("INVALID"))

	r := initAuthRouter(t, ctx, authenticationUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/auth/userinfo", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code, "status code should be 401")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 1, "response should have one message")
	assert.Equal(t, "Unauthorized", message[0], "message should be 'Unauthorized'")
}

func TestAuthHandler_GetUserInfo_shouldReturn200_whenAuthorizationHeaderIsValid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	appUserInfo := &mbuserdomain.AppUserModel{
		AppUserID:      appUserID(t, 123),
		OrganizationID: organizationID(t, 456),
		LoginID:        "LOGIN_ID",
		Username:       "USERNAME",
	}
	authenticationUsecase := new(controllermock.MockAuthenticationUsecase)
	authenticationUsecase.On("GetUserInfo", anyOfCtx, "VALID_TOKEN").Return(appUserInfo, nil)

	r := initAuthRouter(t, ctx, authenticationUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/auth/userinfo", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer VALID_TOKEN")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

	jsonObj := parseJSON(t, respBytes)

	appUserIDExpr := parseExpr(t, "$.appUserId")
	appUserID := appUserIDExpr.Get(jsonObj)
	assert.Equal(t, int64(123), appUserID[0])

	organizationIDExpr := parseExpr(t, "$.organizationId")
	organizationID := organizationIDExpr.Get(jsonObj)
	assert.Equal(t, int64(456), organizationID[0])

	loginIDExpr := parseExpr(t, "$.loginId")
	loginID := loginIDExpr.Get(jsonObj)
	assert.Equal(t, "LOGIN_ID", loginID[0])

	usernameExpr := parseExpr(t, "$.username")
	username := usernameExpr.Get(jsonObj)
	assert.Equal(t, "USERNAME", username[0])
}

func TestAuthHandler_RefreshToken_shouldReturn400_whenRequestBodyIsEmpty(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authenticationUsecase := new(controllermock.MockAuthenticationUsecase)
	r := initAuthRouter(t, ctx, authenticationUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/refresh-token", bytes.NewBuffer([]byte("")))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusBadRequest, w.Code, "status code should be 400")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 1, "response should have one message")
	assert.Equal(t, "Bad Request", message[0], "message should be 'Bad Request'")
}

func TestAuthHandler_RefreshToken_shouldReturn401_whenTokenIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authenticationUsecase := new(controllermock.MockAuthenticationUsecase)
	authenticationUsecase.On("RefreshToken", anyOfCtx, "INVALID_TOKEN").Return("", errors.New("INVALID"))

	r := initAuthRouter(t, ctx, authenticationUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/refresh-token", bytes.NewBuffer([]byte(`{"refreshToken": "INVALID_TOKEN"}`)))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code, "status code should be 401")

	jsonObj := parseJSON(t, respBytes)

	messageExpr := parseExpr(t, "$.message")
	message := messageExpr.Get(jsonObj)
	assert.Len(t, message, 1, "response should have one message")
	assert.Equal(t, "Unauthorized", message[0], "message should be 'Unauthorized'")
}

func TestAuthHandler_RefreshToken_shouldReturn200_whenTokenIsValid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authenticationUsecase := new(controllermock.MockAuthenticationUsecase)
	authenticationUsecase.On("RefreshToken", anyOfCtx, "VALID_TOKEN").Return("ACCESS_TOKEN", nil)

	r := initAuthRouter(t, ctx, authenticationUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/refresh-token", bytes.NewBuffer([]byte(`{"refreshToken": "VALID_TOKEN"}`)))
	require.NoError(t, err)
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

	jsonObj := parseJSON(t, respBytes)

	accessTokenExpr := parseExpr(t, "$.accessToken")
	accessToken := accessTokenExpr.Get(jsonObj)
	assert.Equal(t, "ACCESS_TOKEN", accessToken[0], "accessToken should be 'ACCESS_TOKEN'")
}
