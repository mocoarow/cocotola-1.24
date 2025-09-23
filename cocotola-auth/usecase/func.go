package usecase

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func findUserByID(ctx context.Context, mbrf mbuserservice.RepositoryFactory, operator mbuserdomain.UserInterface, userID *mbuserdomain.UserID) (*mbuserdomain.User, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	user, err := userRepo.FindUserByID(ctx, operator, userID)
	if err != nil {
		return nil, mbliberrors.Errorf("find user by id(%d): %w", userID.Int(), err)
	}
	return user, nil
}

func findUserbyLoginID(ctx context.Context, mbrf mbuserservice.RepositoryFactory, operator mbuserdomain.UserInterface, loginID string) (*mbuserdomain.User, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	user, err := userRepo.FindUserByLoginID(ctx, operator, loginID)
	if err != nil {
		return nil, mbliberrors.Errorf("find user by login id(%s): %w", loginID, err)
	}
	return user, nil
}

func getOrganization(ctx context.Context, mbrf mbuserservice.RepositoryFactory, operator mbuserdomain.UserInterface) (*mbuserdomain.Organization, error) {
	orgRepo := mbrf.NewOrganizationRepository(ctx)
	org, err := orgRepo.GetOrganization(ctx, operator)
	if err != nil {
		return nil, mbliberrors.Errorf("get organization: %w", err)
	}

	return org, nil
}

func createTokenSet(ctx context.Context, authTokenManager service.AuthTokenManager, user *mbuserdomain.User, organization *mbuserdomain.Organization) (*domain.AuthTokenSet, error) {
	targetUser := &usecaseUser{
		userID:         user.UserID,
		organizationID: user.OrganizationID,
		loginID:        user.LoginID,
		username:       user.Username,
	}
	targetOorganization := &usecaseOrganization{
		organizationID: organization.OrganizationID,
		name:           organization.Name,
	}

	tokenSet, err := authTokenManager.CreateTokenSet(ctx, targetUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return tokenSet, nil
}
