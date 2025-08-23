package service

import (
	"context"
	"fmt"
	"log/slog"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var _ SystemAdminInterface = (*SystemAdmin)(nil)

type SystemAdminInterface interface {
	AppUserID() *domain.AppUserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type SystemAdmin struct {
	*domain.SystemAdminModel
	rf          RepositoryFactory
	orgRepo     OrganizationRepository
	appUserRepo AppUserRepository
	logger      *slog.Logger
}

func NewSystemAdmin(ctx context.Context, rf RepositoryFactory) (*SystemAdmin, error) {
	if rf == nil {
		return nil, fmt.Errorf("new system admin. argument 'rf' is nil: %w", libdomain.ErrInvalidArgument)
	}
	orgRepo := rf.NewOrganizationRepository(ctx)
	appUserRepo := rf.NewAppUserRepository(ctx)

	m := &SystemAdmin{
		SystemAdminModel: domain.NewSystemAdminModel(),
		rf:               rf,
		orgRepo:          orgRepo,
		appUserRepo:      appUserRepo,
		logger:           slog.Default().With(slog.String(liblog.LoggerNameKey, "SystemAdmin")),
	}

	return m, nil
}

func (m *SystemAdmin) AppUserID() *domain.AppUserID {
	return m.SystemAdminModel.AppUserID
}
func (m *SystemAdmin) IsSystemAdmin() bool {
	return true
}

func (m *SystemAdmin) FindSystemOwnerByOrganizationID(ctx context.Context, organizationID *domain.OrganizationID) (*SystemOwner, error) {
	sysOwner, err := m.appUserRepo.FindSystemOwnerByOrganizationID(ctx, m, organizationID)
	if err != nil {
		return nil, liberrors.Errorf("m.appUserRepo.FindSystemOwnerByOrganizationID. error: %w", err)
	}

	return sysOwner, nil
}

func (m *SystemAdmin) FindSystemOwnerByOrganizationName(ctx context.Context, organizationName string) (*SystemOwner, error) {
	sysOwner, err := m.appUserRepo.FindSystemOwnerByOrganizationName(ctx, m, organizationName)
	if err != nil {
		return nil, liberrors.Errorf("find system owner by organization name: %w", err)
	}

	return sysOwner, nil
}

func (m *SystemAdmin) FindOrganizationByID(ctx context.Context, organizationID *domain.OrganizationID) (*Organization, error) {
	org, err := m.orgRepo.FindOrganizationByID(ctx, m, organizationID)
	if err != nil {
		return nil, liberrors.Errorf("m.orgRepo.FindOrganizationByID: %w", err)
	}

	return org, nil
}

func (m *SystemAdmin) FindOrganizationByName(ctx context.Context, name string) (*Organization, error) {
	org, err := m.orgRepo.FindOrganizationByName(ctx, m, name)
	if err != nil {
		return nil, liberrors.Errorf("m.orgRepo.FindOrganizationByName. error: %w", err)
	}

	return org, nil
}

func (m *SystemAdmin) addSystemOwnerToOrganization(ctx context.Context, authorizationManager AuthorizationManager, organizationID *domain.OrganizationID, organizationName string) (*SystemOwner, error) {
	systemOwnerID, err := m.appUserRepo.AddSystemOwner(ctx, m, organizationID)
	if err != nil {
		return nil, liberrors.Errorf("failed to AddSystemOwner. error: %w", err)
	}

	systemOwner, err := m.appUserRepo.FindSystemOwnerByOrganizationName(ctx, m, organizationName)
	if err != nil {
		return nil, liberrors.Errorf("failed to FindSystemOwnerByOrganizationName. error: %w", err)
	}

	// 3. add policy to "system-owner" user
	rbacSystemOwner := NewRBACAppUser(organizationID, systemOwnerID)
	rbacAllUserRolesObject := NewRBACAllUserRolesObject(organizationID)
	// - "system-owner" user "can" "set" "all-user-roles"
	if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, m, organizationID, rbacSystemOwner, RBACSetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
		return nil, err
	}

	// - "system-owner" user "can" "unset" "all-user-roles"
	if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, m, organizationID, rbacSystemOwner, RBACUnsetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
		return nil, err
	}

	return systemOwner, nil
}

func (m *SystemAdmin) AddOrganization(ctx context.Context, param OrganizationAddParameterInterface) (*domain.OrganizationID, error) {
	// 1. add organization
	organizationID, err := m.orgRepo.AddOrganization(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("failed to AddOrganization. error: %w", err)
	}

	userGroupRepo := m.rf.NewUserGroupRepository(ctx)

	// // add system-owner-group
	// systemOwnerGroupID, err := userGroupRepo.AddSystemOwnerGroup(ctx, m, organizationID)
	// if err != nil {
	// 	return nil, liberrors.Errorf("userGroupRepo.AddSystemOwnerRole. error: %w", err)
	// }

	authorizationManager, err := m.rf.NewAuthorizationManager(ctx)
	if err != nil {
		return nil, liberrors.Errorf("failed to NewAuthorizationManager. error: %w", err)
	}

	// 2. add "system-owner" user
	// 3. add policy to "system-owner" user
	systemOwner, err := m.addSystemOwnerToOrganization(ctx, authorizationManager, organizationID, param.Name())
	if err != nil {
		return nil, liberrors.Errorf("failed to addSystemOwnertoOrganization. error: %w", err)
	}

	// rbacRepo := m.rf.NewRBACRepository(ctx)
	// rbacDomain := NewRBACOrganization(organizationID)

	// // 3. add policy to "system-owner" user
	// rbacSystemOwner := NewRBACAppUser(organizationID, systemOwnerID)
	rbacAllUserRolesObject := NewRBACAllUserRolesObject(organizationID)
	// // - "system-owner" user "can" "set" "all-user-roles"
	// if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, m, organizationID, rbacSystemOwner, RBACSetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, err
	// }

	// // - "system-owner" user "can" "unset" "all-user-roles"
	// if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, m, organizationID, rbacSystemOwner, RBACUnsetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, err
	// }

	// // "system-owner" "can" "set" "all-user-roles"
	// if err := rbacRepo.AddPolicy(rbacDomain, rbacAppUser, RBACSetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	// }

	// // "system-owner" "can" "unset" "all-user-roles"
	// if err := rbacRepo.AddPolicy(rbacDomain, rbacAppUser, RBACUnsetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	// }

	// pairOfUserAndGroup := m.rf.NewPairOfUserAndGroupRepository(ctx)

	// if err := authorizationManager.AddUserToGroupBySystemAdmin(ctx, m, organizationID, systemOwnerID, systemOwnerGroupID); err != nil {
	// 	return nil, err
	// }
	// // systen-owner belongs to system-owner-group
	// if err := pairOfUserAndGroup.AddPairOfUserAndGroupToSystemOwner(ctx, m, systemOwner, systemOwnerGroupID); err != nil {
	// 	return nil, err
	// }

	// 4. add owner-group
	if _, err := userGroupRepo.AddOwnerGroup(ctx, systemOwner, organizationID); err != nil {
		return nil, err
	}

	// 5. add policty to "owner" group
	ownerGroup, err := userGroupRepo.FindUserGroupByKey(ctx, systemOwner, OwnerGroupKey)
	if err != nil {
		return nil, err
	}

	rbacOwnerGroup := NewRBACUserRole(organizationID, ownerGroup.UserGroupID())
	// - "owner" group "can" "set" "all-user-roles"
	if err := authorizationManager.AddPolicyToGroupBySystemAdmin(ctx, m, organizationID, rbacOwnerGroup, RBACSetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
		return nil, err
	}

	// - "owner" group "can" "unset" "all-user-roles"
	if err := authorizationManager.AddPolicyToGroupBySystemAdmin(ctx, m, organizationID, rbacOwnerGroup, RBACUnsetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
		return nil, err
	}

	// 6. add first owner
	ownerID, err := systemOwner.AddFirstOwner(ctx, param.FirstOwner())
	if err != nil {
		return nil, liberrors.Errorf("m.initFirstOwner. error: %w", err)
	}

	m.logger.InfoContext(ctx, fmt.Sprintf("SystemOwnerID:%d, ownerID: %d", systemOwner.AppUserID().Int(), ownerID.Int()))

	return organizationID, nil
}

func NewRBACOrganization(organizationID *domain.OrganizationID) domain.RBACDomain {
	return domain.NewRBACDomain(fmt.Sprintf("domain:%d", organizationID.Int()))
}

func NewRBACAppUser(organizationID *domain.OrganizationID, appUserID *domain.AppUserID) domain.RBACUser {
	return domain.NewRBACUser(fmt.Sprintf("user:%d", appUserID.Int()))
}

//	func NewRBACUserRole(userRoleID domain.UserGroupID) domain.RBACRole {
//		return domain.NewRBACRole(fmt.Sprintf("role_%d", userRoleID.Int()))
//	}
func NewRBACUserRole(organizationID *domain.OrganizationID, userGroupID *domain.UserGroupID) domain.RBACRole {
	return domain.NewRBACRole(fmt.Sprintf("domain:%d,role:%d", organizationID.Int(), userGroupID.Int()))
}

func NewRBACUserRoleObject(organizationID *domain.OrganizationID, userRoleID *domain.UserGroupID) domain.RBACObject {
	return domain.NewRBACObject(fmt.Sprintf("domain:%d,role:%d", organizationID.Int(), userRoleID.Int()))
}

func NewRBACAllUserRolesObject(organizationID *domain.OrganizationID) domain.RBACObject {
	return domain.NewRBACObject(fmt.Sprintf("domain:%d,role:*", organizationID.Int()))
}

var RBACSetAction = domain.NewRBACAction("Set")
var RBACUnsetAction = domain.NewRBACAction("Unset")

var RBACAllowEffect = domain.NewRBACEffect("allow")
var RBACDenyEffect = domain.NewRBACEffect("deny")
