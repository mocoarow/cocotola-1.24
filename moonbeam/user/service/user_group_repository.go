package service

import (
	"context"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

// type UserGroupAddParameterInterface interface {
// 	Key() string
// 	Name() string
// 	Description() string
// }

type AddUserGroupParameter struct {
	Key         string
	Name        string
	Description string
}

func NewUserGroupAddParameter(key, name, description string) (*AddUserGroupParameter, error) {
	m := &AddUserGroupParameter{
		Key:         key,
		Name:        name,
		Description: description,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

// func (p *UserGroupAddParameter) Key() string {
// 	return p.KeyInternal
// }
// func (p *UserGroupAddParameter) Name() string {
// 	return p.NameInternal
// }
// func (p *UserGroupAddParameter) Description() string {
// 	return p.DescriptionInternal
// }

type UserGroupRepository interface {
	FindAllUserGroups(ctx context.Context, operator domain.UserInterface) ([]*domain.UserGroupModel, error)

	FindSystemOwnerGroup(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroupModel, error)

	FindUserGroupByKey(ctx context.Context, operator domain.UserInterface, key string) (*domain.UserGroupModel, error)
	FindUserGroupByID(ctx context.Context, operator domain.UserInterface, userGroupID *domain.UserGroupID) (*domain.UserGroupModel, error)
	AddOwnerGroup(ctx context.Context, operator domain.SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)
	AddPublicGroup(ctx context.Context, operator domain.SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)

	AddSystemOwnerGroup(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)

	AddUserGroup(ctx context.Context, operator domain.OwnerInterface, parameter *AddUserGroupParameter) (*domain.UserGroupID, error)
}
