package api

import (
	"context"
)

type CocotolaRBACClient interface {
	AddPolicyToUser(ctx context.Context, param *AddPolicyToUserParameter) error
	CheckAuthorization(ctx context.Context, param *AuthorizeRequest) (bool, error)
}
