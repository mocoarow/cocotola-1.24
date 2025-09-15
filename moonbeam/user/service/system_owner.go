package service

import (
	"context"
	"log/slog"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

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
	spaceRepo     SpaceRepository
	// pairOfUserAndGroup PairOfUserAndGroupRepository
	// rbacRepo             RBACRepository
	authorizationManager AuthorizationManager
	appUserEventHandler  libservice.ResourceEventHandler
	spaceEventHandler    libservice.ResourceEventHandler
	logger               *slog.Logger
}

func NewSystemOwner(ctx context.Context, rf RepositoryFactory, systemOwnerModel *domain.SystemOwnerModel) (*SystemOwner, error) {
	orgRepo := rf.NewOrganizationRepository(ctx)
	appUserRepo := rf.NewAppUserRepository(ctx)
	userGroupRepo := rf.NewUserGroupRepository(ctx)
	spaceRepo := rf.NewSpaceRepository(ctx)
	// pairOfUserAndGroup := rf.NewPairOfUserAndGroupRepository(ctx)
	// rbacRepo := rf.NewRBACRepository(ctx)
	authorizationManager, err := rf.NewAuthorizationManager(ctx)
	if err != nil {
		return nil, liberrors.Errorf("NewAuthorizationManager: %w", err)
	}
	appUserEventHandler := rf.NewAppUserEventHandler(ctx)
	spaceEventHandler := rf.NewSpaceEventHandler(ctx)

	m := &SystemOwner{
		SystemOwnerModel: systemOwnerModel,
		orgRepo:          orgRepo,
		appUserRepo:      appUserRepo,
		userGroupRepo:    userGroupRepo,
		spaceRepo:        spaceRepo,
		// pairOfUserAndGroup:   pairOfUserAndGroup,
		// rbacRepo:             rbacRepo,
		authorizationManager: authorizationManager,
		appUserEventHandler:  appUserEventHandler,
		spaceEventHandler:    spaceEventHandler,
		logger:               slog.Default().With(slog.String(liblog.LoggerNameKey, "SystemOwner")),
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *SystemOwner) GetAppUserID() *domain.AppUserID {
	return m.AppUserModel.AppUserID
}
func (m *SystemOwner) GetOrganizationID() *domain.OrganizationID {
	return m.AppUserModel.OrganizationID
}

//	func (m *SystemOwner) LoginID() string {
//		return m.AppUserModel.LoginID
//	}
//
//	func (m *SystemOwner) Username() string {
//		return m.AppUserModel.Username
//	}
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

func (m *SystemOwner) GetPublidDefaultSpace(ctx context.Context) (*Space, error) {
	space, err := m.spaceRepo.FindPublicSpaceByKey(ctx, PublicDefaultSpaceKey)
	if err != nil {
		return nil, liberrors.Errorf("m.spaceRepo.FindPublicSpaceByKey. err: %w", err)
	}

	return space, nil
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

func (m *SystemOwner) AddFirstOwner(ctx context.Context, param *AddAppUserParameter) (*domain.AppUserID, error) {
	// rbacAppUser := NewRBACAppUser(m.GetOrganizationID(), m.GetAppUserID())
	rbacAllUserRolesObject := domain.NewRBACAllUserRolesObject(m.GetOrganizationID())

	// Can "the operator" "set" "all-user-roles" ?
	ok, err := m.authorizationManager.CheckAuthorization(ctx, m, RBACSetAction, rbacAllUserRolesObject)
	if err != nil {
		return nil, liberrors.Errorf("CheckAuthorization: %w", err)
	} else if !ok {
		return nil, libdomain.ErrPermissionDenied
	}

	// add owner
	firstOwnerID, err := m.appUserRepo.AddAppUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("AddAppUser: %w", err)
	}

	ownerGroup, err := m.userGroupRepo.FindUserGroupByKey(ctx, m, OwnerGroupKey)
	if err != nil {
		return nil, liberrors.Errorf("FindUserGroupByKey: %w", err)
	}

	// add owner to owner-group
	if err := m.authorizationManager.AddUserToGroup(ctx, m, firstOwnerID, ownerGroup.UserGroupID); err != nil {
		return nil, liberrors.Errorf("AddUserToGroup: %w", err)
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

func (m *SystemOwner) AddAppUser(ctx context.Context, param *AddAppUserParameter) (*domain.AppUserID, error) {
	m.logger.InfoContext(ctx, "AddStudent")
	appUserID, err := m.appUserRepo.AddAppUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("m.appUserRepo.AddAppUser. err: %w", err)
	}

	go m.appUserEventHandler.OnAdd(context.Background(), map[string]int{
		"organizationId": m.GetOrganizationID().Int(),
		"appUserId":      appUserID.Int(),
	})

	return appUserID, nil
}

func (m *SystemOwner) AddSpace(ctx context.Context, param *AddSpaceParameter) (*domain.SpaceID, error) {
	spaceID, err := m.spaceRepo.AddSpace(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("spaceRepo.AddSpace. err: %w", err)
	}

	go m.spaceEventHandler.OnAdd(context.Background(), map[string]int{
		"organizationId": m.GetOrganizationID().Int(),
		"spaceId":        spaceID.Int(),
	})

	return spaceID, nil
}
func (m *SystemOwner) VerifyPassword(ctx context.Context, loginID, password string) (bool, error) {
	ok, err := m.appUserRepo.VerifyPassword(ctx, m, loginID, password)
	if err != nil {
		return false, liberrors.Errorf("m.appUserRepo.VerifyPassword. err: %w", err)
	}

	return ok, nil
}
