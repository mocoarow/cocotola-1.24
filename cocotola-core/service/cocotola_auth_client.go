package service

import (
	"context"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
)

type CocotolaAuthClient interface {
	RetrieveUserInfo(ctx context.Context, bearerToken string) (*libapi.AppUserInfoResponse, error)
}
