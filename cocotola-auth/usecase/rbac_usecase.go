package usecase

import (
	"context"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type RBACUsecase struct {
	systemToken  libdomain.SystemToken
	txManager    mbuserservice.TransactionManager
	nonTxManager mbuserservice.TransactionManager
}

func NewRBACUsecase(systemToken libdomain.SystemToken, txManager, nonTxManager mbuserservice.TransactionManager) *RBACUsecase {
	return &RBACUsecase{
		systemToken:  systemToken,
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *RBACUsecase) AddPolicyToUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, subject mbuserdomain.RBACSubject, listOfActionObjectEffect []mbuserdomain.RBACActionObjectEffect) error {
	return mblibservice.Do0(ctx, u.txManager, func(rf mbuserservice.RepositoryFactory) error { //nolint:wrapcheck
		// mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		// if err != nil {
		// 	return mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		// }

		sysAdmin := service.NewSystemAdmin(u.systemToken)

		authorizationManager, err := rf.NewAuthorizationManager(ctx)
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

func (u *RBACUsecase) CheckAuthorization(ctx context.Context, operator mbuserservice.OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject) (bool, error) {
	ok, err := service.CheckAuthorization(ctx, operator, action, object, u.nonTxManager)
	if err != nil {
		return false, mbliberrors.Errorf("CheckAuthorization: %w", err)
	}

	return ok, nil
}
