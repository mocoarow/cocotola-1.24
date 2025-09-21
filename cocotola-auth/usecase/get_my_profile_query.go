package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type GetMyProfileQuery struct {
	nonTxManager mbuserservice.TransactionManager
	logger       *slog.Logger
}

func NewGetMyProfileQuery(nonTxManager mbuserservice.TransactionManager) *GetMyProfileQuery {
	return &GetMyProfileQuery{
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "GetMyProfileQuery")),
	}
}

func (u *GetMyProfileQuery) Execute(ctx context.Context, operator mbuserdomain.UserInterface) (*domain.ProfileModel, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SpaceID, error) {
		spaceManager, err := rf.NewSpaceManager(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewSpaceManager: %w", err)
		}
		privateSpace, err := spaceManager.GetPersonalSpace(ctx, operator)
		if err != nil {
			return nil, mbliberrors.Errorf("GetPersonalSpace: %w", err)
		}
		return privateSpace.SpaceID, nil
	}
	privateSpaceID, err := mblibservice.Do1(ctx, u.nonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return &domain.ProfileModel{
		PrivateSpaceID: privateSpaceID,
	}, nil
}
