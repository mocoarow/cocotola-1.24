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
	userRepo   UserRepository
	userGroupRepo UserGroupRepository
	spaceRepo     SpaceRepository
	// pairOfUserAndGroup PairOfUserAndGroupRepository
	// rbacRepo             RBACRepository
	authorizationManager AuthorizationManager
	userEventHandler  libservice.ResourceEventHandler
	spaceEventHandler    libservice.ResourceEventHandler
	logger               *slog.Logger
}

func NewSystemOwner(ctx context.Context, rf RepositoryFactory, systemOwnerModel *domain.SystemOwnerModel) (*SystemOwner, error) {
	orgRepo := rf.NewOrganizationRepository(ctx)
	userRepo := rf.NewUserRepository(ctx)
	userGroupRepo := rf.NewUserGroupRepository(ctx)
	spaceRepo := rf.NewSpaceRepository(ctx)
	// pairOfUserAndGroup := rf.NewPairOfUserAndGroupRepository(ctx)
	// rbacRepo := rf.NewRBACRepository(ctx)
	authorizationManager, err := rf.NewAuthorizationManager(ctx)
	if err != nil {
		return nil, liberrors.Errorf("NewAuthorizationManager: %w", err)
	}
	userEventHandler := rf.NewUserEventHandler(ctx)
	spaceEventHandler := rf.NewSpaceEventHandler(ctx)

	m := &SystemOwner{
		SystemOwnerModel: systemOwnerModel,
		orgRepo:          orgRepo,
		userRepo:      userRepo,
		userGroupRepo:    userGroupRepo,
		spaceRepo:        spaceRepo,
		// pairOfUserAndGroup:   pairOfUserAndGroup,
		// rbacRepo:             rbacRepo,
		authorizationManager: authorizationManager,
		userEventHandler:  userEventHandler,
		spaceEventHandler:    spaceEventHandler,
		logger:               slog.Default().With(slog.String(liblog.LoggerNameKey, "SystemOwner")),
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *SystemOwner) GetUserID() *domain.UserID {
	return m.UserModel.UserID
}
func (m *SystemOwner) GetOrganizationID() *domain.OrganizationID {
	return m.UserModel.OrganizationID
}

//	func (m *SystemOwner) LoginID() string {
//		return m.UserModel.LoginID
//	}
//
//	func (m *SystemOwner) Username() string {
//		return m.UserModel.Username
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

func (m *SystemOwner) FindUserByID(ctx context.Context, id *domain.UserID) (*User, error) {
	user, err := m.userRepo.FindUserByID(ctx, m, id)
	if err != nil {
		return nil, liberrors.Errorf("m.userRepo.FindUserByID. err: %w", err)
	}

	return user, nil
}

func (m *SystemOwner) FindUserByLoginID(ctx context.Context, loginID string) (*User, error) {
	user, err := m.userRepo.FindUserByLoginID(ctx, m, loginID)
	if err != nil {
		return nil, liberrors.Errorf("m.userRepo.FindUserByLoginID. err: %w", err)
	}

	return user, nil
}

func (m *SystemOwner) AddFirstOwner(ctx context.Context, param *AddUserParameter) (*domain.UserID, error) {
	// rbacUser := NewRBACUser(m.GetOrganizationID(), m.GetUserID())
	rbacAllUserRolesObject := domain.NewRBACAllUserRolesObjectFromOrganization(m.GetOrganizationID())

	// Can "the operator" "set" "all-user-roles" ?
	ok, err := m.authorizationManager.CheckAuthorization(ctx, m, RBACSetAction, rbacAllUserRolesObject)
	if err != nil {
		return nil, liberrors.Errorf("CheckAuthorization: %w", err)
	} else if !ok {
		return nil, libdomain.ErrPermissionDenied
	}

	// add owner
	firstOwnerID, err := m.userRepo.AddUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("AddUser: %w", err)
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

	// rbacDomain := NewRBACDomainFromOrganization(m.GetOrganizationID())

	// // "owner" "can" "set" "all-user-roles"
	// if err := m.rbacRepo.AddPolicy(rbacDomain, rbacUser, RBACSetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	// }

	// // "owner" "can" "unset" "all-user-roles"
	// if err := m.rbacRepo.AddPolicy(rbacDomain, rbacUser, RBACUnsetAction, rbacAllUserRolesObject, RBACAllowEffect); err != nil {
	// 	return nil, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	// }

	return firstOwnerID, nil
}

func (m *SystemOwner) AddUser(ctx context.Context, param *AddUserParameter) (*domain.UserID, error) {
	m.logger.InfoContext(ctx, "AddStudent")
	userID, err := m.userRepo.AddUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("m.userRepo.AddUser. err: %w", err)
	}

	go m.userEventHandler.OnAdd(context.Background(), map[string]int{
		"organizationId": m.GetOrganizationID().Int(),
		"userId":      userID.Int(),
	})

	return userID, nil
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
	ok, err := m.userRepo.VerifyPassword(ctx, m, loginID, password)
	if err != nil {
		return false, liberrors.Errorf("m.userRepo.VerifyPassword. err: %w", err)
	}

	return ok, nil
}
