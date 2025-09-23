package initialize

import (
	"context"
	"errors"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func findOrganizationByName(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.Organization, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.Organization, error) {
		orgRepo := mbrf.NewOrganizationRepository(ctx)
		org, err := orgRepo.FindOrganizationByName(ctx, systemAdmin, organizationName)
		if err != nil {
			if errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
				return nil, mbliberrors.Errorf("organization not found(%s): %w", organizationName, err)
			}
			return nil, mbliberrors.Errorf("find organization by name(%s): %w", organizationName, err)
		}
		return org, nil
	}
	org, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return org, nil
}

func findSystemOwnerByOrganizationID(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationID *mbuserdomain.OrganizationID) (*mbuserdomain.SystemOwner, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) {
		userRepo := mbrf.NewUserRepository(ctx)
		sysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, systemAdmin, organizationID)
		if err != nil {
			return nil, mbliberrors.Errorf("find system owner by organization id(%d): %w", organizationID.Int(), err)
		}

		return sysOwner, nil
	}
	sysOwner, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return sysOwner, nil
}

func findSystemOwnerByOrganizationName(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.SystemOwner, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) {
		userRepo := mbrf.NewUserRepository(ctx)
		sysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
		if err != nil {
			return nil, mbliberrors.Errorf("find system owner by organization name(%s): %w", organizationName, err)
		}

		return sysOwner, nil
	}
	sysOwner, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return sysOwner, nil
}

func findUserByLoginID(ctx context.Context, systemOwner mbuserdomain.SystemOwnerInterface, mbNonTxManager mbuserservice.TransactionManager, loginID string) (*mbuserdomain.User, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.User, error) {
		userRepo := mbrf.NewUserRepository(ctx)
		user, err := userRepo.FindUserByLoginID(ctx, systemOwner, loginID)
		if err != nil {
			return nil, mbliberrors.Errorf("find user by login id(%s): %w", loginID, err)
		}

		return user, nil
	}
	user, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return user, nil
}

func findPublicSpaceByKey(ctx context.Context, systemOwner mbuserdomain.SystemOwnerInterface, mbNonTxManager mbuserservice.TransactionManager, key string) (*mbuserdomain.Space, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.Space, error) {
		spaceRepo := mbrf.NewSpaceRepository(ctx)
		publicDefaultSpace, err := spaceRepo.FindPublicSpaceByKey(ctx, systemOwner, key)
		if err != nil {
			return nil, mbliberrors.Errorf("find public default space by key(%s): %w", key, err)
		}

		return publicDefaultSpace, nil
	}
	publicDefaultSpace, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return publicDefaultSpace, nil
}
