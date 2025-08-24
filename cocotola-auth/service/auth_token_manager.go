package service

import (
	"context"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

type AppUserInfo struct {
	// AppUserID        int
	LoginID          string
	Username         string
	OrganizationID   int
	OrganizationName string
}

type AppUserInterface interface {
	// AppUserID() *mbuserdomain.AppUserID
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
	GetUserInfo(ctx context.Context, tokenString string) (*AppUserInfo, error)

	CreateTokenSet(ctx context.Context, appUser AppUserInterface, organizationUsecase OrganizationInterface) (*domain.AuthTokenSet, error)
	RefreshToken(ctx context.Context, accessToken string) (string, error)
}
