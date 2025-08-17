//go:build small

package gateway_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	gatewaymock "github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway/mocks"
)

func newGoogleAuthClient(t *testing.T, httpClient gateway.HTTPClient) *gateway.GoogleAuthClient {
	t.Helper()
	return &gateway.GoogleAuthClient{
		HTTPClient:   httpClient,
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		RedirectURI:  "REDIRECT_URI",
		GrantType:    "authorization_code",
	}
}

func Test_GoogleAuthClient_NewGoogleAuthClient(t *testing.T) {
	t.Parallel()
	httpClient := new(gatewaymock.MockHTTPClient)
	c := gateway.NewGoogleAuthClient(
		httpClient, "CLIENT_ID", "CLIENT_SECRET", "REDIRECT_URI",
	)
	assert.Equal(t, "CLIENT_ID", c.ClientID)
	assert.Equal(t, "CLIENT_SECRET", c.ClientSecret)
	assert.Equal(t, "REDIRECT_URI", c.RedirectURI)
	assert.Equal(t, "authorization_code", c.GrantType)
	assert.Equal(t, httpClient, c.HTTPClient)
}

func Test_GoogleAuthClient_RetrieveAccessToken_shouldReturnTokenSet_whenReturnedStatusCodeIs200(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	httpClient := new(gatewaymock.MockHTTPClient)
	httpClient.On("Do", mock.Anything).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"access_token":"ACCESS_TOKEN","refresh_token":"REFRESH_TOKEN"}`)),
	}, nil)
	c := newGoogleAuthClient(t, httpClient)

	// when
	tokenSet, err := c.RetrieveAccessToken(ctx, "CODE")

	// then
	assert.Nil(t, err)
	assert.Equal(t, "ACCESS_TOKEN", tokenSet.AccessToken)
	assert.Equal(t, "REFRESH_TOKEN", tokenSet.RefreshToken)
}

func Test_GoogleAuthClient_RetrieveAccessToken_shouldReturnAuthenticationError_whenReturnedStatusCodeIs400(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	httpClient := new(gatewaymock.MockHTTPClient)
	httpClient.On("Do", mock.Anything).Return(&http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("")),
	}, nil)
	c := newGoogleAuthClient(t, httpClient)

	// when
	tokenSet, err := c.RetrieveAccessToken(ctx, "CODE")

	// then
	assert.ErrorIs(t, err, domain.ErrUnauthenticated)
	assert.Nil(t, tokenSet)
}

func Test_GoogleAuthClient_RetrieveAccessToken_shouldReturnAuthenticationError_whenReturnedStatusCodeIs401(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	httpClient := new(gatewaymock.MockHTTPClient)
	httpClient.On("Do", mock.Anything).Return(&http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(strings.NewReader("")),
	}, nil)
	c := newGoogleAuthClient(t, httpClient)

	// when
	tokenSet, err := c.RetrieveAccessToken(ctx, "CODE")

	// then
	assert.ErrorIs(t, err, domain.ErrUnauthenticated)
	assert.Nil(t, tokenSet)
}

func Test_GoogleAuthClient_RetrieveAccessToken_shouldReturnOtherError_whenErrorOccurred(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	otherError := errors.New("ERROR")
	httpClient := new(gatewaymock.MockHTTPClient)
	httpClient.On("Do", mock.Anything).Return(nil, otherError)
	c := newGoogleAuthClient(t, httpClient)

	// when
	tokenSet, err := c.RetrieveAccessToken(ctx, "CODE")

	// then
	assert.ErrorIs(t, err, otherError)
	assert.Nil(t, tokenSet)
}

func Test_GoogleAuthClient_RetrieveUserInfo(t *testing.T) {
	t.Skip()
	ctx := context.Background()
	accessToken := ""
	c := newGoogleAuthClient(t, http.DefaultClient)
	userInfo, err := c.RetrieveUserInfo(ctx, accessToken)
	require.NoError(t, err)
	assert.Equal(t, userInfo.Email, "")
}
