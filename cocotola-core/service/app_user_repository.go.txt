package service

import (
	"context"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type AppUserRepository interface {
	NewAppUser(ctx context.Context, operator mbuserservice.OperatorInterface) (*AppUser, error)
}
