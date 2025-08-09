//go:build small

package controller_test

import (
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

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libconfig "github.com/mocoarow/cocotola-1.24/lib/config"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/middleware"
	controllermock "github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/mocks"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	servicemock "github.com/mocoarow/cocotola-1.24/cocotola-core/service/mocks"
)

var (
	anyOfCtx = mock.MatchedBy(func(_ context.Context) bool { return true })
	// corsConfig   cors.Config
	corsConfig   *mblibconfig.CORSConfig
	serverConfig *config.ServerConfig
	debugConfig  *libconfig.DebugConfig
	// authTokenManager  auth.AuthTokenManager
)

func init() {
	corsConfig = &mblibconfig.CORSConfig{
		AllowOrigins: []string{"*"},
	}
	serverConfig = &config.ServerConfig{
		HTTPPort:    8080,
		MetricsPort: 8081,
	}
	debugConfig = &libconfig.DebugConfig{
		Gin:  false,
		Wait: false,
	}
}

func initWorkbookRouter(t *testing.T, ctx context.Context, cocotolaAuthClient service.CocotolaAuthClient, workbokQueryUsecase controller.WorkbookQueryUsecase, workbookCommandUsecase controller.WorkbookCommandUsecase) *gin.Engine {
	t.Helper()
	fn := controller.NewInitWorkbookRouterFunc(workbokQueryUsecase, workbookCommandUsecase)

	authMiddleware := middleware.NewAuthMiddleware(cocotolaAuthClient)
	initPublicRouterFuncs := []libcontroller.InitRouterGroupFunc{}
	initPrivateRouterFuncs := []libcontroller.InitRouterGroupFunc{fn}

	router := libcontroller.InitRootRouterGroup(ctx, corsConfig, debugConfig)
	api := router.Group("api")
	v1 := api.Group("v1")

	libcontroller.InitPublicAPIRouterGroup(ctx, v1, initPublicRouterFuncs)
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, authMiddleware, initPrivateRouterFuncs)
	// err := initialize.InitAppServer(ctx, router, authMiddleware, corsConfig, debugConfig, appConfig.Name, initPublicRouterFunc, initPrivateRouterFunc)
	// require.NoError(t, err)

	return router
}

func TestWorkbookHandler_FindWorkbook_shouldReturn200(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	cocotolaAuthClient := new(servicemock.CocotolaAuthClient)
	cocotolaAuthClient.On("RetrieveUserInfo", anyOfCtx, "VALID_TOKEN").Return(&libapi.AppUserInfoResponse{
		AppUserID:      123,
		OrganizationID: 456,
		LoginID:        "LOGIN_ID",
		Username:       "USERNAME",
	}, nil)

	workbookQueryUsecase := new(controllermock.WorkbookQueryUsecase)
	workbookQueryUsecase.On("FindWorkbooks", anyOfCtx, mock.MatchedBy(func(o service.OperatorInterface) bool {
		return o.OrganizationID().Int() == 456 && o.AppUserID().Int() == 123
	}), mock.Anything).Return(&libapi.WorkbookFindResult{
		TotalCount: 789,
		Results: []*libapi.WorkbookFindWorkbookModel{
			{
				ID:   135,
				Name: "WORKBOOK_NAME",
			},
		},
	}, nil)
	workbookCommandUsecase := new(controllermock.WorkbookCommandUsecase)

	// given
	r := initWorkbookRouter(t, ctx, cocotolaAuthClient, workbookQueryUsecase, workbookCommandUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/workbook", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer VALID_TOKEN")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

	jsonObj := parseJSON(t, respBytes)

	totalCountExpr := parseExpr(t, "$.totalCount")
	totalCount := totalCountExpr.Get(jsonObj)
	assert.Len(t, totalCount, 1, "response should have one totalCount")
	assert.Equal(t, int64(789), totalCount[0], "totalCount should be 789")

	resultsExpr := parseExpr(t, "$.results")
	results := resultsExpr.Get(jsonObj)
	assert.Len(t, results, 1, "response should have one results")
	assert.Equal(t, 1, len(results), "results should have one element")

	resultID0Expr := parseExpr(t, "$.results[0].id")
	resultID0 := resultID0Expr.Get(jsonObj)
	assert.Equal(t, int64(135), resultID0[0], "results[0].id should be 135")

	resultName0Expr := parseExpr(t, "$.results[0].name")
	resultName0 := resultName0Expr.Get(jsonObj)
	assert.Equal(t, "WORKBOOK_NAME", resultName0[0], "results[0].name should be 'WORKBOOK_NAME'")
}

func TestWorkbookHandler_FindWorkbook_shouldReturn401_whenAuthorizationHeaderIsEmpty(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	cocotolaAuthClient := new(servicemock.CocotolaAuthClient)
	workbookQueryUsecase := new(controllermock.WorkbookQueryUsecase)
	workbookCommandUsecase := new(controllermock.WorkbookCommandUsecase)

	// given
	r := initWorkbookRouter(t, ctx, cocotolaAuthClient, workbookQueryUsecase, workbookCommandUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/workbook", nil)
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

func TestWorkbookHandler_FindWorkbook_shouldReturn401_whenAuthorizationHeaderIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	cocotolaAuthClient := new(servicemock.CocotolaAuthClient)
	cocotolaAuthClient.On("RetrieveUserInfo", anyOfCtx, "INVALID_TOKEN").Return(nil, errors.New("invalid token"))
	workbookQueryUsecase := new(controllermock.WorkbookQueryUsecase)
	workbookCommandUsecase := new(controllermock.WorkbookCommandUsecase)

	// given
	r := initWorkbookRouter(t, ctx, cocotolaAuthClient, workbookQueryUsecase, workbookCommandUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/workbook", nil)
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

// func TestWorkbookHandler_RetrieveWorkbookByID_shouldReturn200(t *testing.T) {
// 	t.Parallel()
// 	ctx := context.Background()

// 	// given
// 	cocotolaAuthClient := new(servicemock.CocotolaAuthClient)
// 	cocotolaAuthClient.On("RetrieveUserInfo", anyOfCtx, "VALID_TOKEN").Return(&libapi.AppUserInfoResponse{
// 		AppUserID:      123,
// 		OrganizationID: 456,
// 		LoginID:        "LOGIN_ID",
// 		Username:       "USERNAME",
// 	}, nil)

// 	workbookQueryUsecase := new(controllermock.WorkbookQueryUsecase)
// 	workbookQueryUsecase.On("RetrieveWorkbookByID", anyOfCtx, organizationID(t, 456), appUserID(t, 123), 246).Return(&workbookfind.Result{
// 		TotalCount: 789,
// 		Results: []*workbookfind.WorkbookModel{
// 			{
// 				ID:   135,
// 				Name: "WORKBOOK_NAME",
// 			},
// 		},
// 	}, nil)
// 	workbookCommandUsecase := new(controllermock.WorkbookCommandUsecase)

// 	// given
// 	r := initWorkbookRouter(t, ctx, cocotolaAuthClient, workbookQueryUsecase, workbookCommandUsecase)
// 	w := httptest.NewRecorder()

// 	// when
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/v1/workbook", nil)
// 	require.NoError(t, err)
// 	req.Header.Set("Authorization", "Bearer VALID_TOKEN")
// 	r.ServeHTTP(w, req)
// 	respBytes := readBytes(t, w.Body)

// 	// then
// 	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")

// 	jsonObj := parseJSON(t, respBytes)

// 	totalCountExpr := parseExpr(t, "$.totalCount")
// 	totalCount := totalCountExpr.Get(jsonObj)
// 	assert.Equal(t, int64(789), totalCount[0], "totalCount should be 789")

// 	resultsExpr := parseExpr(t, "$.results")
// 	results := resultsExpr.Get(jsonObj)
// 	assert.Equal(t, 1, len(results), "results should have one element")

// 	resultID0Expr := parseExpr(t, "$.results[0].id")
// 	resultID0 := resultID0Expr.Get(jsonObj)
// 	assert.Equal(t, int64(135), resultID0[0], "results[0].id should be 135")

// 	resultName0Expr := parseExpr(t, "$.results[0].name")
// 	resultName0 := resultName0Expr.Get(jsonObj)
// 	assert.Equal(t, "WORKBOOK_NAME", resultName0[0], "results[0].name should be 'WORKBOOK_NAME'")
// }
