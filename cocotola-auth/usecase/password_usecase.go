package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type PasswordUsecae struct {
	systemToken      libdomain.SystemToken
	txManager        mbuserservice.TransactionManager
	nonTxManager     mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
}

func NewPassword(systemToken libdomain.SystemToken, txManager, nonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) *PasswordUsecae {
	return &PasswordUsecae{
		systemToken:      systemToken,
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
	}
}

func (u *PasswordUsecae) Authenticate(ctx context.Context, loginID, password, organizationName string) (*domain.AuthTokenSet, error) {
	var tokenSet *domain.AuthTokenSet

	if strings.Contains(loginID, "guest@@") {
		return nil, fmt.Errorf("guest cannot authenticate with password")
	}

	targetOorganization, targetUser, err := mblibservice.Do2(ctx, u.txManager, func(rf mbuserservice.RepositoryFactory) (*organization, *user, error) {
		action, err := service.NewSystemOwnerAction(ctx, u.systemToken, rf,
			service.WithOrganizationByName(organizationName),
		)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("new organization action: %w", err)
		} else if action.Organization == nil {
			return nil, nil, mbliberrors.Errorf("organization is nil")
		}

		verified, err := action.SystemOwner.VerifyPassword(ctx, loginID, password)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("action.userRepo.VerifyPassword: %w", err)
		} else if !verified {
			return nil, nil, domain.ErrUnauthenticated
		}

		tmpUser, err := action.SystemOwner.FindUserByLoginID(ctx, loginID)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("find user by login id: %w", err)
		}

		targetOorganization := &organization{
			organizationID: action.Organization.OrganizationID,
			name:           action.Organization.Name,
		}

		targetUser := &user{
			userID:         tmpUser.UserID,
			organizationID: tmpUser.OrganizationID,
			loginID:        tmpUser.LoginID,
			username:       tmpUser.Username,
		}

		return targetOorganization, targetUser, nil
	})
	if err != nil {
		if errors.Is(err, mbuserservice.ErrUserNotFound) {
			return nil, mbliberrors.Errorf("user not found: %w", domain.ErrUnauthenticated)
		}

		return nil, mbliberrors.Errorf("authenticate: %w", err)
	}

	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}
	tokenSet = tokenSetTmp

	return tokenSet, nil
}
