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
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.Organization, error) {
		orgRepo := rf.NewOrganizationRepository(ctx)
		org, err := orgRepo.FindOrganizationByName(ctx, systemAdmin, organizationName)
		if err != nil {
			if errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
				return nil, mbliberrors.Errorf("organization not found(%s): %w", organizationName, err)
			}
			return nil, mbliberrors.Errorf("find organization by name(%s): %w", organizationName, err)
		}
		return org, nil
	}
	orgModel, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return orgModel, nil
}

func findSystemOwnerByOrganizationID(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationID *mbuserdomain.OrganizationID) (*mbuserdomain.SystemOwner, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) {
		userRepo := rf.NewUserRepository(ctx)
		sysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, systemAdmin, organizationID)
		if err != nil {
			return nil, mbliberrors.Errorf("find system owner by organization id(%d): %w", organizationID.Int(), err)
		}

		return sysOwner, nil
	}
	sysOwner, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err
	}
	return sysOwner, nil
}

func findSystemOwnerByOrganizationName(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.SystemOwner, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) {
		userRepo := rf.NewUserRepository(ctx)
		sysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
		if err != nil {
			return nil, mbliberrors.Errorf("find system owner by organization name(%s): %w", organizationName, err)
		}

		return sysOwner, nil
	}
	sysOwner, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err
	}
	return sysOwner, nil
}

func findUserByLoginID(ctx context.Context, systemOwner mbuserdomain.SystemOwnerInterface, mbNonTxManager mbuserservice.TransactionManager, loginID string) (*mbuserdomain.User, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.User, error) {
		userRepo := rf.NewUserRepository(ctx)
		user, err := userRepo.FindUserByLoginID(ctx, systemOwner, loginID)
		if err != nil {
			return nil, mbliberrors.Errorf("find user by login id(%s): %w", loginID, err)
		}

		return user, nil
	}
	userModel, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func findPublicSpaceByKey(ctx context.Context, systemOwner mbuserdomain.SystemOwnerInterface, mbNonTxManager mbuserservice.TransactionManager, key string) (*mbuserdomain.SpaceModel, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SpaceModel, error) {
		spaceRepo := rf.NewSpaceRepository(ctx)
		publicDefaultSpace, err := spaceRepo.FindPublicSpaceByKey(ctx, systemOwner, mbuserservice.PublicDefaultSpaceKey)
		if err != nil {
			return nil, mbliberrors.Errorf("find public default space by key(%s): %w", mbuserservice.PublicDefaultSpaceKey, err)
		}

		return publicDefaultSpace, nil
	}
	publicDefaultSpaceModel, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return publicDefaultSpaceModel, nil
}
