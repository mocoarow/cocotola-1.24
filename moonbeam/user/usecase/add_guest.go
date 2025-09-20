package usecase

import (
	"context"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type AddGuestCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewAddGuestCommand(txManager service.TransactionManager, nonTxManager service.TransactionManager) *AddGuestCommand {
	return &AddGuestCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *AddGuestCommand) Execute(ctx context.Context, operator service.SystemOwnerInterface, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator); err != nil {
		return nil, liberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	newUserID, err := u.execute(ctx, operator, param, aoeList)
	if err != nil {
		return nil, liberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(ctx, operator, newUserID); err != nil {
		return nil, liberrors.Errorf("callback: %w", err)
	}

	return newUserID, nil
}

func (u *AddGuestCommand) checkAuthorization(ctx context.Context, operator service.SystemOwnerInterface) error {
	return nil
}

func (u *AddGuestCommand) execute(ctx context.Context, operator service.SystemOwnerInterface, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	userID, err := libservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.UserID, error) {
		return AddUser(ctx, operator, rf, param, aoeList)
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return userID, nil
}

func (u *AddGuestCommand) callback(ctx context.Context, operator service.SystemOwnerInterface, newUserID *domain.UserID) error {
	return nil
}
