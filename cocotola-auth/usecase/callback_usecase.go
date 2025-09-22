package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type CallbackUsecase struct {
	systemToken                libdomain.SystemToken
	mbTxManager                mbuserservice.TransactionManager
	mbNonTxManager             mbuserservice.TransactionManager
	cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient
	logger                     *slog.Logger
}

func NewCallbackUsecase(systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient) *CallbackUsecase {
	return &CallbackUsecase{
		systemToken:                systemToken,
		mbTxManager:                mbTxManager,
		mbNonTxManager:             mbNonTxManager,
		cocotolaCoreCallbackClient: cocotolaCoreCallbackClient,
		logger:                     slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

func (u *CallbackUsecase) OnAddUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error {
	ctx, span := tracer.Start(ctx, "CallbackUsecase.OnAddUser")
	defer span.End()

	u.logger.InfoContext(ctx, "OnAddUser", slog.Int("user_id", userID.Int()))

	sysAdmin := service.NewSystemAdmin(u.systemToken)
	sysOwner, err := u.findSystemOwnerByOrganizationID(ctx, sysAdmin, organizationID)
	if err != nil {
		return mbliberrors.Errorf("findSystemOwnerByOrganizationID: %w", err)
	}
	command := NewCallbackOnAddUserCommand(ctx, u.mbTxManager, u.mbNonTxManager, u.cocotolaCoreCallbackClient)
	if err := command.Execute(ctx, sysOwner, organizationID, userID); err != nil {
		return mbliberrors.Errorf("command.Execute: %w", err)
	}

	return nil
}

func (u *CallbackUsecase) findSystemOwnerByOrganizationID(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationID *mbuserdomain.OrganizationID) (*mbuserdomain.SystemOwner, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) { //nolint:wrapcheck
		return service.FindSystemOwnerByOrganizationID(ctx, mbrf, operator, organizationID)
	})
}
