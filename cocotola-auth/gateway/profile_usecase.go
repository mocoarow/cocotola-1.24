package gateway

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

type ProfileQueryUsecase struct {
	// dialect      mblibgateway.DialectRDBMS
	// driverName   string
	// db           *gorm.DB
	nonTxManager mbuserservice.TransactionManager
	logger       *slog.Logger
}

func NewProfileQueryUsecase(nonTxManager mbuserservice.TransactionManager,

/*dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB*/) *ProfileQueryUsecase {
	return &ProfileQueryUsecase{
		// dialect:      dialect,
		// driverName:   driverName,
		// db:           db,
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "ProfileQueryUsecase")),
	}
}

func (u *ProfileQueryUsecase) GetMyProfile(ctx context.Context, operator mbuserdomain.UserInterface) (*domain.ProfileModel, error) {
	privateSpaceID, err := mblibservice.Do1(ctx, u.nonTxManager, func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SpaceID, error) {
		// mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		// if err != nil {
		// 	return nil, mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		// }
		spaeManager, err := rf.NewSpaceManager(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewSpaceManager: %w", err)
		}
		privateSpace, err := spaeManager.GetPersonalSpace(ctx, operator)
		if err != nil {
			return nil, mbliberrors.Errorf("GetPersonalSpace: %w", err)
		}
		return privateSpace.SpaceID, nil
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return &domain.ProfileModel{
		PrivateSpaceID: privateSpaceID,
	}, nil
}
