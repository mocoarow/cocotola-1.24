package gateway

import (
	"context"

	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type userGroupEntity struct {
	BaseModelEntity
	ID             int
	OrganizationID int
	KeyName        string
	Name           string
	Description    string
	Deleted        bool
}

func (e *userGroupEntity) TableName() string {
	return UserGroupTableName
}

func (e *userGroupEntity) toUserGroup() (*domain.UserGroup, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, liberrors.Errorf("toBaseModel: %w", err)
	}

	userGroupID, err := domain.NewUserGroupID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewUser: %w", err)
	}

	organizationID, err := domain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOrganizationID: %w", err)
	}

	userGroupModel, err := domain.NewUserGroup(baseModel, userGroupID, organizationID, e.KeyName, e.Name, e.Description)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewUserGroup: %w", err)
	}

	return userGroupModel, nil
}

type userGroupRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
}

var _ service.UserGroupRepository = (*userGroupRepository)(nil)

func NewUserGroupRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB) service.UserGroupRepository {
	return &userGroupRepository{
		dialect: dialect,
		db:      db,
	}
}

func (r *userGroupRepository) FindAllUserGroups(ctx context.Context, operator domain.UserInterface) ([]*domain.UserGroup, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.FindAllUserGroups")
	defer span.End()

	userGroups := []userGroupEntity{}
	if result := r.db.Where(&userGroupEntity{ //nolint:exhaustruct
		OrganizationID: operator.GetOrganizationID().Int(),
	}).Find(&userGroups); result.Error != nil {
		return nil, result.Error
	}

	userGroupModels := make([]*domain.UserGroup, len(userGroups))
	for i, e := range userGroups {
		m, err := e.toUserGroup()
		if err != nil {
			return nil, liberrors.Errorf("toUserGroup: %w", err)
		}
		userGroupModels[i] = m
	}

	return userGroupModels, nil
}

func (r *userGroupRepository) FindSystemOwnerGroup(ctx context.Context, _ domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroup, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.FindSystemOwnerGroup")
	defer span.End()

	var userGroup userGroupEntity
	if result := r.db.Where(&userGroupEntity{ //nolint:exhaustruct
		OrganizationID: organizationID.Int(),
		KeyName:        service.SystemOwnerGroupKey,
	}).First(&userGroup); result.Error != nil {
		return nil, result.Error
	}

	return userGroup.toUserGroup()
}

func (r *userGroupRepository) FindUserGroupByID(ctx context.Context, operator domain.UserInterface, userGroupID *domain.UserGroupID) (*domain.UserGroup, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.FindUserGroupByID")
	defer span.End()

	var userGroup userGroupEntity
	if result := r.db.Where("organization_id = ?", operator.GetOrganizationID().Int()).
		Where("id = ? and deleted = ?", userGroupID.Int(), r.dialect.BoolDefaultValue()).
		First(&userGroup); result.Error != nil {
		return nil, result.Error
	}

	return userGroup.toUserGroup()
}

func (r *userGroupRepository) FindUserGroupByKey(ctx context.Context, operator domain.UserInterface, key string) (*domain.UserGroup, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.FindUserGroupByKey")
	defer span.End()

	var userGroup userGroupEntity
	if result := r.db.Where("organization_id = ?", operator.GetOrganizationID().Int()).
		Where("key_name = ? and deleted = ?", key, r.dialect.BoolDefaultValue()).
		First(&userGroup); result.Error != nil {
		return nil, result.Error
	}

	return userGroup.toUserGroup()
}

func (r *userGroupRepository) addUserGroup(userID *domain.UserID, organizationID *domain.OrganizationID, key, name string) (*domain.UserGroupID, error) {
	userGroup := userGroupEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: userID.Int(),
			UpdatedBy: userID.Int(),
		},
		OrganizationID: organizationID.Int(),
		KeyName:        key,
		Name:           name,
	}
	if result := r.db.Create(&userGroup); result.Error != nil {
		return nil, liberrors.Errorf(": %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	userGroupID, err := domain.NewUserGroupID(userGroup.ID)
	if err != nil {
		return nil, liberrors.Errorf("NewUserGroupID: %w", err)
	}

	return userGroupID, nil
}

func (r *userGroupRepository) AddSystemOwnerGroup(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.AddSystemOwnerGroup")
	defer span.End()

	return r.addUserGroup(operator.GetUserID(), organizationID, service.SystemOwnerGroupKey, service.SystemOwnerGroupName)
}

func (r *userGroupRepository) AddOwnerGroup(ctx context.Context, operator domain.SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.AddOwnerGroup")
	defer span.End()

	return r.addUserGroup(operator.GetUserID(), organizationID, service.OwnerGroupKey, service.OwnerGroupName)
}

func (r *userGroupRepository) AddPublicGroup(ctx context.Context, operator domain.SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.AddPublicGroup")
	defer span.End()

	return r.addUserGroup(operator.GetUserID(), organizationID, service.PublicGroupKey, service.PublicGroupName)
}

func (r *userGroupRepository) AddUserGroup(ctx context.Context, operator domain.OwnerInterface, param *service.AddUserGroupParameter) (*domain.UserGroupID, error) {
	_, span := tracer.Start(ctx, "userGroupRepository.AddUserGroup")
	defer span.End()

	userGroup := userGroupEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		KeyName:        param.Key,
		Name:           param.Name,
		Description:    param.Description,
	}
	if result := r.db.Create(&userGroup); result.Error != nil {
		return nil, liberrors.Errorf(": %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	userGroupID, err := domain.NewUserGroupID(userGroup.ID)
	if err != nil {
		return nil, liberrors.Errorf("NewUserGroupID: %w", err)
	}

	return userGroupID, nil
}
