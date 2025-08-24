package gateway

import (
	"context"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type SpaceEntity struct {
	mbusergateway.BaseModelEntity
	ID             int
	OrganizationID int
	OwnerID        int
	Key            string
	Name           string
}

func (e *SpaceEntity) TableName() string {
	return "core_space"
}

func (e *SpaceEntity) ToModel() (*domain.SpaceModel, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to base model: %w", err)
	}

	organizationID, err := mbuserdomain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("new organization id(%d): %w", e.OrganizationID, err)
	}

	spaceID, err := domain.NewSpaceID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(%d): %w", e.ID, err)
	}

	ownerID, err := mbuserdomain.NewAppUserID(e.OwnerID)
	if err != nil {
		return nil, mbliberrors.Errorf("new app user id(%d): %w", e.OwnerID, err)
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
		return nil, mbliberrors.Errorf("new space model: %w", err)
	}

	return spaceModel, nil
}

func (e *SpaceEntity) toSpace() (*service.Space, error) {
	spaceModel, err := e.ToModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to space model: %w", err)
	}
	space := &service.Space{SpaceModel: spaceModel}
	return space, nil
}

type spaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) service.SpaceRepository {
	return &spaceRepository{
		db: db,
	}
}

func (r *spaceRepository) AddSpace(ctx context.Context, operator mbuserservice.OperatorInterface, param *service.SpaceAddParameter) (*domain.SpaceID, error) {
	_, span := tracer.Start(ctx, "spaceRepository.AddSpace")
	defer span.End()

	spaceE := SpaceEntity{
		BaseModelEntity: mbusergateway.BaseModelEntity{
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		OwnerID:        operator.AppUserID().Int(),
		Key:            param.Key,
		Name:           param.Name,
	}
	if result := r.db.Create(&spaceE); result.Error != nil {
		return nil, mbliberrors.Errorf("add space entity: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrSpaceAlreadyExists))
	}

	spaceID, err := domain.NewSpaceID(spaceE.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(%d): %w", spaceE.ID, err)
	}

	return spaceID, nil
}

func (r *spaceRepository) UpdateSpace(ctx context.Context, operator service.OperatorInterface, spaceID *domain.SpaceID, version int, param *service.SpaceUpdateParameter) error {
	_, span := tracer.Start(ctx, "spaceRepository.UpdateSpace")
	defer span.End()

	if result := r.db.Model(&SpaceEntity{}).
		Where("organization_id = ?", uint(operator.OrganizationID().Int())).
		Where("id = ?", spaceID.Int()).
		Where("version = ?", version).
		Updates(map[string]interface{}{
			"version": gorm.Expr("version + 1"),
			"name":    param.Name,
		}); result.Error != nil {
		return mbliberrors.Errorf("spaceRepository.UpdateSpace: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrSpaceAlreadyExists))
	}

	return nil
}

func (r *spaceRepository) FindSpaces(ctx context.Context, operator service.OperatorInterface) ([]*service.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.FindSpaces")
	defer span.End()

	var spacesE []SpaceEntity
	if result := r.db.Model(&SpaceEntity{}).
		Where("organization_id = ?", uint(operator.OrganizationID().Value)).
		Where("owner_id = ?", uint(operator.AppUserID().Value)).
		Find(&spacesE); result.Error != nil {
		return nil, mbliberrors.Errorf("spaceRepository.FindSpaces: %w", result.Error)
	}

	var spaces []*service.Space
	for _, spaceE := range spacesE {
		space, err := spaceE.toSpace()
		if err != nil {
			return nil, err
		}
		spaces = append(spaces, space)
	}

	return spaces, nil
}

func (r *spaceRepository) GetSpaceByID(ctx context.Context, operator service.OperatorInterface, spaceID *domain.SpaceID) (*service.Space, error) {
	_, span := tracer.Start(ctx, "spaceRepository.GetSpaceByID")
	defer span.End()

	var spaceE SpaceEntity
	if result := r.db.Model(&SpaceEntity{}).
		Where("organization_id = ?", uint(operator.OrganizationID().Int())).
		Where("owner_id = ?", uint(operator.AppUserID().Int())).
		Where("id = ?", spaceID.Int()).First(&spaceE); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, service.ErrSpaceNotFound
		}
		return nil, mbliberrors.Errorf("spaceRepository.GetSpaceByID: %w", result.Error)
	}

	space, err := spaceE.toSpace()
	if err != nil {
		return nil, mbliberrors.Errorf("spaceE.toSpace: %w", err)
	}
	return space, nil
}
