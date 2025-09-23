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

type spaceEntity struct {
	BaseModelEntity
	ID             int
	OrganizationID int
	OwnerID        int
	KeyName        string
	Name           string
	SpaceType      string
	Deleted        bool
}

func (e *spaceEntity) TableName() string {
	return SpaceTableName
}

func (e *spaceEntity) toSpace() (*domain.Space, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, liberrors.Errorf("to base model: %w", err)
	}

	organizationID, err := domain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, liberrors.Errorf("new organization id(%d): %w", e.OrganizationID, err)
	}

	spaceID, err := domain.NewSpaceID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("new space id(%d): %w", e.ID, err)
	}

	ownerID, err := domain.NewUserID(e.OwnerID)
	if err != nil {
		return nil, liberrors.Errorf("new user id(%d): %w", e.OwnerID, err)
	}

	spaceModel, err := domain.NewSpace(
		baseModel,
		spaceID,
		organizationID,
		ownerID,
		e.KeyName,
		e.Name,
		e.SpaceType,
	)
	if err != nil {
		return nil, liberrors.Errorf("new space: %w", err)
	}

	return spaceModel, nil
}

type spaceEntities []spaceEntity

func (e spaceEntities) toSpaces() ([]*domain.Space, error) {
	spaces := make([]*domain.Space, len(e))
	for i, spaceE := range e {
		space, err := spaceE.toSpace()
		if err != nil {
			return nil, liberrors.Errorf("to space: %w", err)
		}
		spaces[i] = space
	}

	return spaces, nil
}

type spaceRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
}

var _ service.SpaceRepository = (*spaceRepository)(nil)

func NewSpaceRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB) service.SpaceRepository {
	return &spaceRepository{
		dialect: dialect,
		db:      db,
	}
}

func (r *spaceRepository) AddSpace(ctx context.Context, operator domain.UserInterface, param *service.AddSpaceParameter) (*domain.SpaceID, error) {
	_, span := tracer.Start(ctx, "spaceRepository.AddSpace")
	defer span.End()

	spaceE := spaceEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		OwnerID:        operator.GetUserID().Int(),
		KeyName:        param.Key,
		Name:           param.Name,
		SpaceType:      param.SpaceType,
	}
	if result := r.db.Create(&spaceE); result.Error != nil {
		return nil, liberrors.Errorf("add space entity: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrSpaceAlreadyExists))
	}

	spaceID, err := domain.NewSpaceID(spaceE.ID)
	if err != nil {
		return nil, liberrors.Errorf("new space id(%d): %w", spaceE.ID, err)
	}

	return spaceID, nil
}

// func (r *spaceRepository) UpdateSpace(ctx context.Context, operator domain.UserInterface, spaceID *domain.SpaceID, version int, param *service.SpaceUpdateParameter) error {
// 	_, span := tracer.Start(ctx, "spaceRepository.UpdateSpace")
// 	defer span.End()

// 	if result := r.db.Model(
// 		&spaceEntity{}, //nolint:exhaustruct
// 	).
// 		Where("organization_id = ?", uint(operator.GetOrganizationID().Int())).
// 		Where("id = ?", spaceID.Int()).
// 		Where("version = ?", version).
// 		Updates(map[string]interface{}{
// 			"version":   gorm.Expr("version + 1"),
// 			"name":      param.Name,
// 			"is_public": param.IsPublic,
// 		}); result.Error != nil {
// 		return liberrors.Errorf("spaceRepository.UpdateSpace: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrSpaceAlreadyExists))
// 	}

// 	return nil
// }

func (r *spaceRepository) FindPublicSpaces(ctx context.Context, operator domain.UserInterface) ([]*domain.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.FindPublicSpaces")
	defer span.End()

	var spacesE spaceEntities
	if result := r.db.WithContext(ctx).Model(
		&spaceEntity{}, //nolint:exhaustruct
	).
		Where("organization_id = ?", uint(operator.GetOrganizationID().Value)).
		Where("space_type = ?", "public").
		Find(&spacesE); result.Error != nil {
		return nil, liberrors.Errorf("spaceRepository.FindPublicSpaces: %w", result.Error)
	}

	spaces, err := spacesE.toSpaces()
	if err != nil {
		return nil, liberrors.Errorf("spacesE.toSpaces: %w", err)
	}
	return spaces, nil
}

func (r *spaceRepository) FindPublicSpaceByKey(ctx context.Context, operator domain.UserInterface, keyName string) (*domain.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.FindPublicSpaceByKey")
	defer span.End()

	var spaceE spaceEntity
	if result := r.db.Model(&spaceE).
		Where("organization_id = ?", uint(operator.GetOrganizationID().Value)).
		Where("key_name = ?", keyName).
		Where("space_type = ?", "public").
		First(&spaceE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrSpaceNotFound
		}

		return nil, liberrors.Errorf("spaceRepository.FindPublicSpaceByKey: %w", result.Error)
	}

	space, err := spaceE.toSpace()
	if err != nil {
		return nil, liberrors.Errorf("spaceE.toSpace: %w", err)
	}

	return space, nil
}

func (r *spaceRepository) GetSpaceByID(ctx context.Context, operator domain.UserInterface, spaceID *domain.SpaceID) (*domain.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.GetSpaceByID")
	defer span.End()

	var spaceE spaceEntity
	if result := r.db.Model(
		&spaceEntity{}, //nolint:exhaustruct
	).
		Where("organization_id = ?", uint(operator.GetOrganizationID().Int())).
		Where("owner_id = ?", uint(operator.GetUserID().Int())).
		Where("id = ?", spaceID.Int()).First(&spaceE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrSpaceNotFound
		}

		return nil, liberrors.Errorf("spaceRepository.GetSpaceByID: %w", result.Error)
	}

	space, err := spaceE.toSpace()
	if err != nil {
		return nil, liberrors.Errorf("spaceE.toSpace: %w", err)
	}

	return space, nil
}
