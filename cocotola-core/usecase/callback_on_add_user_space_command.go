package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type CallbackOnAddUserSpaceCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   libapi.CocotolaRBACClient
	logger       *slog.Logger
}

func NewCallbackOnAddUserSpaceCommand(txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) *CallbackOnAddUserSpaceCommand {
	return &CallbackOnAddUserSpaceCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackOnAddUserSpaceCommandUsecase")),
	}
}

func (u *CallbackOnAddUserSpaceCommand) Execute(ctx context.Context, operator mbuserdomain.UserInterface, spaceID *mbuserdomain.SpaceID) error {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator); err != nil {
		return mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	if err := u.execute(ctx, operator, spaceID); err != nil {
		return mbliberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(); err != nil {
		return mbliberrors.Errorf("callback: %w", err)
	}

	return nil
}

func (u *CallbackOnAddUserSpaceCommand) checkAuthorization(_ context.Context, _ mbuserdomain.UserInterface) error {
	return nil
}

func (u *CallbackOnAddUserSpaceCommand) execute(ctx context.Context, operator mbuserdomain.UserInterface, spaceID *mbuserdomain.SpaceID) error {
	fn := func(rf service.RepositoryFactory) error {
		folderRepo := rf.NewFolderRepository(ctx)

		param := service.AddFolderParameter{
			SpaceID:  spaceID,
			FolderID: domain.EmptyFolderID,
			Name:     "Root",
		}
		spaceID, err := folderRepo.AddFolder(ctx, operator, &param)
		if err != nil {
			return mbliberrors.Errorf("AddFolder: %w", err)
		}

		u.logger.InfoContext(ctx, "OnAddUser: AddSpace", slog.Int("space_id", spaceID.Int()))
		return nil
	}
	if err := mblibservice.Do0(ctx, u.nonTxManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}

func (u *CallbackOnAddUserSpaceCommand) callback() error {
	return nil
}
