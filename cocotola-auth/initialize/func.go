package initialize

import (
	"context"
	"errors"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func findOrganizationByName(ctx context.Context, systemAdmin mbuserservice.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.OrganizationModel, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.OrganizationModel, error) {
		orgRepo := rf.NewOrganizationRepository(ctx)
		org, err := orgRepo.FindOrganizationByName(ctx, systemAdmin, organizationName)
		if err != nil {
			if errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
				return nil, mbliberrors.Errorf("organization not found(%s): %w", organizationName, err)
			}
			return nil, mbliberrors.Errorf("find organization by name(%s): %w", organizationName, err)
		}
		return org.OrganizationModel, nil
	}
	orgModel, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return orgModel, nil
}

func findSystemOwnerByOrganizationName(ctx context.Context, systemAdmin mbuserservice.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.SystemOwnerModel, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwnerModel, error) {
		userRepo := rf.NewUserRepository(ctx)
		sysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
		if err != nil {
			return nil, mbliberrors.Errorf("find system owner by organization name(%s): %w", organizationName, err)
		}

		return sysOwner.SystemOwnerModel, nil
	}
	sysOwner, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err
	}
	return sysOwner, nil
}

func findUserByLoginID(ctx context.Context, systemOwner mbuserservice.SystemOwnerInterface, mbNonTxManager mbuserservice.TransactionManager, loginID string) (*mbuserdomain.UserModel, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.UserModel, error) {
		userRepo := rf.NewUserRepository(ctx)
		user, err := userRepo.FindUserByLoginID(ctx, systemOwner, loginID)
		if err != nil {
			return nil, mbliberrors.Errorf("find user by login id(%s): %w", loginID, err)
		}

		return user.UserModel, nil
	}
	userModel, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func findPublicSpaceByKey(ctx context.Context, systemOwner mbuserservice.SystemOwnerInterface, mbNonTxManager mbuserservice.TransactionManager, key string) (*mbuserdomain.SpaceModel, error) {
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.SpaceModel, error) {
		spaceRepo := rf.NewSpaceRepository(ctx)
		publicDefaultSpace, err := spaceRepo.FindPublicSpaceByKey(ctx, systemOwner, mbuserservice.PublicDefaultSpaceKey)
		if err != nil {
			return nil, mbliberrors.Errorf("find public default space by key(%s): %w", mbuserservice.PublicDefaultSpaceKey, err)
		}

		return publicDefaultSpace.SpaceModel, nil
	}
	publicDefaultSpaceModel, err := mblibservice.Do1(ctx, mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return publicDefaultSpaceModel, nil
}
