package usecase

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type PasswordUsecae struct {
	systemToken      libdomain.SystemToken
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPassword(systemToken libdomain.SystemToken, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordUsecae {
	return &PasswordUsecae{
		systemToken:      systemToken,
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordUsecae) Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	var tokenSet *domain.AuthTokenSet

	targetOorganization, targetAppUser, err := mblibservice.Do2(ctx, u.txManager, func(rf service.RepositoryFactory) (*organization, *appUser, error) {
		action, err := service.NewSystemOwnerAction(ctx, u.systemToken, rf,
			// service.WithOrganizationRepository(),
			service.WithOrganizationByName(organizationName),
			// service.WithAppUserRepository(),
		)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("new organization action: %w", err)
		}
		// if action.AppUserRepo == nil {
		// 	return nil, nil, mbliberrors.Errorf("app user repository is nil")
		// }
		if action.Organization == nil {
			return nil, nil, mbliberrors.Errorf("organization is nil")
		}

		verified, err := action.SystemOwner.VerifyPassword(ctx, loginID, password)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("action.appUserRepo.VerifyPassword: %w", err)
		}

		if !verified {
			return nil, nil, domain.ErrUnauthenticated
		}

		tmpAppUser, err := action.SystemOwner.FindAppUserByLoginID(ctx, loginID)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("find app user by login id: %w", err)
		}

		targetOorganization := &organization{
			organizationID: action.Organization.OrganizationModel.OrganizationID,
			name:           action.Organization.OrganizationModel.Name,
		}

		targetAppUser := &appUser{
			appUserID:      tmpAppUser.AppUserModel.AppUserID,
			organizationID: tmpAppUser.AppUserModel.OrganizationID,
			loginID:        tmpAppUser.AppUserModel.LoginID,
			username:       tmpAppUser.AppUserModel.Username,
		}

		return targetOorganization, targetAppUser, nil
	})

	if err != nil {
		if errors.Is(err, mbuserservice.ErrAppUserNotFound) {
			return nil, mbliberrors.Errorf("app user not found: %w", domain.ErrUnauthenticated)
		}
		return nil, mbliberrors.Errorf("authenticate: %w", err)
	}

	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetAppUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}
	tokenSet = tokenSetTmp
	return tokenSet, nil
}
