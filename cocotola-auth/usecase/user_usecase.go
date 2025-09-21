package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type UserUsecase struct {
	systemToken      libdomain.SystemToken
	mbTxManager      mbuserservice.TransactionManager
	mbNonTxManager   mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
	logger           *slog.Logger
}

func NewUserUsecase(systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *UserUsecase {
	return &UserUsecase{
		systemToken:      systemToken,
		mbTxManager:      mbTxManager,
		mbNonTxManager:   mbNonTxManager,
		authTokenManager: authTokenManager,
		logger:           slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-UserUsecase")),
	}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, operator mbuserdomain.UserInterface, param *mbuserservice.AddUserParameter) (*domain.AuthTokenSet, error) {
	command, err := NewRegisterUserCommand(ctx, u.mbTxManager, u.mbNonTxManager, u.authTokenManager)
	if err != nil {
		return nil, mbliberrors.Errorf("NewRegisterUserCommand. err: %w", err)
	}

	tokenSet, err := command.Execute(ctx, operator, param)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}

	return tokenSet, nil
}
