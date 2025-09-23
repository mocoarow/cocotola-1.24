package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type ProfileUsecase struct {
	nonTxManager mbuserservice.TransactionManager
	logger       *slog.Logger
}

func NewProfileUsecase(nonTxManager mbuserservice.TransactionManager) *ProfileUsecase {
	return &ProfileUsecase{
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileUsecase")),
	}
}

func (u *ProfileUsecase) GetMyProfile(ctx context.Context, operator mbuserdomain.UserInterface) (*domain.ProfileModel, error) {
	command := NewGetMyProfileQuery(u.nonTxManager)
	profile, err := command.Execute(ctx, operator)
	if err != nil {
		return nil, mbliberrors.Errorf("GetMyProfileQuery.Execute: %w", err)
	}
	return profile, err
}
