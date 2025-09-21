package service

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func CheckAuthorization(ctx context.Context, operator mbuserdomain.UserInterface, action mbuserdomain.RBACAction, object mbuserdomain.RBACObject, mbNonTxManager mbuserservice.TransactionManager) (bool, error) {
	return mblibservice.Do1(ctx, mbNonTxManager, func(rf mbuserservice.RepositoryFactory) (bool, error) { //nolint:wrapcheck
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return false, mbliberrors.Errorf("NewAuthorizationManager: %w", err)
		}

		return authorizationManager.CheckAuthorization(ctx, operator, action, object)
	})
}
