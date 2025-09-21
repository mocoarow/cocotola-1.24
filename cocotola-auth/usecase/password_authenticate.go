package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type PasswordAuthenticateCommand struct {
	mbTxManager      mbuserservice.TransactionManager
	mbNonTxManager   mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPasswordAuthenticateCommand(ctx context.Context, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordAuthenticateCommand {
	return &PasswordAuthenticateCommand{
		mbTxManager:      mbTxManager,
		mbNonTxManager:   mbNonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordAuthenticateCommand) Execute(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, systemAdmin, loginID); err != nil {
		return nil, mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	tokenSet, err := u.execute(ctx, systemAdmin, loginID, password, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(ctx, systemAdmin); err != nil {
		return nil, mbliberrors.Errorf("callback: %w", err)
	}
	return tokenSet, nil
}

func (u *PasswordAuthenticateCommand) checkAuthorization(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, loginID string) error {
	if strings.Contains(loginID, "guest@@") {
		return domain.ErrUnauthenticated
	}

	return nil
}
func (u *PasswordAuthenticateCommand) execute(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	sysOwner, err := u.findSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("findSystemOwnerByOrganizationName: %w", err)
	}

	fn := func(rf mbuserservice.RepositoryFactory) error {
		userRepo := rf.NewUserRepository(ctx)
		ok, err := userRepo.VerifyPassword(ctx, sysOwner, loginID, password)
		if err != nil {
			return mbliberrors.Errorf("action.userRepo.VerifyPassword: %w", err)
		} else if !ok {
			return domain.ErrUnauthenticated
		}
		return nil
	}
	if err := mblibservice.Do0(ctx, u.mbNonTxManager, fn); err != nil {
		return nil, err //nolint:wrapcheck
	}

	user, err := u.findUserbyLoginID(ctx, sysOwner, loginID)
	if err != nil {
		if errors.Is(err, mbuserservice.ErrUserNotFound) {
			return nil, domain.ErrUnauthenticated
		}
		return nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	org, err := u.getOrganization(ctx, sysOwner)
	if err != nil {
		return nil, mbliberrors.Errorf("getOrganization: %w", err)
	}
	tokenSet, err := u.createTokenSet(ctx, user, org)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return tokenSet, nil
}

func (u *PasswordAuthenticateCommand) callback(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface) error {
	return nil
}

func (u *PasswordAuthenticateCommand) findSystemOwnerByOrganizationName(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationName string) (*mbuserdomain.SystemOwner, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) {
		return findSystemOwnerByOrganizationName(ctx, mbrf, operator, organizationName)
	})
}
func (u *PasswordAuthenticateCommand) findUserbyLoginID(ctx context.Context, operator mbuserdomain.UserInterface, loginID string) (*mbuserdomain.User, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.User, error) {
		return findUserbyLoginID(ctx, mbrf, operator, loginID)
	})
}

func (u *PasswordAuthenticateCommand) getOrganization(ctx context.Context, operator mbuserdomain.UserInterface) (*mbuserdomain.Organization, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.Organization, error) {
		return getOrganization(ctx, mbrf, operator)
	})
}

func (u *PasswordAuthenticateCommand) createTokenSet(ctx context.Context, userModel *mbuserdomain.User, organizationModel *mbuserdomain.Organization) (*domain.AuthTokenSet, error) {
	return createTokenSet(ctx, u.authTokenManager, userModel, organizationModel)
}
