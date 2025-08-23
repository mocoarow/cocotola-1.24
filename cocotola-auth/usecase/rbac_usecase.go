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

func (u *RBACUsecase) AddPolicyToUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, subject mbuserdomain.RBACSubject, listOfActionObjectEffect []mbuserdomain.RBACActionObjectEffect) error {
	return mblibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		rsrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
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

		for _, aoe := range listOfActionObjectEffect {
			action := aoe.Action
			object := aoe.Object
			effect := aoe.Effect
			if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAdmin, organizationID, subject, action, object, effect); err != nil {
				return err
			}
		}

		return nil
	})
}

func (u *RBACUsecase) Authorize(ctx context.Context, operator service.OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject) (bool, error) {
	return service.Authorize(ctx, operator, action, object, u.nonTxManager)
}
