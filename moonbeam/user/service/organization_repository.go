//go:generate mockery --output mock --name OrganizationRepository
package service

import (
	"context"
	"errors"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var ErrOrganizationNotFound = errors.New("organization not found")
var ErrOrganizationAlreadyExists = errors.New("organization already exists")

// type OrganizationAddParameterInterface interface {
// 	Name() string
// 	FirstOwner() *AddUserParameter
// }

type AddOrganizationParameter struct {
	Name       string `validate:"required"`
	FirstOwner *AddUserParameter
}

func NewOrganizationAddParameter(name string, firstOwner *AddUserParameter) (*AddOrganizationParameter, error) {
	m := &AddOrganizationParameter{
		Name:       name,
		FirstOwner: firstOwner,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

// func (p *AddOrganizationParameter) Name() string {
// 	return p.NameInternal
// }
// func (p *OrganizationAddParameter) FirstOwner() *AddUserParameter {
// 	return p.FirstOwnerInternal
// }

type OrganizationRepository interface {
	GetOrganization(ctx context.Context, operator UserInterface) (*Organization, error)

	FindOrganizationByName(ctx context.Context, operator SystemAdminInterface, name string) (*Organization, error)

	FindOrganizationByID(ctx context.Context, operator SystemAdminInterface, id *domain.OrganizationID) (*Organization, error)

	AddOrganization(ctx context.Context, operator SystemAdminInterface, param *AddOrganizationParameter) (*domain.OrganizationID, error)

	// FindOrganizationByName(ctx context.Context, operator SystemAdmin, name string) (Organization, error)
	// FindOrganization(ctx context.Context, operator User) (Organization, error)
}
