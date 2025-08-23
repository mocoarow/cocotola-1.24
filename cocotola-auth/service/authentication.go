package service

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
)

func GetUserInfo(ctx context.Context, systemToken libdomain.SystemToken, authTokenManager AuthTokenManager, nonTxManager TransactionManager, bearerToken string) (*mbuserdomain.AppUserModel, error) {
	// TODO: Check whether the token is registered in the Database

	appUserInfo, err := authTokenManager.GetUserInfo(ctx, bearerToken)
	if err != nil {
		return nil, err
	}

	var targetAppUserModel *mbuserdomain.AppUserModel

	if err := nonTxManager.Do(ctx, func(rf RepositoryFactory) error {
		action, err := NewOrganizationAction(ctx, systemToken, rf,
			WithOrganizationByName(appUserInfo.OrganizationName),
		)
		if err != nil {
			return mbliberrors.Errorf("new organization action: %w", err)
		}

		appUser, err := action.SystemOwner.FindAppUserByLoginID(ctx, appUserInfo.LoginID)
		if err != nil {
			return err
		}

		targetAppUserModel = appUser.AppUserModel
		return nil
	}); err != nil {
		return nil, mbliberrors.Errorf("RegisterAppUser. err: %w", err)
	}

	return targetAppUserModel, nil
}
