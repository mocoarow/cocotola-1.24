package service

import (
	"context"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type UserInfo struct {
	// UserID        int
	LoginID          string
	Username         string
	OrganizationID   int
	OrganizationName string
}

type UserInterface interface {
	// UserID() *mbuserdomain.UserID
	OrganizationID() *mbuserdomain.OrganizationID
	LoginID() string
	Username() string
	// GetUserGroups() []domain.UserGroupModel
}

type OrganizationInterface interface {
	OrganizationID() *mbuserdomain.OrganizationID
	Name() string
}

type AuthTokenManager interface {
	SignInWithIDToken(ctx context.Context, idToken string) (*domain.AuthTokenSet, error)
	GetUserInfo(ctx context.Context, tokenString string) (*UserInfo, error)

	CreateTokenSet(ctx context.Context, user UserInterface, organizationUsecase OrganizationInterface) (*domain.AuthTokenSet, error)
	RefreshToken(ctx context.Context, accessToken string) (string, error)
}
