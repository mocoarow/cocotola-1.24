package api

import (
	"context"
)

type CocotolaAuthClient interface {
	RetrieveUserInfo(ctx context.Context, bearerToken string) (*AppUserInfoResponse, error)
}
