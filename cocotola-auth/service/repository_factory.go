package service

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
)

//	type OperatorInterface interface {
//		UserID() *mbuserdomain.UserID
//		OrganizationID() *mbuserdomain.OrganizationID
//		// LoginID() string
//		// Username() string
//	}
type RepositoryFactory interface {
	// NewMoonBeamRepositoryFactory(ctx context.Context) (mbuserservice.RepositoryFactory, error)

	NewStateRepository(ctx context.Context) (StateRepository, error)
}

type TransactionManager mblibservice.TransactionManagerT[RepositoryFactory]
