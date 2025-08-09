package usecase

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type RBACUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewRBACUsecase(txManager, nonTxManager service.TransactionManager) *RBACUsecase {
	return &RBACUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *RBACUsecase) AddPolicyToUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, subject mbuserdomain.RBACSubject, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject, effect mbuserdomain.RBACEffect) error {
	return mblibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewmoonbeamRepositoryFactory(ctx)
		if err != nil {
			return err
		}

		sysAdmin, err := mbuserservice.NewSystemAdmin(ctx, rsrf)
		if err != nil {
			return err
		}

		authorizationManager, err := rsrf.NewAuthorizationManager(ctx)
		if err != nil {
			return err
		}

		if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAdmin, organizationID, subject, action, object, effect); err != nil {
			return err
		}

		return nil
	})
}

func (u *RBACUsecase) Authorize(ctx context.Context, operator service.OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject) (bool, error) {
	return mblibservice.Do1(ctx, u.nonTxManager, func(rf service.RepositoryFactory) (bool, error) {
		rsrf, err := rf.NewmoonbeamRepositoryFactory(ctx)
		if err != nil {
			return false, err
		}

		authorizationManager, err := rsrf.NewAuthorizationManager(ctx)
		if err != nil {
			return false, err
		}

		return authorizationManager.Authorize(ctx, operator, action, object)
	})
}
