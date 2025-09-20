package usecase

import (
	"context"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type SystemOwnerByOrganizationName interface {
	Get(ctx context.Context, rf mbuserservice.RepositoryFactory, organizationName string) (*mbuserservice.SystemOwner, error)
}
