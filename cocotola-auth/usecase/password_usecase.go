package usecase

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type PasswordUsecae struct {
	systemToken      libdomain.SystemToken
	txManager        mbuserservice.TransactionManager
	nonTxManager     mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPassword(systemToken libdomain.SystemToken, txManager, nonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordUsecae {
	return &PasswordUsecae{
		systemToken:      systemToken,
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordUsecae) Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	sysAdmin := service.NewSystemAdmin(u.systemToken)

	command := NewPasswordAuthenticateCommand(ctx, u.txManager, u.nonTxManager, u.authTokenManager)
	tokenSet, err := command.Execute(ctx, sysAdmin, loginID, password, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("command.Execute: %w", err)
	}
	return tokenSet, nil
}
