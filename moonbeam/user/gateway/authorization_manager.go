package gateway

import (
	"context"

	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type authorizationManager struct {
	dialect  libgateway.DialectRDBMS
	db       *gorm.DB
	rf       service.RepositoryFactory
	rbacRepo service.RBACRepository
}

func NewAuthorizationManager(ctx context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) (service.AuthorizationManager, error) {
	rbacRepo, err := newRBACRepository(ctx, db)
	if err != nil {
		return nil, err
	}

	return &authorizationManager{
		dialect:  dialect,
		db:       db,
		rf:       rf,
		rbacRepo: rbacRepo,
	}, nil
}

// func (m *authorizationManager) Init(ctx context.Context) error {
// 	rbacRepo, err := newRBACRepository(ctx, m.db)
// 	if err != nil {
// 		return err
// 	}
// 	m.rbacRepo = rbacRepo
// 	return m.rbacRepo.Init()
// }

func (m *authorizationManager) AddUserToGroupBySystemAdmin(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error {
	pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, m.dialect, m.db, m.rf)

	if err := pairOfUserAndGroupRepo.AddPairOfUserAndGroupBySystemAdmin(ctx, operator, organizationID, appUserID, userGroupID); err != nil {
		return liberrors.Errorf("AddPairOfUserAndGroupBySystemAdmin: %w", err)
	}

	rbacUser := domain.NewRBACUserFromUser(appUserID)
	rbacUserRole := domain.NewRBACRoleFromGroup(organizationID, userGroupID)
	rbacDomain := domain.NewRBACDomainFromOrganization(organizationID)

	// app-user belongs to user-role
	if err := m.rbacRepo.AddSubjectGroupingPolicy(ctx, rbacDomain, rbacUser, rbacUserRole); err != nil {
		return liberrors.Errorf("rbacRepo.AddSubjectGroupingPolicy. err: %w", err)
	}

	return nil
}

func (m *authorizationManager) AddUserToGroup(ctx context.Context, operator service.UserInterface, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error {
	pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, m.dialect, m.db, m.rf)

	if err := pairOfUserAndGroupRepo.AddPairOfUserAndGroup(ctx, operator, appUserID, userGroupID); err != nil {
		return liberrors.Errorf("AddPairOfUserAndGroup: %w", err)
	}

	organizationID := operator.GetOrganizationID()

	rbacUser := domain.NewRBACUserFromUser(appUserID)
	rbacUserRole := domain.NewRBACRoleFromGroup(organizationID, userGroupID)
	rbacDomain := domain.NewRBACDomainFromOrganization(organizationID)

	// app-user belongs to user-role
	if err := m.rbacRepo.AddSubjectGroupingPolicy(ctx, rbacDomain, rbacUser, rbacUserRole); err != nil {
		return liberrors.Errorf("rbacRepo.AddNamedGroupingPolicy. err: %w", err)
	}

	return nil
}

func (m *authorizationManager) AddPolicyToUser(ctx context.Context, operator service.UserInterface, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	rbacDomain := domain.NewRBACDomainFromOrganization(operator.GetOrganizationID())

	if err := m.rbacRepo.AddPolicy(ctx, rbacDomain, subject, action, object, effect); err != nil {
		return liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	}

	return nil
}

func (m *authorizationManager) AddPolicyToUserBySystemAdmin(ctx context.Context, _ service.SystemAdminInterface, organizationID *domain.OrganizationID, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	rbacDomain := domain.NewRBACDomainFromOrganization(organizationID)

	if err := m.rbacRepo.AddPolicy(ctx, rbacDomain, subject, action, object, effect); err != nil {
		return liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	}

	return nil
}
func (m *authorizationManager) AddPolicyToUserBySystemOwner(ctx context.Context, operator service.SystemOwnerInterface, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	organizationID := operator.GetOrganizationID()
	rbacDomain := domain.NewRBACDomainFromOrganization(organizationID)

	if err := m.rbacRepo.AddPolicy(ctx, rbacDomain, subject, action, object, effect); err != nil {
		return liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	}

	return nil
}

func (m *authorizationManager) AddPolicyToGroup(ctx context.Context, operator service.UserInterface, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	rbacDomain := domain.NewRBACDomainFromOrganization(operator.GetOrganizationID())

	if err := m.rbacRepo.AddPolicy(ctx, rbacDomain, subject, action, object, effect); err != nil {
		return liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	}

	return nil
}

func (m *authorizationManager) AddPolicyToGroupBySystemAdmin(ctx context.Context, _ service.SystemAdminInterface, organizationID *domain.OrganizationID, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	rbacDomain := domain.NewRBACDomainFromOrganization(organizationID)

	if err := m.rbacRepo.AddPolicy(ctx, rbacDomain, subject, action, object, effect); err != nil {
		return liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	}

	return nil
}

func (m *authorizationManager) AddObjectToObject(ctx context.Context, operator service.SystemOwnerInterface, child, parent domain.RBACObject) error {
	rbacDomain := domain.NewRBACDomainFromOrganization(operator.GetOrganizationID())

	if err := m.rbacRepo.AddObjectGroupingPolicy(ctx, rbacDomain, child, parent); err != nil {
		return liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
	}

	return nil
}

func (m *authorizationManager) CheckAuthorization(ctx context.Context, operator service.UserInterface, rbacAction domain.RBACAction, rbacObject domain.RBACObject) (bool, error) {
	rbacDomain := domain.NewRBACDomainFromOrganization(operator.GetOrganizationID())

	userGroupRepo := m.rf.NewUserGroupRepository(ctx)
	userGroups, err := userGroupRepo.FindAllUserGroups(ctx, operator)
	if err != nil {
		return false, liberrors.Errorf("FindAllUserGroups: %w", err)
	}

	rbacRoles := make([]domain.RBACRole, 0)
	for _, userGroup := range userGroups {
		rbacRoles = append(rbacRoles, domain.NewRBACRoleFromGroup(operator.GetOrganizationID(), userGroup.UserGroupID))
	}

	rbacOperator := domain.NewRBACUserFromUser(operator.GetUserID())
	e, err := m.rbacRepo.NewEnforcerWithGroupsAndUsers(ctx, rbacRoles, []domain.RBACUser{rbacOperator})
	if err != nil {
		return false, liberrors.Errorf("NewEnforcerWithGroupsAndUsers: %w", err)
	}

	ok, err := e.Enforce(rbacOperator.Subject(), rbacObject.Object(), rbacAction.Action(), rbacDomain.Domain())
	if err != nil {
		return false, liberrors.Errorf("Enforce: %w", err)
	}

	return ok, nil
}
