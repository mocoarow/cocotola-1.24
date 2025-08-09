package service

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type OperatorInterface interface {
	AppUserID() *mbuserdomain.AppUserID
	OrganizationID() *mbuserdomain.OrganizationID
	// LoginID() string
	// Username() string
}
type RepositoryFactory interface {
	NewmoonbeamRepositoryFactory(ctx context.Context) (mbuserservice.RepositoryFactory, error)

	NewStateRepository(ctx context.Context) (StateRepository, error)
}
type TransactionManager mblibservice.TransactionManagerT[RepositoryFactory]
