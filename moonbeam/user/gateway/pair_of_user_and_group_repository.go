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

var (
	PairOfUserAndGroupTableName = "user_n_group"
)

type pairOfUserAndGroupRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

type pairOfUserAndGroupEntity struct {
	JunctionModelEntity
	OrganizationID int
	AppUserID      int
	UserGroupID    int
}

func (u *pairOfUserAndGroupEntity) TableName() string {
	return PairOfUserAndGroupTableName
}

func NewPairOfUserAndGroupRepository(ctx context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.PairOfUserAndGroupRepository {
	return &pairOfUserAndGroupRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *pairOfUserAndGroupRepository) AddPairOfUserAndGroupBySystemAdmin(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.AddPairOfUserAndGroupToSystemOwner")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndGroupEntity{
		JunctionModelEntity: JunctionModelEntity{
			CreatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: organizationID.Int(),
		AppUserID:      appUserID.Int(),
		UserGroupID:    userGroupID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
	}

	// if err := r.add(ctx, operator.GetAppUserID(), systemOwner.GetOrganizationID(), systemOwner.GetAppUserID(), userGroupID,
	// /* service.SystemOwnerGroupKey*/
	// ); err != nil {
	// 	return nil
	// }

	return nil
}

func (r *pairOfUserAndGroupRepository) AddPairOfUserAndGroup(ctx context.Context, operator service.AppUserInterface, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.AddPairOfUserAndGroup")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndGroupEntity{
		JunctionModelEntity: JunctionModelEntity{
			CreatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		AppUserID:      appUserID.Int(),
		UserGroupID:    userGroupID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
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

	// if err := r.add(ctx, operator.GetAppUserID(), operator.GetOrganizationID(), appUserID, userGroupID,
	// /*userGroup.GetKey()*/
	// ); err != nil {
	// 	return err
	// }
	return nil
}

// func (r *pairOfUserAndGroupRepository) add(ctx context.Context, operatorID domain.AppUserID, organizationID domain.OrganizationID, appUserID domain.AppUserID, userGroupID domain.UserGroupID,
// 	/*userGroupKey string*/
// 	) error {
// 	// add pairOfOuserAndRole
// 	pairOfUserAndGroup := pairOfUserAndGroupEntity{
// 		JunctionModelEntity: JunctionModelEntity{
// 			CreatedBy: operatorID.Int(),
// 		},
// 		OrganizationID: organizationID.Int(),
// 		AppUserID:      appUserID.Int(),
// 		UserGroupID:    userGroupID.Int(),
// 	}
// 	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
// 		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
// 	}

// 	rbacRepo := r.rf.NewRBACRepository(ctx)
// 	rbacAppUser := service.NewRBACAppUser(organizationID, appUserID)
// 	rbacUserRole := service.NewRBACUserRole(organizationID, userGroupID)
// 	rbacDomain := service.NewRBACOrganization(organizationID)

// 	// app-user belongs to user-role
// 	if err := rbacRepo.AddSubjectGroupingPolicy(rbacDomain, rbacAppUser, rbacUserRole); err != nil {
// 		return liberrors.Errorf("rbacRepo.AddNamedGroupingPolicy. err: %w", err)
// 	}

// 	return nil
// }

func (r *pairOfUserAndGroupRepository) RemovePairOfUserAndGroup(ctx context.Context, operator service.AppUserInterface, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndGroupRepository.RemovePairOfUserAndGroup")
	defer span.End()

	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.OrganizationID()}
	db := wrappedDB.
		WherePairOfUserAndGroup().
		Where("app_user_id = ?", appUserID.Int()).
		Where("user_group_id = ?", userGroupID.Int()).
		db
	result := db.Delete(&pairOfUserAndGroupEntity{})
	if result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
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

	// if err := r.remove(ctx, operator.GetAppUserID(), operator.GetOrganizationID(), appUserID, userGroupID,
	// /*userGroup.GetKey()*/); err != nil {
	// 	return err
	// }
	return nil
}

// func (r *pairOfUserAndGroupRepository) remove(ctx context.Context, operatorID domain.AppUserID, organizationID domain.OrganizationID, appUserID domain.AppUserID, userGroupID domain.UserGroupID,

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
// 		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("ERROR")
// 	}

// 	rbacRepo := r.rf.NewRBACRepository(ctx)
// 	rbacAppUser := service.NewRBACAppUser(organizationID, appUserID)
// 	rbacUserRole := service.NewRBACUserRole(organizationID, userGroupID)
// 	rbacDomain := service.NewRBACOrganization(organizationID)

// 	// remove relationship
// 	if err := rbacRepo.RemoveSubjectGroupingPolicy(rbacDomain, rbacAppUser, rbacUserRole); err != nil {
// 		return liberrors.Errorf("rbacRepo.RemoveSubjectGroupingPolicy. err: %w", err)
// 	}

// 	return nil
// }

// func (r *pairOfUserAndGroupRepository) enforce(ctx context.Context, operator domain.AppUserModel, rbacObject domain.RBACObject, rbacAction domain.RBACAction) (bool, error) {
// 	rbacDomain := service.NewRBACOrganization(operator.GetOrganizationID())

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
// 	rbacOperator := service.NewRBACAppUser(operator.GetOrganizationID(), operator.GetAppUserID())
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

func (r *pairOfUserAndGroupRepository) FindUserGroupsByUserID(ctx context.Context, operator service.AppUserInterface, appUserID *domain.AppUserID) ([]*domain.UserGroupModel, error) {
	userGroups := []userGroupEntity{}
	if result := r.db.Table("user_group").Select("user_group.*").
		Where("user_group.organization_id = ?", operator.OrganizationID().Int()).
		Where("user_group.removed = ?", r.dialect.BoolDefaultValue()).
		Where("app_user.organization_id = ?", operator.OrganizationID().Int()).
		Where("app_user.id = ? and app_user.removed = ?", appUserID.Int(), r.dialect.BoolDefaultValue()).
		Joins("inner join user_n_group on user_group.id = user_n_group.user_group_id").
		Joins("inner join app_user on user_n_group.app_user_id = app_user.id").
		Order("user_group.key_name").
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
