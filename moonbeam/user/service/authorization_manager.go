package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type AuthorizationManager interface {
	// Init(ctx context.Context) error

	AddUserToGroup(ctx context.Context, operator UserInterface, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error

	AddUserToGroupBySystemAdmin(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error

	// RemoveUserFromGroup()

	// AddGroupToGroup(ctx context.Context, operator domain.UserModel, src domain.UserGroupID, dst domain.UserGroupID) error
	AddObjectToObject(ctx context.Context, operator SystemOwnerInterface, child, parent domain.RBACObject) error

	// RemoveGroupFromGroup()

	// AddObjectToObject()

	// RemoveObjectFromObject()

	AddPolicyToUser(ctx context.Context, operator UserInterface, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error

	AddPolicyToUserBySystemAdmin(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error

	AddPolicyToUserBySystemOwner(ctx context.Context, operator SystemOwnerInterface, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error

	AddPolicyToGroup(ctx context.Context, operator UserInterface, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error

	AddPolicyToGroupBySystemAdmin(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error

	// AddPolicyToGroup()

	// RemovePolicyToGroup()

	CheckAuthorization(ctx context.Context, operator UserInterface, rbacAction domain.RBACAction, rbacObject domain.RBACObject) (bool, error)
}
