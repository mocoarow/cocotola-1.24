package service

import (
	"context"
	"errors"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var ErrAppUserNotFound = errors.New("AppUser not found")
var ErrAppUserAlreadyExists = errors.New("AppUser already exists")

var ErrSystemOwnerNotFound = errors.New("SystemOwner not found")

// type AppUserAddParameterInterface interface {
// 	LoginID() string
// 	Username() string
// 	Password() string
// 	Provider() string
// 	ProviderLoginID() string
// 	ProviderAuthToken() string
// 	ProviderRefreshToken() string
// }

type AddAppUserParameter struct {
	LoginID              string
	Username             string
	Password             string
	Provider             string
	ProviderLoginID      string
	ProviderAuthToken    string
	providerRefreshToken string
}

func NewAppUserAddParameter(loginID, username, password, provider, providerLoginID, providerAuthToken, providerRefreshToken string) (*AddAppUserParameter, error) {
	m := &AddAppUserParameter{
		LoginID:              loginID,
		Username:             username,
		Password:             password,
		Provider:             provider,
		ProviderLoginID:      providerLoginID,
		ProviderAuthToken:    providerAuthToken,
		providerRefreshToken: providerRefreshToken,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

// func (p *AppUserAddParameter) LoginID() string {
// 	return p.LoginID
// }
// func (p *AppUserAddParameter) Username() string {
// 	return p.Username
// }
// func (p *AppUserAddParameter) Password() string {
// 	return p.Password
// }
// func (p *AppUserAddParameter) Provider() string {
// 	return p.Provider
// }
// func (p *AppUserAddParameter) ProviderLoginID() string {
// 	return p.ProviderLoginID
// }
// func (p *AppUserAddParameter) ProviderAuthToken() string {
// 	return p.ProviderAuthToken
// }
// func (p *AppUserAddParameter) ProviderRefreshToken() string {
// 	return p.providerRefreshToken
// }

type Option string

var IncludeGroups Option = "IncludeGroups"

type AppUserRepository interface {
	FindSystemOwnerByOrganizationID(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID) (*SystemOwner, error)

	FindSystemOwnerByOrganizationName(ctx context.Context, operator SystemAdminInterface, organizationName string, options ...Option) (*SystemOwner, error)

	FindAppUserByID(ctx context.Context, operator AppUserInterface, id *domain.AppUserID, options ...Option) (*AppUser, error)

	FindAppUserByLoginID(ctx context.Context, operator AppUserInterface, loginID string) (*AppUser, error)

	FindOwnerByLoginID(ctx context.Context, operator SystemOwnerInterface, loginID string) (*Owner, error)

	AddAppUser(ctx context.Context, operator OwnerModelInterface, param *AddAppUserParameter) (*domain.AppUserID, error)

	AddSystemOwner(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.AppUserID, error)

	// VerifyPassword(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, loginID, password string) (bool, error)

	VerifyPassword(ctx context.Context, operator SystemOwnerInterface, loginID, password string) (bool, error)

	// AddFirstOwner(ctx context.Context, operator domain.SystemOwnerModel, param FirstOwnerAddParameter) (domain.AppUserID, error)

	// FindAppUserIDs(ctx context.Context, operator domain.SystemOwnerModel, pageNo, pageSize int) ([]domain.AppUserID, error)
}
