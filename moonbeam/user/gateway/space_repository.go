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
	Key            string
	Name           string
	IsPublic       bool
}

func (e *spaceEntity) TableName() string {
	return SpaceTableName
}

func (e *spaceEntity) ToModel() (*domain.SpaceModel, error) {
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

	ownerID, err := domain.NewAppUserID(e.OwnerID)
	if err != nil {
		return nil, liberrors.Errorf("new app user id(%d): %w", e.OwnerID, err)
	}

	spaceModel, err := domain.NewSpaceModel(
		baseModel,
		spaceID,
		organizationID,
		ownerID,
		e.Key,
		e.Name,
	)
	if err != nil {
		return nil, liberrors.Errorf("new space model: %w", err)
	}

	return spaceModel, nil
}

func (e *spaceEntity) toSpace() (*service.Space, error) {
	spaceModel, err := e.ToModel()
	if err != nil {
		return nil, liberrors.Errorf("to space model: %w", err)
	}

	space, err := service.NewSpace(spaceModel)
	if err != nil {
		return nil, liberrors.Errorf("new space: %w", err)
	}

	return space, nil
}

type spaceEntities []spaceEntity

func (e spaceEntities) toSpaces() ([]*service.Space, error) {
	spaces := make([]*service.Space, len(e))
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

func (r *spaceRepository) AddSpace(ctx context.Context, operator service.OperatorInterface, param *service.SpaceAddParameter) (*domain.SpaceID, error) {
	_, span := tracer.Start(ctx, "spaceRepository.AddSpace")
	defer span.End()

	spaceE := spaceEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		OwnerID:        operator.AppUserID().Int(),
		Key:            param.Key,
		Name:           param.Name,
		IsPublic:       param.IsPublic,
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

func (r *spaceRepository) UpdateSpace(ctx context.Context, operator service.OperatorInterface, spaceID *domain.SpaceID, version int, param *service.SpaceUpdateParameter) error {
	_, span := tracer.Start(ctx, "spaceRepository.UpdateSpace")
	defer span.End()

	if result := r.db.Model(
		&spaceEntity{}, //nolint:exhaustruct
	).
		Where("organization_id = ?", uint(operator.OrganizationID().Int())).
		Where("id = ?", spaceID.Int()).
		Where("version = ?", version).
		Updates(map[string]interface{}{
			"version":   gorm.Expr("version + 1"),
			"name":      param.Name,
			"is_public": param.IsPublic,
		}); result.Error != nil {
		return liberrors.Errorf("spaceRepository.UpdateSpace: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrSpaceAlreadyExists))
	}

	return nil
}

func (r *spaceRepository) FindPublicSpaces(ctx context.Context, operator service.OperatorInterface) ([]*service.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.FindPublicSpaces")
	defer span.End()

	var spacesE spaceEntities
	if result := r.db.WithContext(ctx).Model(
		&spaceEntity{}, //nolint:exhaustruct
	).
		Where("organization_id = ?", uint(operator.OrganizationID().Value)).
		Where("is_public = ?", true).
		Find(&spacesE); result.Error != nil {
		return nil, liberrors.Errorf("spaceRepository.FindPublicSpaces: %w", result.Error)
	}

	spaces, err := spacesE.toSpaces()
	if err != nil {
		return nil, liberrors.Errorf("spacesE.toSpaces: %w", err)
	}
	return spaces, nil
}

func (r *spaceRepository) FindPublicSpaceByKey(ctx context.Context, key string) (*service.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.FindPublicSpaceByKey")
	defer span.End()

	var spaceE spaceEntity
	if result := r.db.Model(&spaceE).
		Where("key = ?", key).
		Where("is_public = ?", true).
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

func (r *spaceRepository) GetSpaceByID(ctx context.Context, operator service.OperatorInterface, spaceID *domain.SpaceID) (*service.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.GetSpaceByID")
	defer span.End()

	var spaceE spaceEntity
	if result := r.db.Model(
		&spaceEntity{}, //nolint:exhaustruct
	).
		Where("organization_id = ?", uint(operator.OrganizationID().Int())).
		Where("owner_id = ?", uint(operator.AppUserID().Int())).
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
