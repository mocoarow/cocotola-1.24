package usecase

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func findUserByID(ctx context.Context, mbrf mbuserservice.RepositoryFactory, operator mbuserservice.OperatorInterface, userID *mbuserdomain.UserID) (*mbuserdomain.UserModel, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	user, err := userRepo.FindUserByID(ctx, operator, userID)
	if err != nil {
		return nil, mbliberrors.Errorf("FindUserByLoginID(%d). err: %w", userID.Int(), err)
	}
	return user, nil
}

func findUserbyLoginID(ctx context.Context, mbrf mbuserservice.RepositoryFactory, operator mbuserservice.OperatorInterface, loginID string) (*mbuserdomain.UserModel, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	user, err := userRepo.FindUserByLoginID(ctx, operator, loginID)
	if err != nil {
		return nil, mbliberrors.Errorf("FindUserByLoginID(%s). err: %w", loginID, err)
	}
	return user, nil
}

func findSystemOwnerByOrganizationID(ctx context.Context, mbrf mbuserservice.RepositoryFactory, systemAdmin mbuserservice.SystemAdminInterface, organizationID *mbuserdomain.OrganizationID) (*mbuserdomain.SystemOwnerModel, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	sysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, systemAdmin, organizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("FindSystemOwnerByOrganizationID: %w", err)
	}
	return sysOwner, nil
}

func findSystemOwnerByOrganizationName(ctx context.Context, mbrf mbuserservice.RepositoryFactory, systemAdmin mbuserservice.SystemAdminInterface, organizationName string) (*mbuserdomain.SystemOwnerModel, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	sysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("FindSystemOwnerByOrganizationName: %w", err)
	}
	return sysOwner, nil
}

func getOrganization(ctx context.Context, mbrf mbuserservice.RepositoryFactory, operator mbuserservice.OperatorInterface) (*mbuserdomain.OrganizationModel, error) {
	orgRepo := mbrf.NewOrganizationRepository(ctx)
	org, err := orgRepo.GetOrganization(ctx, operator)
	if err != nil {
		return nil, mbliberrors.Errorf("GetOrganization(). err: %w", err)
	}

	return org.OrganizationModel, nil
}

func createTokenSet(ctx context.Context, authTokenManager service.AuthTokenManager, userModel *mbuserdomain.UserModel, organizationModel *mbuserdomain.OrganizationModel) (*domain.AuthTokenSet, error) {
	targetUser := &user{
		userID:         userModel.UserID,
		organizationID: userModel.OrganizationID,
		loginID:        userModel.LoginID,
		username:       userModel.Username,
	}
	targetOorganization := &organization{
		organizationID: organizationModel.OrganizationID,
		name:           organizationModel.Name,
	}

	tokenSet, err := authTokenManager.CreateTokenSet(ctx, targetUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return tokenSet, nil
}
