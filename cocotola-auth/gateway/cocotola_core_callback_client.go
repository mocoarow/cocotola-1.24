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

type cocotolaCoreCallbackClient struct {
	httpClient   HTTPClient
	coreEndpoint *url.URL
	coreUsername string
	corePassword string
}

func NewCocotolaCoreCallbackClient(httpClient HTTPClient, coreEndpoint *url.URL, coreUsername, corePassword string) service.CocotolaCoreCallbackClient {
	return &cocotolaCoreCallbackClient{
		httpClient:   httpClient,
		coreEndpoint: coreEndpoint,
		coreUsername: coreUsername,
		corePassword: corePassword,
	}
}

func (c *cocotolaCoreCallbackClient) OnAddAppUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.AppUserID) error {
	ctx, span := tracer.Start(ctx, "cocotolaAuthClient.RetrieveUserInfo")
	defer span.End()

	u := *c.coreEndpoint
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
	req.SetBasicAuth(c.coreUsername, c.corePassword)

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
