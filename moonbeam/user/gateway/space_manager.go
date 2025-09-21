package gateway

import (
	"context"
	"errors"

	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type spaceManager struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

func NewSpaceManager(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) (service.SpaceManager, error) {
	return &spaceManager{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}, nil
}

func (m *spaceManager) AddPublicDefaultSpace(ctx context.Context, operator domain.UserInterface) (*domain.SpaceID, error) {
	userRepo := m.rf.NewUserRepository(ctx)

	user, err := userRepo.FindUserByID(ctx, operator, operator.GetUserID())
	if err != nil {
		return nil, liberrors.Errorf("FindUserByID: %w", err)
	}

	spaceRepo := m.rf.NewSpaceRepository(ctx)
	addSpaceParam := service.AddSpaceParameter{
		Key:       service.PublicDefaultSpaceKey,
		Name:      service.PublicDefaultSpaceName,
		SpaceType: "public",
	}

	spaceID, err := spaceRepo.AddSpace(ctx, operator, &addSpaceParam)
	if err != nil {
		return nil, liberrors.Errorf("AddSpace: %w", err)
	}

	pairOfUserAndSpaceRepo := NewPairOfUserAndSpaceRepository(ctx, m.dialect, m.db)
	if err := pairOfUserAndSpaceRepo.AddPairOfUserAndSpace(ctx, user, user.GetUserID(), spaceID); err != nil {
		return nil, liberrors.Errorf("AddPairOfUserAndSpace: %w", err)
	}
	return spaceID, nil
}

func (m *spaceManager) AddPersonalSpace(ctx context.Context, operator domain.UserInterface, param *service.AddPersonalSpaceParameter) (*domain.SpaceID, error) {
	userRepo := m.rf.NewUserRepository(ctx)

	user, err := userRepo.FindUserByID(ctx, operator, param.UserID)
	if err != nil {
		return nil, liberrors.Errorf("FindUserByID: %w", err)
	}

	spaceRepo := m.rf.NewSpaceRepository(ctx)
	addSpaceParam := service.AddSpaceParameter{
		Key:       param.KeyName,
		Name:      param.Name,
		SpaceType: "personal",
	}

	spaceID, err := spaceRepo.AddSpace(ctx, operator, &addSpaceParam)
	if err != nil {
		return nil, liberrors.Errorf("AddSpace: %w", err)
	}

	pairOfUserAndSpaceRepo := NewPairOfUserAndSpaceRepository(ctx, m.dialect, m.db)
	if err := pairOfUserAndSpaceRepo.AddPairOfUserAndSpace(ctx, user, user.GetUserID(), spaceID); err != nil {
		return nil, liberrors.Errorf("AddPairOfUserAndSpace: %w", err)
	}
	return spaceID, nil
}

func (m *spaceManager) AddUserToSpace(_ context.Context, _ domain.SystemOwnerInterface, _ domain.UserID, _ *domain.SpaceID) error {
	return errors.New("not implemented")
}

func (m *spaceManager) GetPersonalSpace(ctx context.Context, operator domain.UserInterface) (*domain.SpaceModel, error) {
	pairOfUserAndSpaceRepo := NewPairOfUserAndSpaceRepository(ctx, m.dialect, m.db)
	spaces, err := pairOfUserAndSpaceRepo.FindMySpaces(ctx, operator)
	if err != nil {
		return nil, liberrors.Errorf("FindMySpaces: %w", err)
	}
	for _, space := range spaces {
		if space.SpaceType == "personal" {
			return space, nil
		}
	}

	return nil, liberrors.Errorf("personal space not found: %w", service.ErrSpaceNotFound)
}
