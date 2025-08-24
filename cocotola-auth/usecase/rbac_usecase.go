package usecase

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
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
	return mblibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error { //nolint:wrapcheck
		mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		}

		sysAdmin, err := mbuserservice.NewSystemAdmin(ctx, mbrf)
		if err != nil {
			return mbliberrors.Errorf("NewSystemAdmin: %w", err)
		}

		authorizationManager, err := mbrf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewAuthorizationManager: %w", err)
		}

		for _, aoe := range listOfActionObjectEffect {
			action := aoe.Action
			object := aoe.Object
			effect := aoe.Effect
			if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAdmin, organizationID, subject, action, object, effect); err != nil {
				return mbliberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
			}
		}

		return nil
	})
}

func (u *RBACUsecase) CheckAuthorization(ctx context.Context, operator service.OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject) (bool, error) {
	ok, err := service.CheckAuthorization(ctx, operator, action, object, u.nonTxManager)
	if err != nil {
		return false, mbliberrors.Errorf("CheckAuthorization: %w", err)
	}

	return ok, nil
}
