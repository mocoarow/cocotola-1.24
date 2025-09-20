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
	mbuserusecase "github.com/mocoarow/cocotola-1.24/moonbeam/user/usecase"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
)

func initApp2(ctx context.Context, systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.UserID, error) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp2"))

	sysOwner, err := findSystemOwnerByOrganizationName(ctx, newSystemAdmin(ctx), mbNonTxManager, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("findSystemOwnerByOrganizationName: %w", err)
	}

	guestLoginID := libdomain.NewGuestLoginID(organizationName)
	guestUserName := libdomain.NewGuestUserName(organizationName)
	// 1. check whether the guest user already exists
	{
		guest, err := findUserByLoginID(ctx, sysOwner, mbNonTxManager, guestLoginID)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("guest already exists. id: %d", guest.GetUserID().Int()))
			return guest.GetUserID(), nil
		} else if !errors.Is(err, mbuserservice.ErrUserNotFound) {
			return nil, mbliberrors.Errorf("find user by login id(%s): %w", guestLoginID, err)
		}
	}

	// 2. find public default space
	publicDefaultSpace, err := findPublicSpaceByKey(ctx, sysOwner, mbNonTxManager, mbuserservice.PublicDefaultSpaceKey)
	if err != nil {
		return nil, mbliberrors.Errorf("find public default space by key(%s): %w", mbuserservice.PublicDefaultSpaceKey, err)
	}

	// 3. add guest user
	guestID := addGuestUser(ctx, mbTxManager, mbNonTxManager, sysOwner, guestLoginID, guestUserName, publicDefaultSpace.SpaceID)

	return guestID, nil
}

func addGuestUser(ctx context.Context, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, systemOwner mbuserservice.SystemOwnerInterface, guestLoginID, guestUserName string, spaceID *mbuserdomain.SpaceID) *mbuserdomain.UserID {
	allowEffect := mbuserservice.RBACAllowEffect
	spaceObject := spaceID.GetRBACObject()

	aoeList := []mbuserusecase.ActionObjectEffect{
		// guest can list decks in the "public" space
		{Action: librbac.ListDecksAction, Object: spaceObject, Effect: allowEffect},
		// guest cat read all decks in the "public" space
		{Action: librbac.ReadDeckAction, Object: spaceObject, Effect: allowEffect},
	}
	fn := func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.UserID, error) {
		addGuestCommand := mbuserusecase.NewAddGuestCommand(mbTxManager, mbNonTxManager)
		addUserParam, err := mbuserservice.NewUserAddParameter(guestLoginID, guestUserName, "", "", "", "", "")
		if err != nil {
			libdomain.CheckError(err)
		}

		guestID, err := addGuestCommand.Execute(ctx, systemOwner, addUserParam, aoeList)
		if err != nil {
			return nil, mbliberrors.Errorf("AddGuestCommand.Execute: %w", err)
		}
		return guestID, nil
	}

	guestID, err := mblibservice.Do1(ctx, mbTxManager, fn)
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
