package gateway

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"path"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	"github.com/mocoarow/cocotola-1.24/lib/api"
	apiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
)

type cocotolaAuthClient struct {
	httpClient   HTTPClient
	authEndpoint *url.URL
	logger       *slog.Logger
}

func NewCocotolaAuthClient(httpClient HTTPClient, authEndpoint *url.URL) api.CocotolaAuthClient {
	return &cocotolaAuthClient{
		httpClient:   httpClient,
		authEndpoint: authEndpoint,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CocotolaAuthClient")),
	}
}

func (c *cocotolaAuthClient) GetUserInfo(ctx context.Context, bearerToken string) (*apiauth.GetUserInfoResponse, error) {
	ctx, span := tracer.Start(ctx, "cocotolaAuthClient.RetrieveUserInfo")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "api", "v1", "userinfo")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, mbliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, mbliberrors.Errorf("auth request: %w", err)
	}
	defer resp.Body.Close()

	var response apiauth.GetUserInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, mbliberrors.Errorf("json.NewDecoder: %w", err)
	}

	return &response, nil
}

func (c *cocotolaAuthClient) GetMyProfile(ctx context.Context, bearerToken string) (*apiauth.GetMyProfileResponse, error) {
	ctx, span := tracer.Start(ctx, "cocotolaAuthClient.GetMyProfile")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "api", "v1", "profile", "me")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, mbliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, mbliberrors.Errorf("auth request: %w", err)
	}
	defer resp.Body.Close()

	var response apiauth.GetMyProfileResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, mbliberrors.Errorf("json.NewDecoder: %w", err)
	}

	c.logger.InfoContext(ctx, "GetMyProfile response", slog.Any("response", response))

	return &response, nil
}
