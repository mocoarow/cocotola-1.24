package usecase

import (
	"context"
	"log/slog"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type VerifyPasswordCommand struct {
	nonTxManager service.TransactionManager
	logger       *slog.Logger
}

func NewVerifyPasswordCommand(nonTxManager service.TransactionManager) *VerifyPasswordCommand {
	return &VerifyPasswordCommand{
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(liblog.LoggerNameKey, "AddOrganizationCommand")),
	}
}

func (u *VerifyPasswordCommand) Execute(ctx context.Context, operator service.SystemOwnerInterface, loginID, password string) error {
	fn := func(rf service.RepositoryFactory) (bool, error) {
		userRepo := rf.NewUserRepository(ctx)
		ok, err := userRepo.VerifyPassword(ctx, operator, loginID, password)
		if err != nil {
			return false, liberrors.Errorf("m.userRepo.VerifyPassword. err: %w", err)
		}
		return ok, nil
	}
	ok, err := libservice.Do1(ctx, u.nonTxManager, fn)
	if err != nil {
		return err //nolint:wrapcheck
	}
	if !ok {
		return service.ErrUnauthenticated
	}

	return nil
}
