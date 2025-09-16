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

type pairOfUserAndGroupRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

type pairOfUserAndGroupEntity struct {
	JunctionModelEntity
	OrganizationID int
	UserID      int
	UserGroupID    int
}

func (u *pairOfUserAndGroupEntity) TableName() string {
	return PairOfUserAndGroupTableName
}

func NewPairOfUserAndGroupRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.PairOfUserAndGroupRepository {
	return &pairOfUserAndGroupRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *pairOfUserAndGroupRepository) AddPairOfUserAndGroupBySystemAdmin(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.AddPairOfUserAndGroupToSystemOwner")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndGroupEntity{
		JunctionModelEntity: JunctionModelEntity{ //nolint:exhaustruct
			CreatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: organizationID.Int(),
		UserID:      appUserID.Int(),
		UserGroupID:    userGroupID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	// if err := r.add(ctx, operator.GetUserID(), systemOwner.GetOrganizationID(), systemOwner.GetUserID(), userGroupID,
	// /* service.SystemOwnerGroupKey*/
	// ); err != nil {
	// 	return nil
	// }

	return nil
}

func (r *pairOfUserAndGroupRepository) AddPairOfUserAndGroup(ctx context.Context, operator service.UserInterface, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.AddPairOfUserAndGroup")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndGroupEntity{
		JunctionModelEntity: JunctionModelEntity{ //nolint:exhaustruct
			CreatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:      appUserID.Int(),
		UserGroupID:    userGroupID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	// rbacUserRoleObject := service.NewRBACUserRoleObject(operator.GetOrganizationID(), userGroupID)

	// ok, err := r.enforce(ctx, operator, rbacUserRoleObject, service.RBACSetAction)
	// if err != nil {
	// 	return err
	// }
	// if !ok {
	// 	return libdomain.ErrPermissionDenied
	// }

	// // userGroupRepo := r.rf.NewUserGroupRepository(ctx)
	// // userGroup, err := userGroupRepo.FindUserGroupByID(ctx, operator, userGroupID)
	// // if err != nil {
	// // 	return err
	// // }

	// if err := r.add(ctx, operator.GetUserID(), operator.GetOrganizationID(), appUserID, userGroupID,
	// /*userGroup.GetKey()*/
	// ); err != nil {
	// 	return err
	// }
	return nil
}

// func (r *pairOfUserAndGroupRepository) add(ctx context.Context, operatorID domain.UserID, organizationID domain.OrganizationID, appUserID domain.UserID, userGroupID domain.UserGroupID,
// 	/*userGroupKey string*/
// 	) error {
// 	// add pairOfOuserAndRole
// 	pairOfUserAndGroup := pairOfUserAndGroupEntity{
// 		JunctionModelEntity: JunctionModelEntity{
// 			CreatedBy: operatorID.Int(),
// 		},
// 		OrganizationID: organizationID.Int(),
// 		UserID:      appUserID.Int(),
// 		UserGroupID:    userGroupID.Int(),
// 	}
// 	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
// 		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
// 	}

// 	rbacRepo := r.rf.NewRBACRepository(ctx)
// 	rbacUser := service.NewRBACUser(organizationID, appUserID)
// 	rbacUserRole := service.NewRBACUserRole(organizationID, userGroupID)
// 	rbacDomain := service.NewRBACDomainFromOrganization(organizationID)

// 	// app-user belongs to user-role
// 	if err := rbacRepo.AddSubjectGroupingPolicy(rbacDomain, rbacUser, rbacUserRole); err != nil {
// 		return liberrors.Errorf("rbacRepo.AddNamedGroupingPolicy. err: %w", err)
// 	}

// 	return nil
// }

func (r *pairOfUserAndGroupRepository) RemovePairOfUserAndGroup(ctx context.Context, operator service.UserInterface, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.RemovePairOfUserAndGroup")
	defer span.End()

	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.GetOrganizationID()}
	db := wrappedDB.
		WherePairOfUserAndGroup().
		Where("app_user_id = ?", appUserID.Int()).
		Where("user_group_id = ?", userGroupID.Int()).
		db
	result := db.Delete(&pairOfUserAndGroupEntity{}) //nolint:exhaustruct
	if result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}
	if result.RowsAffected == 0 {
		return errors.New("ERROR")
	}

	// rbacUserRoleObject := service.NewRBACUserRoleObject(operator.GetOrganizationID(), userGroupID)

	// ok, err := r.enforce(ctx, operator, rbacUserRoleObject, service.RBACUnsetAction)
	// if err != nil {
	// 	return err
	// }
	// if !ok {
	// 	return libdomain.ErrPermissionDenied
	// }

	// // userGroupRepo := r.rf.NewUserGroupRepository(ctx)
	// // userGroup, err := userGroupRepo.FindUserGroupByID(ctx, operator, userGroupID)
	// // if err != nil {
	// // 	return err
	// // }

	// if err := r.remove(ctx, operator.GetUserID(), operator.GetOrganizationID(), appUserID, userGroupID,
	// /*userGroup.GetKey()*/); err != nil {
	// 	return err
	// }
	return nil
}

// func (r *pairOfUserAndGroupRepository) remove(ctx context.Context, operatorID domain.UserID, organizationID domain.OrganizationID, appUserID domain.UserID, userGroupID domain.UserGroupID,

// /* userGroupKey string*/
// ) error {
// 	// remove pairOfOuserAndRole
// 	wrappedDB := wrappedDB{dialect: r.dialect,db: r.db, organizationID: organizationID}
// 	db := wrappedDB.
// 		WherePairOfUserAndGroup().
// 		Where("`app_user_id` = ?", appUserID.Int()).
// 		Where("`user_group_id` = ?", userGroupID.Int()).
// 		db
// 	result := db.Delete(&pairOfUserAndGroupEntity{})
// 	if result.Error != nil {
// 		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("ERROR")
// 	}

// 	rbacRepo := r.rf.NewRBACRepository(ctx)
// 	rbacUser := service.NewRBACUser(organizationID, appUserID)
// 	rbacUserRole := service.NewRBACUserRole(organizationID, userGroupID)
// 	rbacDomain := service.NewRBACDomainFromOrganization(organizationID)

// 	// remove relationship
// 	if err := rbacRepo.RemoveSubjectGroupingPolicy(rbacDomain, rbacUser, rbacUserRole); err != nil {
// 		return liberrors.Errorf("rbacRepo.RemoveSubjectGroupingPolicy. err: %w", err)
// 	}

// 	return nil
// }

// func (r *pairOfUserAndGroupRepository) enforce(ctx context.Context, operator domain.UserModel, rbacObject domain.RBACObject, rbacAction domain.RBACAction) (bool, error) {
// 	rbacDomain := service.NewRBACDomainFromOrganization(operator.GetOrganizationID())

// 	userGroupRepo:= r.rf.NewUserGroupRepository(ctx)
// 	userGroups,err:= userGroupRepo.FindAllUserGroups(ctx, operator)
// 	if err!=nil{
// 		return false, err
// 	}

// 	rbacRoles := make([]domain.RBACRole, 0)
// 	for _, userGroup := range userGroups{
// 		rbacRoles = append(rbacRoles, service.NewRBACUserRole(operator.GetOrganizationID(), userGroup.GetUerGroupID()))
// 	}

// 	rbacRepo := r.rf.NewRBACRepository(ctx)
// 	rbacOperator := service.NewRBACUser(operator.GetOrganizationID(), operator.GetUserID())
// 	e, err := rbacRepo.NewEnforcerWithGroupsAndUsers(rbacRoles, []domain.RBACUser{rbacOperator})
// 	if err != nil {
// 		return false, err
// 	}

// 	ok, err := e.Enforce(rbacOperator.Subject(), rbacObject.Object(), rbacAction.Action(), rbacDomain.Domain(), )
// 	if err != nil {
// 		return false, err
// 	}
// 	if ok {
// 		return true, nil
// 	}

// 	return false, nil
// }

func (r *pairOfUserAndGroupRepository) FindUserGroupsByUserID(ctx context.Context, operator service.UserInterface, appUserID *domain.UserID) ([]*domain.UserGroupModel, error) {
	userGroups := []userGroupEntity{}
	if result := r.db.WithContext(ctx).Table(UserGroupTableName).Select(UserGroupTableName+".*").
		Where(UserGroupTableName+".organization_id = ?", operator.GetOrganizationID().Int()).
		Where(UserGroupTableName+".deleted = ?", r.dialect.BoolDefaultValue()).
		Where(UserTableName+".organization_id = ?", operator.GetOrganizationID().Int()).
		Where(UserTableName+".id = ? and "+UserTableName+".deleted = ?", appUserID.Int(), r.dialect.BoolDefaultValue()).
		Joins("inner join " + PairOfUserAndGroupTableName + " on " + UserGroupTableName + ".id = " + PairOfUserAndGroupTableName + ".user_group_id").
		Joins("inner join " + UserTableName + " on " + PairOfUserAndGroupTableName + ".app_user_id = " + UserTableName + ".id").
		Order(UserGroupTableName + ".key_name").
		Find(&userGroups); result.Error != nil {
		return nil, result.Error
	}

	userGroupModels := make([]*domain.UserGroupModel, len(userGroups))
	for i, e := range userGroups {
		m, err := e.toUserGroupModel()
		if err != nil {
			return nil, err
		}
		userGroupModels[i] = m
	}

	return userGroupModels, nil
}
