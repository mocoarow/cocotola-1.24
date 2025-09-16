package initialize

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func initApp2(ctx context.Context, systemToken libdomain.SystemToken, _, nonTxManager service.TransactionManager, organizationName string) (*mbuserdomain.UserID, error) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp2"))

	fn := func(rf service.RepositoryFactory) (*mbuserdomain.UserID, error) {
		systemOwnerAction := newSystemOwnerAction(ctx, systemToken, rf, organizationName)

		guestLoginID := libdomain.NewGuestLoginID(organizationName)
		guestUserName := libdomain.NewGuestUserName(organizationName)

		// 1. check whether the guest user already exists
		guest, err := systemOwnerAction.SystemOwner.FindUserByLoginID(ctx, guestLoginID)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("guest already exists. id: %d", guest.GetUserID().Int()))
			return guest.GetUserID(), nil
		} else if !errors.Is(err, mbuserservice.ErrUserNotFound) {
			return nil, mbliberrors.Errorf("find app user by login id(%s): %w", guestLoginID, err)
		}

		// 2. add guest user
		guestID := addGuestUser(ctx, guestLoginID, guestUserName, systemOwnerAction)

		mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		}

		// 3. add "guest" user to "public" group
		userGroupRepo := mbrf.NewUserGroupRepository(ctx)

		publicGroup, err := userGroupRepo.FindUserGroupByKey(ctx, systemOwnerAction.SystemOwner, mbuserservice.PublicGroupKey)
		if err != nil {
			return nil, mbliberrors.Errorf("find public group(%s): %w", mbuserservice.PublicGroupKey, err)
		}

		authorizationManager := initAuthorizationManager(ctx, mbrf)

		if err := authorizationManager.AddUserToGroup(ctx, systemOwnerAction.SystemOwner, guestID, publicGroup.UserGroupID); err != nil {
			return nil, mbliberrors.Errorf("AddUserToGroup: %w", err)
		}

		// space
		spaceRepo := mbrf.NewSpaceRepository(ctx)
		publicDefaultSpace, err := spaceRepo.FindPublicSpaceByKey(ctx, mbuserservice.PublicDefaultSpaceKey)
		if err != nil {
			return nil, mbliberrors.Errorf("find public space(%s): %w", mbuserservice.PublicDefaultSpaceKey, err)
		}

		subject := guestID.GetRBACSubject()
		allowEffect := mbuserservice.RBACAllowEffect
		spaceObject := publicDefaultSpace.SpaceID.GetRBACObject()

		// guest can list decks in the "public" space
		action := librbac.ListDecksAction
		if err := authorizationManager.AddPolicyToUserBySystemOwner(ctx, systemOwnerAction.SystemOwner, subject, action, spaceObject, allowEffect); err != nil {
			return nil, mbliberrors.Errorf("AddPolicyToUserBySystemOwner: %w", err)
		}

		// guest cat read all decks in the "public" space
		action = librbac.ReadDeckAction
		if err := authorizationManager.AddPolicyToUserBySystemOwner(ctx, systemOwnerAction.SystemOwner, subject, action, spaceObject, allowEffect); err != nil {
			return nil, mbliberrors.Errorf("AddPolicyToUserBySystemOwner: %w", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("guest was created. id: %d", guestID.Int()))

		return guestID, nil
	}

	guestID, err := mblibservice.Do1(ctx, nonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return guestID, nil
}

func addGuestUser(ctx context.Context, guestLoginID, guestUserName string, systemOwnerAction *service.SystemOwnerAction) *mbuserdomain.UserID {
	appUserAddParam, err := mbuserservice.NewUserAddParameter(guestLoginID, guestUserName, "", "", "", "", "")
	if err != nil {
		libdomain.CheckError(err)
	}

	guestID, err := systemOwnerAction.SystemOwner.AddUser(ctx, appUserAddParam)
	if err != nil {
		libdomain.CheckError(err)
	}

	return guestID
}

func initAuthorizationManager(ctx context.Context, mbrf mbuserservice.RepositoryFactory) mbuserservice.AuthorizationManager {
	authorizationManager, err := mbrf.NewAuthorizationManager(ctx)
	if err != nil {
		libdomain.CheckError(err)
	}
	return authorizationManager
}
