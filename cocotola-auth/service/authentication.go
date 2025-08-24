package service

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
)

func GetUserInfo(ctx context.Context, systemToken libdomain.SystemToken, authTokenManager AuthTokenManager, nonTxManager TransactionManager, bearerToken string) (*mbuserdomain.AppUserModel, error) {
	// TODO: Check whether the token is registered in the Database

	appUserInfo, err := authTokenManager.GetUserInfo(ctx, bearerToken)
	if err != nil {
		return nil, err
	}

	appUserModel, err := mblibservice.Do1(ctx, nonTxManager, func(rf RepositoryFactory) (*mbuserdomain.AppUserModel, error) {
		action, err := NewSystemOwnerAction(ctx, systemToken, rf,
			WithOrganizationByName(appUserInfo.OrganizationName),
		)
		if err != nil {
			return nil, mbliberrors.Errorf("new organization action: %w", err)
		}

		appUser, err := action.SystemOwner.FindAppUserByLoginID(ctx, appUserInfo.LoginID)
		if err != nil {
			return nil, mbliberrors.Errorf("find app user by login id(%s): %w", appUserInfo.LoginID, err)
		}

		return appUser.AppUserModel, nil
	})
	if err != nil {
		return nil, err
	}

	return appUserModel, nil
}
