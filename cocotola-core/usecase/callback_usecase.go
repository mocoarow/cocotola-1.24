package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type CallbackUsecase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   libapi.CocotolaRBACClient
	logger       *slog.Logger
}

func NewCallbackUsecase(txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) *CallbackUsecase {
	return &CallbackUsecase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

func (u *CallbackUsecase) OnAddUserSpace(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID, spaceID *mbuserdomain.SpaceID) error {
	ctx, span := tracer.Start(ctx, "CallbackUsecase.OnAddUserSpace")
	defer span.End()

	operator := Operator{
		organizationID: organizationID,
		userID:         userID,
	}
	command := NewCallbackOnAddUserSpaceCommand(u.txManager, u.nonTxManager, u.rbacClient)
	if err := command.Execute(ctx, &operator, spaceID); err != nil {
		return mbliberrors.Errorf("command.Execute: %w", err)
	}

	return nil
}
