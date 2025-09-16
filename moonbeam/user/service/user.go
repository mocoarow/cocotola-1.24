package service

import (
	"context"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type OperatorInterface interface {
	GetUserID() *domain.UserID
	GetOrganizationID() *domain.OrganizationID
}

type UserInterface interface {
	GetUserID() *domain.UserID
	GetOrganizationID() *domain.OrganizationID
	// LoginID() string
	// Username() string
	// GetUserGroups() []domain.UserGroupModel
}

type User struct {
	*domain.UserModel
}

func NewUser(_ context.Context, rf RepositoryFactory, userModel *domain.UserModel) (*User, error) {
	if rf == nil {
		return nil, liberrors.Errorf("rf is nil. err: %w", libdomain.ErrInvalidArgument)
	}
	if userModel == nil {
		return nil, liberrors.Errorf("userModel is nil. err: %w", libdomain.ErrInvalidArgument)
	}

	m := &User{
		UserModel: userModel,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *User) GetUserID() *domain.UserID {
	return m.UserModel.UserID
}
func (m *User) GetOrganizationID() *domain.OrganizationID {
	return m.UserModel.OrganizationID
}

// func (m *User) LoginID() string {
// 	return m.UserModel.LoginID
// }
// func (m *User) Username() string {
// 	return m.UserModel.Username
// }
