package service

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func FindSystemOwnerByOrganizationID(ctx context.Context, mbrf mbuserservice.RepositoryFactory, systemAdmin mbuserdomain.SystemAdminInterface, organizationID *mbuserdomain.OrganizationID) (*mbuserdomain.SystemOwner, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	sysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, systemAdmin, organizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("FindSystemOwnerByOrganizationID: %w", err)
	}
	return sysOwner, nil
}

func FindSystemOwnerByOrganizationName(ctx context.Context, mbrf mbuserservice.RepositoryFactory, systemAdmin mbuserdomain.SystemAdminInterface, organizationName string) (*mbuserdomain.SystemOwner, error) {
	userRepo := mbrf.NewUserRepository(ctx)
	sysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("FindSystemOwnerByOrganizationName: %w", err)
	}
	return sysOwner, nil
}
