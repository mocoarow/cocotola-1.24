package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type AppUserID struct {
	Value int `validate:"required,gte=0"`
}

func NewAppUserID(value int) (*AppUserID, error) {
	return &AppUserID{
		Value: value,
	}, nil
}

func (v *AppUserID) Int() int {
	return v.Value
}
func (v *AppUserID) IsAppUserID() bool {
	return true
}
func (v *AppUserID) GetRBACSubject() RBACSubject {
	return NewRBACAppUser(v)
}

type AppUserModel struct {
	*libdomain.BaseModel
	AppUserID      *AppUserID
	OrganizationID *OrganizationID
	LoginID        string `validate:"required"`
	Username       string `validate:"required"`
	UserGroups     []*UserGroupModel
}

func NewAppUserModel(baseModel *libdomain.BaseModel, appUserID *AppUserID, organizationID *OrganizationID, loginID, username string, userGroups []*UserGroupModel) (*AppUserModel, error) {
	m := &AppUserModel{
		BaseModel:      baseModel,
		AppUserID:      appUserID,
		OrganizationID: organizationID,
		LoginID:        loginID,
		Username:       username,
		UserGroups:     userGroups,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate app user model: %w", err)
	}

	return m, nil
}
