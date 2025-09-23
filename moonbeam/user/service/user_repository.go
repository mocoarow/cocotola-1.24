package service

import (
	"context"
	"errors"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")

var ErrSystemOwnerNotFound = errors.New("system owner not found")

var ErrUnauthenticated = errors.New("unauthenticated")

// type UserAddParameterInterface interface {
// 	LoginID() string
// 	Username() string
// 	Password() string
// 	Provider() string
// 	ProviderLoginID() string
// 	ProviderAuthToken() string
// 	ProviderRefreshToken() string
// }

type AddUserParameter struct {
	LoginID              string
	Username             string
	Password             string
	Provider             string
	ProviderLoginID      string
	ProviderAuthToken    string
	providerRefreshToken string
}

func NewAddUserParameter(loginID, username, password, provider, providerLoginID, providerAuthToken, providerRefreshToken string) (*AddUserParameter, error) {
	m := &AddUserParameter{
		LoginID:              loginID,
		Username:             username,
		Password:             password,
		Provider:             provider,
		ProviderLoginID:      providerLoginID,
		ProviderAuthToken:    providerAuthToken,
		providerRefreshToken: providerRefreshToken,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("new add user parameter: %w", err)
	}

	return m, nil
}

// func (p *UserAddParameter) LoginID() string {
// 	return p.LoginID
// }
// func (p *UserAddParameter) Username() string {
// 	return p.Username
// }
// func (p *UserAddParameter) Password() string {
// 	return p.Password
// }
// func (p *UserAddParameter) Provider() string {
// 	return p.Provider
// }
// func (p *UserAddParameter) ProviderLoginID() string {
// 	return p.ProviderLoginID
// }
// func (p *UserAddParameter) ProviderAuthToken() string {
// 	return p.ProviderAuthToken
// }
// func (p *UserAddParameter) ProviderRefreshToken() string {
// 	return p.providerRefreshToken
// }

type Option string

var IncludeGroups Option = "IncludeGroups"

type UserRepository interface {
	FindSystemOwnerByOrganizationID(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.SystemOwner, error)

	FindSystemOwnerByOrganizationName(ctx context.Context, operator domain.SystemAdminInterface, organizationName string, options ...Option) (*domain.SystemOwner, error)

	FindUserByID(ctx context.Context, operator domain.UserInterface, id *domain.UserID, options ...Option) (*domain.User, error)

	FindUserByLoginID(ctx context.Context, operator domain.UserInterface, loginID string) (*domain.User, error)

	FindOwnerByLoginID(ctx context.Context, operator domain.SystemOwnerInterface, loginID string) (*domain.Owner, error)

	AddUser(ctx context.Context, operator domain.UserInterface, param *AddUserParameter) (*domain.UserID, error)

	AddSystemOwner(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserID, error)

	// VerifyPassword(ctx context.Context, operator SystemAdminInterface, organizationID *domain.OrganizationID, loginID, password string) (bool, error)

	VerifyPassword(ctx context.Context, operator domain.SystemOwnerInterface, loginID, password string) (bool, error)

	// AddFirstOwner(ctx context.Context, operator domain.SystemOwner, param FirstOwnerAddParameter) (domain.UserID, error)

	// FindUserIDs(ctx context.Context, operator domain.SystemOwner, pageNo, pageSize int) ([]domain.UserID, error)
}
