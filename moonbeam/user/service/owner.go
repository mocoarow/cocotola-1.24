package service

import (
	"context"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type OwnerModelInterface interface {
	UserInterface
	IsOwner() bool
	// GetUserGroups() []domain.UserGroupModel
}

type Owner struct {
	rf RepositoryFactory
	*domain.OwnerModel
}

func NewOwner(rf RepositoryFactory, ownerModel *domain.OwnerModel) *Owner {
	m := &Owner{
		rf:         rf,
		OwnerModel: ownerModel,
	}

	return m
}

func (m *Owner) AddUser(ctx context.Context, param *AddUserParameter) (*domain.UserID, error) {
	userRepo := m.rf.NewUserRepository(ctx)
	userID, err := userRepo.AddUser(ctx, m, param)
	if err != nil {
		return nil, liberrors.Errorf("m.userRepo.AddUser. err: %w", err)
	}

	return userID, nil
}

func (m *Owner) GetUserID() *domain.UserID {
	return m.UserModel.UserID
}
func (m *Owner) GetOrganizationID() *domain.OrganizationID {
	return m.UserModel.OrganizationID
}

//	func (m *Owner) LoginID() string {
//		return m.UserModel.LoginID
//	}
//
//	func (m *Owner) Username() string {
//		return m.UserModel.Username
//	}
func (m *Owner) IsOwner() bool {
	return true
}
