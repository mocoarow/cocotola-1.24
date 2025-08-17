package service

import (
	"context"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type UserGroupAddParameterInterface interface {
	Key() string
	Name() string
	Description() string
}

type UserGroupAddParameter struct {
	KeyInternal         string
	NameInternal        string
	DescriptionInternal string
}

func NewUserGroupAddParameter(key, name, description string) (*UserGroupAddParameter, error) {
	m := &UserGroupAddParameter{
		KeyInternal:         key,
		NameInternal:        name,
		DescriptionInternal: description,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (p *UserGroupAddParameter) Key() string {
	return p.KeyInternal
}
func (p *UserGroupAddParameter) Name() string {
	return p.NameInternal
}
func (p *UserGroupAddParameter) Description() string {
	return p.DescriptionInternal
}

type UserGroupRepository interface {
	FindAllUserGroups(ctx context.Context, operator AppUserInterface) ([]*domain.UserGroupModel, error)

	FindSystemOwnerGroup(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID) (*UserGroup, error)

	FindUserGroupByKey(ctx context.Context, operator AppUserInterface, key string) (*UserGroup, error)
	FindUserGroupByID(ctx context.Context, operator AppUserInterface, userGroupID *domain.UserGroupID) (*UserGroup, error)
	AddOwnerGroup(ctx context.Context, operator SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)

	AddSystemOwnerGroup(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)

	AddUserGroup(ctx context.Context, operator OwnerModelInterface, parameter UserGroupAddParameterInterface) (*domain.UserGroupID, error)
}
