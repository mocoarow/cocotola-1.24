package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type PairOfUserAndGroupRepository interface {
	AddPairOfUserAndGroupBySystemAdmin(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error

	AddPairOfUserAndGroup(ctx context.Context, operator UserInterface, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error

	RemovePairOfUserAndGroup(ctx context.Context, operator UserInterface, appUserID *domain.UserID, userGroupID *domain.UserGroupID) error

	FindUserGroupsByUserID(ctx context.Context, operator UserInterface, appUserID *domain.UserID) ([]*domain.UserGroupModel, error)
}
