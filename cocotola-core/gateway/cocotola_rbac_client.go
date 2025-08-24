package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type cocotolaRBACClient struct {
	httpClient   HTTPClient
	authEndpoint *url.URL
	authUsername string
	authPassword string
}

func NewCocotolaRBACClient(httpClient HTTPClient, authEndpoint *url.URL, authUsername, authPassword string) service.CocotolaRBACClient {
	return &cocotolaRBACClient{
		httpClient:   httpClient,
		authEndpoint: authEndpoint,
		authUsername: authUsername,
		authPassword: authPassword,
	}
}

func (c *cocotolaRBACClient) AddPolicyToUser(ctx context.Context, param *libapi.AddPolicyToUserParameter) error {
	ctx, span := tracer.Start(ctx, "cocotolaRBACClient.AddPolicyToUser")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "v1", "rbac", "policy", "user")

	jsonParam, err := json.Marshal(param)
	if err != nil {
		return mbliberrors.Errorf("json.Marshal. err: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), bytes.NewBuffer(jsonParam))
	if err != nil {
		return mbliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.authUsername, c.authPassword)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return mbliberrors.Errorf("auth request. err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return mbliberrors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *cocotolaRBACClient) CheckAuthorization(ctx context.Context, param *libapi.AuthorizeRequest) (bool, error) {
	ctx, span := tracer.Start(ctx, "cocotolaRBACClient.CheckAuthorization")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "v1", "rbac", "check-authorization")

	jsonParam, err := json.Marshal(param)
	if err != nil {
		return false, mbliberrors.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(jsonParam))
	if err != nil {
		return false, mbliberrors.Errorf("new http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.authUsername, c.authPassword)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, mbliberrors.Errorf("check-authorization request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, mbliberrors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	apiResp := libapi.AuthorizeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return false, mbliberrors.Errorf("decode response body: %w", err)
	}
	return apiResp.Authorized, nil
}
