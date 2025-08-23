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

func (c *cocotolaRBACClient) Authorize(ctx context.Context, param *libapi.AuthorizeRequest) error {
	ctx, span := tracer.Start(ctx, "cocotolaRBACClient.Authorize")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "v1", "rbac", "authorize")

	jsonParam, err := json.Marshal(param)
	if err != nil {
		return mbliberrors.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), bytes.NewBuffer(jsonParam))
	if err != nil {
		return mbliberrors.Errorf("new http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.authUsername, c.authPassword)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return mbliberrors.Errorf("authorize request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return mbliberrors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
