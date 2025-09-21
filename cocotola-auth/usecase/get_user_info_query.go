package usecase

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type GetUserInfoQuery struct {
	mbNonTxManager   mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewGetUserInfoQuery(mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *GetUserInfoQuery {
	return &GetUserInfoQuery{
		mbNonTxManager:   mbNonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *GetUserInfoQuery) Execute(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, bearerToken string) (*mbuserdomain.User, error) {
	// TODO: Check whether the token is registered in the Database
	userInfo, err := u.authTokenManager.GetUserInfo(ctx, bearerToken)
	if err != nil {
		return nil, mbliberrors.Errorf("GetUserInfo: %w", err)
	}
	sysOwner, err := u.findSystemOwnerByOrganizationName(ctx, systemAdmin, userInfo.OrganizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("findSystemOwnerByOrganizationName: %w", err)
	}
	user, err := u.findUserbyLoginID(ctx, sysOwner, userInfo.LoginID)
	if err != nil {
		return nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	return user, nil
}

func (u *GetUserInfoQuery) findSystemOwnerByOrganizationName(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationName string) (*mbuserdomain.SystemOwner, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) {
		return findSystemOwnerByOrganizationName(ctx, mbrf, operator, organizationName)
	})
}

func (u *GetUserInfoQuery) findUserbyLoginID(ctx context.Context, operator mbuserdomain.UserInterface, loginID string) (*mbuserdomain.User, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.User, error) {
		return findUserbyLoginID(ctx, mbrf, operator, loginID)
	})
}
