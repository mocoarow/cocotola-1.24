package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type cocotolaAuthCallbackClient struct {
	httpClient   HTTPClient
	authEndpoint *url.URL
	authUsername string
	authPassword string
}

func NewCocotolaAuthCallbackClient(httpClient HTTPClient, authEndpoint *url.URL, authUsername, authPassword string) service.CocotolaAuthCallbackClient {
	return &cocotolaAuthCallbackClient{
		httpClient:   httpClient,
		authEndpoint: authEndpoint,
		authUsername: authUsername,
		authPassword: authPassword,
	}
}

func (c *cocotolaAuthCallbackClient) OnAddAppUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.AppUserID) error {
	ctx, span := tracer.Start(ctx, "cocotolaAuthCallbackClient.OnAddAppUser")
	defer span.End()

	u := *c.authEndpoint
	u.Path = path.Join(u.Path, "api", "v1", "callback", "on-add-user")

	apiReq := libapi.CallbackOnAddAppUserRequest{
		OrganizationID: organizationID.Int(),
		AppUserID:      appUserID.Int(),
	}
	jsonReq, err := json.Marshal(apiReq)
	if err != nil {
		return mbliberrors.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(jsonReq))
	if err != nil {
		return mbliberrors.Errorf("http.NewRequestWithContext. err: %w", err)
	}
	req.SetBasicAuth(c.authUsername, c.authPassword)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return mbliberrors.Errorf("auth request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return mbliberrors.Errorf("invalid status code: %d", resp.StatusCode)
	}

	return nil
}
