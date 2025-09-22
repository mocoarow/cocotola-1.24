package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type UserID struct {
	Value int `validate:"required,gte=0"`
}

func NewUserID(value int) (*UserID, error) {
	return &UserID{
		Value: value,
	}, nil
}

func (v *UserID) Int() int {
	return v.Value
}
func (v *UserID) IsUserID() bool {
	return true
}
func (v *UserID) GetRBACSubject() RBACSubject {
	return NewRBACUserFromUser(v)
}

type User struct {
	*libdomain.BaseModel
	UserID         *UserID         `validate:"required"`
	OrganizationID *OrganizationID `validate:"required"`
	LoginID        string          `validate:"required"`
	Username       string          `validate:"required"`
	UserGroups     []*UserGroup
}

func NewUser(baseModel *libdomain.BaseModel, userID *UserID, organizationID *OrganizationID, loginID, username string, userGroups []*UserGroup) (*User, error) {
	m := &User{
		BaseModel:      baseModel,
		UserID:         userID,
		OrganizationID: organizationID,
		LoginID:        loginID,
		Username:       username,
		UserGroups:     userGroups,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate user model: %w", err)
	}

	return m, nil
}

func (m *User) GetUserID() *UserID {
	return m.UserID
}
func (m *User) GetOrganizationID() *OrganizationID {
	return m.OrganizationID
}
