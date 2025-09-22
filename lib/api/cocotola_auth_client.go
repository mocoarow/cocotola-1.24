package api

import (
	"context"

	apiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
)

type CocotolaAuthClient interface {
	GetUserInfo(ctx context.Context, bearerToken string) (*apiauth.GetUserInfoResponse, error)
	GetMyProfile(ctx context.Context, obearerToken string) (*apiauth.GetMyProfileResponse, error)
}
