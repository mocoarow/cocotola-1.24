package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type OrganizationID struct {
	Value int `validate:"required,gte=1"`
}

func NewOrganizationID(value int) (*OrganizationID, error) {
	return &OrganizationID{
		Value: value,
	}, nil
}

func (v *OrganizationID) Int() int {
	return v.Value
}
func (v *OrganizationID) IsOrganizationID() bool {
	return true
}
func (v *OrganizationID) GetRBACDomain() RBACDomain {
	return NewRBACDomainFromOrganization(v)
}

type Organization struct {
	*libdomain.BaseModel
	OrganizationID *OrganizationID `validate:"required"`
	Name           string          `validate:"required"`
}

func NewOrganization(basemodel *libdomain.BaseModel, organizationID *OrganizationID, name string) (*Organization, error) {
	m := &Organization{
		BaseModel:      basemodel,
		OrganizationID: organizationID,
		Name:           name,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate organization: %w", err)
	}

	return m, nil
}
