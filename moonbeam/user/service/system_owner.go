package service

import (
	"context"
	"log/slog"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var _ SystemOwnerInterface = (*SystemOwner)(nil)

type SystemOwnerInterface interface {
	OwnerModelInterface
	IsSystemOwner() bool
	// GetUserGroups() []domain.UserGroupModel
}

type SystemOwner struct {
	*domain.SystemOwnerModel
	orgRepo       OrganizationRepository
	appUserRepo   AppUserRepository
	userGroupRepo UserGroupRepository
	// pairOfUserAndGroup PairOfUserAndGroupRepository
	// rbacRepo             RBACRepository
	authorizationManager AuthorizationManager
	logger               *slog.Logger
}

func NewSystemOwner(ctx context.Context, rf RepositoryFactory, systemOwnerModel *domain.SystemOwnerModel) (*SystemOwner, error) {
	orgRepo := rf.NewOrganizationRepository(ctx)
	appUserRepo := rf.NewAppUserRepository(ctx)
	userGroupRepo := rf.NewUserGroupRepository(ctx)
	// pairOfUserAndGroup := rf.NewPairOfUserAndGroupRepository(ctx)
	// rbacRepo := rf.NewRBACRepository(ctx)
	authorizationManager, err := rf.NewAuthorizationManager(ctx)
	if err != nil {
		return nil, err
	}

	m := &SystemOwner{
		SystemOwnerModel: systemOwnerModel,
		orgRepo:          orgRepo,
		appUserRepo:      appUserRepo,
		userGroupRepo:    userGroupRepo,
		// pairOfUserAndGroup:   pairOfUserAndGroup,
		// rbacRepo:             rbacRepo,
		authorizationManager: authorizationManager,
		logger:               slog.Default().With(slog.String(liblog.LoggerNameKey, "SystemOwner")),
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}
func (m *SystemOwner) AppUserID() *domain.AppUserID {
	return m.AppUserModel.AppUserID
}
func (m *SystemOwner) OrganizationID() *domain.OrganizationID {
	return m.AppUserModel.OrganizationID
}
func (m *SystemOwner) LoginID() string {
	return m.AppUserModel.LoginID
}
func (m *SystemOwner) Username() string {
	return m.AppUserModel.Username
}
func (m *SystemOwner) IsOwner() bool {
	return true
}
func (m *SystemOwner) IsSystemOwner() bool {
	return true
}

func (m *SystemOwner) GetOrganization(ctx context.Context) (*Organization, error) {
	org, err := m.orgRepo.GetOrganization(ctx, m)
	if err != nil {
		return nil, liberrors.Errorf("m.orgRepo.GetOrganization. err: %w", err)
	}

	return org, nil
}

func (m *SystemOwner) FindAppUserByID(ctx context.Context, id *domain.AppUserID) (*AppUser, error) {
	appUser, err := m.appUserRepo.FindAppUserByID(ctx, m, id)
	if err != nil {
		return nil, liberrors.Errorf("m.appUserRepo.FindAppUserByID. err: %w", err)
	}

	return appUser, nil
}

func (m *SystemOwner) FindAppUserByLoginID(ctx context.Context, loginID string) (*AppUser, error) {
	appUser, err := m.appUserRepo.FindAppUserByLoginID(ctx, m, loginID)
	if err != nil {
		return nil, liberrors.Errorf("m.appUserRepo.FindAppUserByLoginID. err: %w", err)
	}

	return appUser, nil
}

func (m *SystemOwner) AddFirstOwner(ctx context.Context, param AppUserAddParameterInterface) (*domain.AppUserID, error) {
	// rbacAppUser := NewRBACAppUser(m.GetOrganizationID(), m.GetAppUserID())
	rbacAllUserRolesObject := NewRBACAllUserRolesObject(m.OrganizationID())

	// Can "the operator" "set" "all-user-roles" ?
	ok, err := m.authorizationManager.Authorize(ctx, m, RBACSetAction, rbacAllUserRolesObject)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, libdomain.ErrPermissionDenied
	}

	// add owner
	firstOwnerID, err := m.appUserRepo.AddAppUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("failed to AddFirstOwner. error: %w", err)
	}

	ownerGroup, err := m.userGroupRepo.FindUserGroupByKey(ctx, m, OwnerGroupKey)
	if err != nil {
		return nil, err
	}

	// add owner to owner-group
	if err := m.authorizationManager.AddUserToGroup(ctx, m, firstOwnerID, ownerGroup.UserGroupID()); err != nil {
		return nil, err
	}

	// add owner to owner-group
	// if err := m.pairOfUserAndGroup.AddPairOfUserAndGroup(ctx, m, ownerID, ownerGroup.GetUerGroupID()); err != nil {
	// 	return nil, err
	// }

	// rbacDomain := NewRBACOrganization(m.GetOrganizationID())

	// // "owner" "can" "set" "all-user-roles"
	// if err := m.rbacRepo.AddPolicy(rbacDomain, rbacAppUser, RBACSetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	// }

	// // "owner" "can" "unset" "all-user-roles"
	// if err := m.rbacRepo.AddPolicy(rbacDomain, rbacAppUser, RBACUnsetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	// }

	return firstOwnerID, nil
}

func (m *SystemOwner) AddAppUser(ctx context.Context, param AppUserAddParameterInterface) (*domain.AppUserID, error) {
	m.logger.InfoContext(ctx, "AddStudent")
	appUserID, err := m.appUserRepo.AddAppUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("m.appUserRepo.AddAppUser. err: %w", err)
	}

	return appUserID, nil
}
