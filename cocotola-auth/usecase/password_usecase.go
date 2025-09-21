package usecase

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type PasswordUsecae struct {
	systemToken      libdomain.SystemToken
	mbTxManager      mbuserservice.TransactionManager
	mbNonTxManager   mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPassword(systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordUsecae {
	return &PasswordUsecae{
		systemToken:      systemToken,
		mbTxManager:      mbTxManager,
		mbNonTxManager:   mbNonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordUsecae) Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	sysAdmin := service.NewSystemAdmin(u.systemToken)
	sysOwner, err := u.findSystemOwnerByOrganizationName(ctx, sysAdmin, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("Do1: %w", err)
	}

	command := NewPasswordAuthenticateCommand(ctx, u.mbTxManager, u.mbNonTxManager, u.authTokenManager)
	tokenSet, err := command.Execute(ctx, sysOwner, loginID, password)
	if err != nil {
		return nil, mbliberrors.Errorf("command.Execute: %w", err)
	}
	return tokenSet, nil
}

func (u *PasswordUsecae) findSystemOwnerByOrganizationName(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationName string) (*mbuserdomain.SystemOwner, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) { //nolint:wrapcheck
		return findSystemOwnerByOrganizationName(ctx, mbrf, operator, organizationName)
	})
}
