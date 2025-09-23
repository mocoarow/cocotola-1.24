package service

import (
	"context"
	"errors"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var ErrUserGroupNotFound = errors.New("user group not found")
var ErrUserGroupAlreadyExists = errors.New("user group already exists")

type AddUserGroupParameter struct {
	Key         string
	Name        string
	Description string
}

func NewAddUserGroupParameter(key, name, description string) (*AddUserGroupParameter, error) {
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

type UserGroupRepository interface {
	FindAllUserGroups(ctx context.Context, operator domain.UserInterface) ([]*domain.UserGroup, error)

	FindSystemOwnerGroup(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroup, error)

	FindUserGroupByKey(ctx context.Context, operator domain.UserInterface, key string) (*domain.UserGroup, error)
	FindUserGroupByID(ctx context.Context, operator domain.UserInterface, userGroupID *domain.UserGroupID) (*domain.UserGroup, error)
	AddOwnerGroup(ctx context.Context, operator domain.SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)
	AddPublicGroup(ctx context.Context, operator domain.SystemOwnerInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)

	AddSystemOwnerGroup(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserGroupID, error)

	AddUserGroup(ctx context.Context, operator domain.OwnerInterface, parameter *AddUserGroupParameter) (*domain.UserGroupID, error)
}
