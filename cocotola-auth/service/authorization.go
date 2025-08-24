package service

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func CheckAuthorization(ctx context.Context, operator OperatorInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject, nonTxManager TransactionManager) (bool, error) {
	return mblibservice.Do1(ctx, nonTxManager, func(rf RepositoryFactory) (bool, error) {
		mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		if err != nil {
			return false, err
		}

		authorizationManager, err := mbrf.NewAuthorizationManager(ctx)
		if err != nil {
			return false, err
		}

		return authorizationManager.CheckAuthorization(ctx, operator, action, object)
	})
}
