package usecase

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type GuestAuthenticateCommand struct {
	mbTxManager      mbuserservice.TransactionManager
	mbNonTxManager   mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewGuestAuthenticateCommand(_ context.Context, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *GuestAuthenticateCommand {
	return &GuestAuthenticateCommand{
		mbTxManager:      mbTxManager,
		mbNonTxManager:   mbNonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *GuestAuthenticateCommand) Execute(ctx context.Context, systemOwner mbuserdomain.SystemOwnerInterface, organizationName string) (*domain.AuthTokenSet, error) {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, systemOwner); err != nil {
		return nil, mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	tokenSet, err := u.execute(ctx, systemOwner, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(ctx, systemOwner); err != nil {
		return nil, mbliberrors.Errorf("callback: %w", err)
	}
	return tokenSet, nil
}

func (u *GuestAuthenticateCommand) checkAuthorization(_ context.Context, _ mbuserdomain.SystemOwnerInterface) error {
	return nil
}

func (u *GuestAuthenticateCommand) execute(ctx context.Context, systemOwner mbuserdomain.SystemOwnerInterface, organizationName string) (*domain.AuthTokenSet, error) {
	guestLoginID := libdomain.NewGuestLoginID(organizationName)
	user, err := u.findUserbyLoginID(ctx, systemOwner, guestLoginID)
	if err != nil {
		if errors.Is(err, mbuserservice.ErrUserNotFound) {
			return nil, domain.ErrUnauthenticated
		}
		return nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	org, err := u.getOrganization(ctx, systemOwner)
	if err != nil {
		return nil, mbliberrors.Errorf("getOrganization: %w", err)
	}
	tokenSet, err := u.createTokenSet(ctx, user, org)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return tokenSet, nil
}

func (u *GuestAuthenticateCommand) callback(_ context.Context, _ mbuserdomain.SystemOwnerInterface) error {
	return nil
}

func (u *GuestAuthenticateCommand) findUserbyLoginID(ctx context.Context, operator mbuserdomain.UserInterface, loginID string) (*mbuserdomain.User, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.User, error) { //nolint:wrapcheck
		return findUserbyLoginID(ctx, mbrf, operator, loginID)
	})
}

func (u *GuestAuthenticateCommand) getOrganization(ctx context.Context, operator mbuserdomain.UserInterface) (*mbuserdomain.Organization, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.Organization, error) { //nolint:wrapcheck
		return getOrganization(ctx, mbrf, operator)
	})
}

func (u *GuestAuthenticateCommand) createTokenSet(ctx context.Context, userModel *mbuserdomain.User, organizationModel *mbuserdomain.Organization) (*domain.AuthTokenSet, error) {
	return createTokenSet(ctx, u.authTokenManager, userModel, organizationModel)
}
