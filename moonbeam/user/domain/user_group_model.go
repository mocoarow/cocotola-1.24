package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type UserGroupID struct {
	Value int
}

func NewUserGroupID(value int) (*UserGroupID, error) {
	return &UserGroupID{
		Value: value,
	}, nil
}

func (v *UserGroupID) Int() int {
	return v.Value
}
func (v *UserGroupID) IsUserGroupID() bool {
	return true
}

type UserGroupModel struct {
	*libdomain.BaseModel
	UserGroupID    *UserGroupID
	OrganizationID *OrganizationID
	Key            string `validate:"required"`
	Name           string `validate:"required"`
	Description    string
}

// NewUserGroupModel returns a new UserGroupModel
func NewUserGroupModel(baseModel *libdomain.BaseModel, userGroupID *UserGroupID, organizationID *OrganizationID, key, name, description string) (*UserGroupModel, error) {
	m := &UserGroupModel{
		BaseModel:      baseModel,
		UserGroupID:    userGroupID,
		OrganizationID: organizationID,
		Key:            key,
		Name:           name,
		Description:    description,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate user group model: %w", err)
	}

	return m, nil
}
