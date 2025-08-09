package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type PairOfUserAndGroupRepository interface {
	AddPairOfUserAndGroupBySystemAdmin(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error

	AddPairOfUserAndGroup(ctx context.Context, operator AppUserInterface, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error

	RemovePairOfUserAndGroup(ctx context.Context, operator AppUserInterface, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error

	FindUserGroupsByUserID(ctx context.Context, operator AppUserInterface, appUserID *domain.AppUserID) ([]*domain.UserGroupModel, error)
}
