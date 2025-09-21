package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type SystemOwner struct {
	*OwnerModel
}

func NewSystemOwnerModel(user *OwnerModel) (*SystemOwner, error) {
	m := &SystemOwner{
		OwnerModel: user,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate system owner model: %w", err)
	}

	return m, nil
}

func (m *SystemOwner) IsOwner() bool {
	return true
}
func (m *SystemOwner) IsSystemOwner() bool {
	return true
}
func (m *SystemOwner) GetOrganizationID() *OrganizationID {
	return m.OwnerModel.GetOrganizationID()
}
func (m *SystemOwner) GetUserID() *UserID {
	return m.OwnerModel.GetUserID()
}
