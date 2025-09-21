package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type Callback struct {
	systemToken                libdomain.SystemToken
	txManager                  mbuserservice.TransactionManager
	nonTxManager               mbuserservice.TransactionManager
	cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient
	logger                     *slog.Logger
}

func NewCallback(systemToken libdomain.SystemToken, txManager, nonTxManager mbuserservice.TransactionManager, cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient) *Callback {
	return &Callback{
		systemToken:                systemToken,
		txManager:                  txManager,
		nonTxManager:               nonTxManager,
		cocotolaCoreCallbackClient: cocotolaCoreCallbackClient,
		logger:                     slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

func (u *Callback) OnAddUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error {
	u.logger.InfoContext(ctx, "OnAddUser", slog.Int("user_id", userID.Int()))

	sysAdmin := service.NewSystemAdmin(u.systemToken)
	command := NewAddPersonalSpaceCommand(ctx, u.txManager, u.nonTxManager, u.cocotolaCoreCallbackClient)
	if err := command.Execute(ctx, sysAdmin, organizationID, userID); err != nil {
		return mbliberrors.Errorf("command.Execute: %w", err)
	}

	return nil
}
