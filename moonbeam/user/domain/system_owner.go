package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type SystemOwner struct {
	*Owner `validate:"required"`
}

func NewSystemOwner(user *Owner) (*SystemOwner, error) {
	m := &SystemOwner{
		Owner: user,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate system owner: %w", err)
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
	return m.Owner.GetOrganizationID()
}
func (m *SystemOwner) GetUserID() *UserID {
	return m.Owner.GetUserID()
}
