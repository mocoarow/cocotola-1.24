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

type organizationRepository struct {
	db *gorm.DB
}

type organizationEntity struct {
	BaseModelEntity
	ID   int
	Name string
}

func (e *organizationEntity) TableName() string {
	return OrganizationTableName
}

func (e *organizationEntity) toModel() (*service.Organization, error) {
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

	org, err := service.NewOrganization(organizationModel)
	if err != nil {
		return nil, liberrors.Errorf("service.NewOrganization. err: %w", err)
	}

	return org, nil
}

func NewOrganizationRepository(ctx context.Context, db *gorm.DB) service.OrganizationRepository {
	return &organizationRepository{
		db: db,
	}
}

func (r *organizationRepository) GetOrganization(ctx context.Context, operator service.AppUserInterface) (*service.Organization, error) {
	_, span := tracer.Start(ctx, "organizationRepository.GetOrganization")
	defer span.End()

	organization := organizationEntity{}

	if result := r.db.Where(organizationEntity{
		ID: operator.OrganizationID().Int(),
	}).First(&organization); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrOrganizationNotFound
		}
		return nil, result.Error
	}

	return organization.toModel()
}

func (r *organizationRepository) FindOrganizationByName(ctx context.Context, operator service.SystemAdminInterface, name string) (*service.Organization, error) {
	_, span := tracer.Start(ctx, "organizationRepository.FindOrganizationByName")
	defer span.End()

	organization := organizationEntity{}

	if result := r.db.Where(organizationEntity{
		Name: name,
	}).First(&organization); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrOrganizationNotFound
		}
		return nil, result.Error
	}

	return organization.toModel()
}

func (r *organizationRepository) FindOrganizationByID(ctx context.Context, operator service.SystemAdminInterface, id *domain.OrganizationID) (*service.Organization, error) {
	_, span := tracer.Start(ctx, "organizationRepository.FindOrganizationByID")
	defer span.End()

	organization := organizationEntity{}

	if result := r.db.Where(organizationEntity{
		ID: id.Int(),
	}).First(&organization); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrOrganizationNotFound
		}
		return nil, result.Error
	}

	return organization.toModel()
}

func (r *organizationRepository) AddOrganization(ctx context.Context, operator service.SystemAdminInterface, param service.OrganizationAddParameterInterface) (*domain.OrganizationID, error) {
	_, span := tracer.Start(ctx, "organizationRepository.AddOrganization")
	defer span.End()

	organization := organizationEntity{
		BaseModelEntity: BaseModelEntity{
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		Name: param.Name(),
	}

	if result := r.db.Create(&organization); result.Error != nil {
		return nil, liberrors.Errorf("db.Create. err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrOrganizationAlreadyExists))
	}

	if organization.ID == 0 {
		return nil, liberrors.Errorf("organization.ID is 0")
	}

	organizationID, err := domain.NewOrganizationID(organization.ID)
	if err != nil {
		return nil, err
	}

	return organizationID, nil
}
