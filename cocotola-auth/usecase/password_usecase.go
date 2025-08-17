package usecase

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type PasswordUsecae struct {
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPassword(txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordUsecae {
	return &PasswordUsecae{
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordUsecae) Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	var tokenSet *domain.AuthTokenSet

	targetOorganization, targetAppUser, err := mblibservice.Do2(ctx, u.txManager, func(rf service.RepositoryFactory) (*organization, *appUser, error) {
		action, err := NewOrganizationAction(ctx, rf,
			WithOrganizationRepository(),
			WithOrganization(organizationName),
			WithAppUserRepository(),
		)
		if err != nil {
			return nil, nil, err
		}

		verified, err := action.appUserRepo.VerifyPassword(ctx, action.systemAdmin, action.organization.OrganizationModel.OrganizationID, loginID, password)
		if err != nil {
			return nil, nil, err
		}

		if !verified {
			return nil, nil, domain.ErrUnauthenticated
		}

		tmpAppUser, err := action.appUserRepo.FindAppUserByLoginID(ctx, action.systemOwner, loginID)
		if err != nil {
			return nil, nil, err
		}

		targetOorganization := &organization{
			organizationID: action.organization.OrganizationModel.OrganizationID,
			name:           action.organization.OrganizationModel.Name,
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
			return nil, mbliberrors.Errorf("AppUserNotFound. err: %w", domain.ErrUnauthenticated)
		}
		return nil, mbliberrors.Errorf("RegisterAppUser. err: %w", err)
	}

	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetAppUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}
	tokenSet = tokenSetTmp
	return tokenSet, nil
}
