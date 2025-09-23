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

type pairOfUserAndGroupEntity struct {
	JunctionModelEntity
	OrganizationID int
	UserID         int
	UserGroupID    int
}

func (u *pairOfUserAndGroupEntity) TableName() string {
	return PairOfUserAndGroupTableName
}

type pairOfUserAndGroupRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

var _ service.PairOfUserAndGroupRepository = (*pairOfUserAndGroupRepository)(nil)

func NewPairOfUserAndGroupRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.PairOfUserAndGroupRepository {
	return &pairOfUserAndGroupRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *pairOfUserAndGroupRepository) AddPairOfUserAndGroupBySystemAdmin(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID, userID *domain.UserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.AddPairOfUserAndGroupToSystemOwner")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndGroupEntity{
		JunctionModelEntity: JunctionModelEntity{ //nolint:exhaustruct
			CreatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: organizationID.Int(),
		UserID:         userID.Int(),
		UserGroupID:    userGroupID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrPairOfUserAndGroupAlreadyExists))
	}

	return nil
}

func (r *pairOfUserAndGroupRepository) AddPairOfUserAndGroup(ctx context.Context, operator domain.UserInterface, userID *domain.UserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.AddPairOfUserAndGroup")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndGroupEntity{
		JunctionModelEntity: JunctionModelEntity{ //nolint:exhaustruct
			CreatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:         userID.Int(),
		UserGroupID:    userGroupID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrPairOfUserAndGroupAlreadyExists))
	}

	return nil
}

func (r *pairOfUserAndGroupRepository) RemovePairOfUserAndGroup(ctx context.Context, operator domain.UserInterface, userID *domain.UserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.RemovePairOfUserAndGroup")
	defer span.End()

	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.GetOrganizationID()}
	db := wrappedDB.
		WherePairOfUserAndGroup().
		Where("user_id = ?", userID.Int()).
		Where("user_group_id = ?", userGroupID.Int()).
		db
	result := db.Delete(&pairOfUserAndGroupEntity{}) //nolint:exhaustruct
	if result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrPairOfUserAndGroupAlreadyExists))
	}
	if result.RowsAffected == 0 {
		return errors.New("ERROR")
	}

	return nil
}

func (r *pairOfUserAndGroupRepository) FindUserGroupsByUserID(ctx context.Context, operator domain.UserInterface, userID *domain.UserID) ([]*domain.UserGroup, error) {
	userGroups := []userGroupEntity{}
	if result := r.db.WithContext(ctx).Table(UserGroupTableName).Select(UserGroupTableName+".*").
		Where(UserGroupTableName+".organization_id = ?", operator.GetOrganizationID().Int()).
		Where(UserGroupTableName+".deleted = ?", r.dialect.BoolDefaultValue()).
		Where(UserTableName+".organization_id = ?", operator.GetOrganizationID().Int()).
		Where(UserTableName+".id = ? and "+UserTableName+".deleted = ?", userID.Int(), r.dialect.BoolDefaultValue()).
		Joins("inner join " + PairOfUserAndGroupTableName + " on " + UserGroupTableName + ".id = " + PairOfUserAndGroupTableName + ".user_group_id").
		Joins("inner join " + UserTableName + " on " + PairOfUserAndGroupTableName + ".user_id = " + UserTableName + ".id").
		Order(UserGroupTableName + ".key_name").
		Find(&userGroups); result.Error != nil {
		return nil, result.Error
	}

	userGroupModels := make([]*domain.UserGroup, len(userGroups))
	for i, e := range userGroups {
		m, err := e.toUserGroup()
		if err != nil {
			return nil, err
		}
		userGroupModels[i] = m
	}

	return userGroupModels, nil
}
