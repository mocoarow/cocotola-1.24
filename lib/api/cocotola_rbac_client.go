package api

import (
	"context"

	apiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
)

type CocotolaRBACClient interface {
	AddPolicyToUser(ctx context.Context, param *apiauth.AddPolicyToUserParameter) error
	CheckAuthorization(ctx context.Context, param *apiauth.AuthorizeRequest) (bool, error)
}
