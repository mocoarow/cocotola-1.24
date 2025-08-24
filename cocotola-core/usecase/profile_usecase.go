package usecase

import (
	"context"
	"log/slog"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

type ProfileUsecase struct {
	nonTxManager service.TransactionManager
	logger       *slog.Logger
}

func NewProfileUsecase(nonTxManager service.TransactionManager) *ProfileUsecase {
	return &ProfileUsecase{
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileUsecase")),
	}
}

func (u *ProfileUsecase) GetMyProfile(ctx context.Context, operator service.OperatorInterface) (*domain.ProfileModel, error) {
	privateSpaceID, err := mblibservice.Do1(ctx, u.nonTxManager, func(rf service.RepositoryFactory) (*domain.SpaceID, error) {
		pairofUserAndSpaceRepo, err := rf.NewPairOfUserAndSpaceRepository(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewPairOfUserAndSpaceRepository: %w", err)
		}
		spaces, err := pairofUserAndSpaceRepo.FindSpacesByUserID(ctx, operator, operator.AppUserID())
		if err != nil {
			return nil, mbliberrors.Errorf("FindSpacesByUserID: %w", err)
		}

		for _, space := range spaces {
			u.logger.InfoContext(ctx, "GetMyProfile: space", slog.Int("space_id", space.SpaceID.Int()), slog.String("space_key", space.Key))
		}

		for _, space := range spaces {
			if space.IsPrivate() {
				return space.SpaceID, nil
			}
		}

		return nil, service.ErrSpaceNotFound
	})
	if err != nil {
		return nil, mbliberrors.Errorf("add deck: %w", err)
	}

	return &domain.ProfileModel{
		PrivateSpaceID: privateSpaceID,
	}, nil
}
