package usecase

import (
	"context"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type SystemOwnerByOrganizationName interface {
	Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*mbuserservice.SystemOwner, error)
}
