package service

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
)

type RepositoryFactory interface {
	NewOrganizationRepository(ctx context.Context) OrganizationRepository
	NewAppUserRepository(ctx context.Context) AppUserRepository
	NewUserGroupRepository(ctx context.Context) UserGroupRepository

	// NewPairOfUserAndGroupRepository(ctx context.Context) PairOfUserAndGroupRepository

	// NewRBACRepository(ctx context.Context) RBACRepository

	NewAuthorizationManager(ctx context.Context) (AuthorizationManager, error)
	NewAppUserEventHandler(ctx context.Context) mblibservice.ResourceEventHandler
}

type TransactionManager mblibservice.TransactionManagerT[RepositoryFactory]
