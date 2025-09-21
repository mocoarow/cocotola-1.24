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

type organizationEntity struct {
	BaseModelEntity
	ID   int
	Name string
}

func (e *organizationEntity) TableName() string {
	return OrganizationTableName
}

func (e *organizationEntity) toModel() (*domain.Organization, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, liberrors.Errorf("e.ToBaseModel: %w", err)
	}

	organizationID, err := domain.NewOrganizationID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOrganizationID. err: %w", err)
	}

	organizationModel, err := domain.NewOrganizationModel(baseModel, organizationID, e.Name)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOrganizationModel. err: %w", err)
	}

	return organizationModel, nil
}

type organizationRepository struct {
	db *gorm.DB
}

var _ service.OrganizationRepository = (*organizationRepository)(nil)

func NewOrganizationRepository(_ context.Context, db *gorm.DB) service.OrganizationRepository {
	return &organizationRepository{
		db: db,
	}
}

func (r *organizationRepository) GetOrganization(ctx context.Context, operator domain.UserInterface) (*domain.Organization, error) {
	_, span := tracer.Start(ctx, "organizationRepository.GetOrganization")
	defer span.End()

	var organization organizationEntity
	if result := r.db.Where(organizationEntity{ //nolint:exhaustruct
		ID: operator.GetOrganizationID().Int(),
	}).First(&organization); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrOrganizationNotFound
		}

		return nil, result.Error
	}

	return organization.toModel()
}

func (r *organizationRepository) FindOrganizationByName(ctx context.Context, _ domain.SystemAdminInterface, name string) (*domain.Organization, error) {
	_, span := tracer.Start(ctx, "organizationRepository.FindOrganizationByName")
	defer span.End()

	var organization organizationEntity
	if result := r.db.Where(organizationEntity{ //nolint:exhaustruct
		Name: name,
	}).First(&organization); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrOrganizationNotFound
		}

		return nil, result.Error
	}

	return organization.toModel()
}

func (r *organizationRepository) FindOrganizationByID(ctx context.Context, _ domain.SystemAdminInterface, id *domain.OrganizationID) (*domain.Organization, error) {
	_, span := tracer.Start(ctx, "organizationRepository.FindOrganizationByID")
	defer span.End()

	var organization organizationEntity
	if result := r.db.Where(organizationEntity{ //nolint:exhaustruct
		ID: id.Int(),
	}).First(&organization); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrOrganizationNotFound
		}

		return nil, result.Error
	}

	return organization.toModel()
}

func (r *organizationRepository) AddOrganization(ctx context.Context, operator domain.SystemAdminInterface, organizationName string) (*domain.OrganizationID, error) {
	_, span := tracer.Start(ctx, "organizationRepository.AddOrganization")
	defer span.End()

	organization := organizationEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		Name: organizationName,
	}

	if result := r.db.Create(&organization); result.Error != nil {
		return nil, liberrors.Errorf("db.Create. err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrOrganizationAlreadyExists))
	}

	if organization.ID == 0 {
		return nil, liberrors.Errorf("organization.ID is 0")
	}

	organizationID, err := domain.NewOrganizationID(organization.ID)
	if err != nil {
		return nil, liberrors.Errorf("NewOrganizationID: %w", err)
	}

	return organizationID, nil
}
