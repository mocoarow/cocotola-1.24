package service

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
)

func GetUserInfo(ctx context.Context, systemToken libdomain.SystemToken, authTokenManager AuthTokenManager, nonTxManager TransactionManager, bearerToken string) (*mbuserdomain.UserModel, error) {
	// TODO: Check whether the token is registered in the Database
	userInfo, err := authTokenManager.GetUserInfo(ctx, bearerToken)
	if err != nil {
		return nil, mbliberrors.Errorf("GetUserInfo: %w", err)
	}

	userModel, err := mblibservice.Do1(ctx, nonTxManager, func(rf RepositoryFactory) (*mbuserdomain.UserModel, error) {
		action, err := NewSystemOwnerAction(ctx, systemToken, rf,
			WithOrganizationByName(userInfo.OrganizationName),
		)
		if err != nil {
			return nil, mbliberrors.Errorf("new organization action: %w", err)
		}

		user, err := action.SystemOwner.FindUserByLoginID(ctx, userInfo.LoginID)
		if err != nil {
			return nil, mbliberrors.Errorf("find user by login id(%s): %w", userInfo.LoginID, err)
		}

		return user.UserModel, nil
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return userModel, nil
}
