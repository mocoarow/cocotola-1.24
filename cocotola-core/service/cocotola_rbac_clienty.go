package service

import (
	"context"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
)

type CocotolaRBACClient interface {
	AddPolicyToUser(ctx context.Context, param *libapi.AddPolicyToUserParameter) error
}
