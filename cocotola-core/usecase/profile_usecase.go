package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type ProfileUsecase struct {
	nonTxManager       service.TransactionManager
	cocotolaAuthClient libapi.CocotolaAuthClient
	logger             *slog.Logger
}

func NewProfileUsecase(nonTxManager service.TransactionManager, cocotolaAuthClient libapi.CocotolaAuthClient) *ProfileUsecase {
	return &ProfileUsecase{
		nonTxManager:       nonTxManager,
		cocotolaAuthClient: cocotolaAuthClient,
		logger:             slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileUsecase")),
	}
}

func (u *ProfileUsecase) GetMyProfile(ctx context.Context, operator mbuserdomain.UserInterface, bearerToken string) (*domain.ProfileModel, error) {
	command := NewGetMyProfileQuery(u.nonTxManager, u.cocotolaAuthClient)
	profile, err := command.Execute(ctx, operator, bearerToken)
	if err != nil {
		return nil, mbliberrors.Errorf("GetMyProfileQuery.Execute: %w", err)
	}
	return profile, err
}
