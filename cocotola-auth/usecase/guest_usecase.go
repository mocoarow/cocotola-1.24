package usecase

import (
	"context"
	"errors"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type GuestUsecae struct {
	systemToken      libdomain.SystemToken
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewGuest(systemToken libdomain.SystemToken, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) *GuestUsecae {
	return &GuestUsecae{
		systemToken:      systemToken,
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *GuestUsecae) Authenticate(ctx context.Context, organizationName string) (*domain.AuthTokenSet, error) {
	var tokenSet *domain.AuthTokenSet

	targetOorganization, targetUser, err := mblibservice.Do2(ctx, u.txManager, func(rf service.RepositoryFactory) (*organization, *user, error) {
		action, err := service.NewSystemOwnerAction(ctx, u.systemToken, rf,
			service.WithOrganizationByName(organizationName),
		)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("new organization action: %w", err)
		} else if action.Organization == nil {
			return nil, nil, mbliberrors.Errorf("organization is nil")
		}

		guestLoginID := libdomain.NewGuestLoginID(organizationName)
		tmpUser, err := action.SystemOwner.FindUserByLoginID(ctx, guestLoginID)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("find user by login id: %w", err)
		}

		targetOorganization := &organization{
			organizationID: action.Organization.OrganizationID,
			name:           action.Organization.Name,
		}

		targetUser := &user{
			userID:         tmpUser.UserID,
			organizationID: tmpUser.OrganizationID,
			loginID:        tmpUser.LoginID,
			username:       tmpUser.Username,
		}

		return targetOorganization, targetUser, nil
	})
	if err != nil {
		if errors.Is(err, mbuserservice.ErrUserNotFound) {
			return nil, mbliberrors.Errorf("user not found: %w", domain.ErrUnauthenticated)
		}

		return nil, mbliberrors.Errorf("authenticate: %w", err)
	}

	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}
	tokenSet = tokenSetTmp

	return tokenSet, nil
}
