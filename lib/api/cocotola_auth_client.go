package api

import (
	"context"

	apiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
)

type CocotolaAuthClient interface {
	RetrieveUserInfo(ctx context.Context, bearerToken string) (*apiauth.UserInfoResponse, error)
}
