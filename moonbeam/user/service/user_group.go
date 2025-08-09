package service

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

// type UserGroup interface {
// 	domain.UserGroupModel
// }

type UserGroup struct {
	*domain.UserGroupModel
}

// NewUserGroup returns a new UserGroup
func NewUserGroup(userGroupModel *domain.UserGroupModel) (*UserGroup, error) {
	m := &UserGroup{
		userGroupModel,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *UserGroup) UserGroupID() *domain.UserGroupID {
	return m.UserGroupModel.UserGroupID
}
func (m *UserGroup) OrganizationID() *domain.OrganizationID {
	return m.UserGroupModel.OrganizationID
}
func (m *UserGroup) Key() string {
	return m.UserGroupModel.Key
}
func (m *UserGroup) Name() string {
	return m.UserGroupModel.Name
}
func (m *UserGroup) Description() string {
	return m.UserGroupModel.Description
}
