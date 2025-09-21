package usecase

import (
	"context"
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
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type AddPersonalSpaceCommand struct {
	mbTxManager                mbuserservice.TransactionManager
	mbNonTxManager             mbuserservice.TransactionManager
	cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient
	logger                     *slog.Logger
}

func NewAddPersonalSpaceCommand(_ context.Context, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient) *AddPersonalSpaceCommand {
	return &AddPersonalSpaceCommand{
		mbTxManager:                mbTxManager,
		mbNonTxManager:             mbNonTxManager,
		cocotolaCoreCallbackClient: cocotolaCoreCallbackClient,
		logger:                     slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-AddPersonalspaceCommand")),
	}
}

func (u *AddPersonalSpaceCommand) Execute(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator); err != nil {
		return mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	if err := u.execute(ctx, operator, organizationID, userID); err != nil {
		return mbliberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(); err != nil {
		return mbliberrors.Errorf("callback: %w", err)
	}

	return nil
}

func (u *AddPersonalSpaceCommand) checkAuthorization(_ context.Context, _ mbuserdomain.SystemAdminInterface) error {
	return nil
}

func (u *AddPersonalSpaceCommand) execute(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error {
	sysOwner, err := u.findSystemOwnerByOrganizationID(ctx, operator, organizationID)
	if err != nil {
		return mbliberrors.Errorf("findSystemOwnerByOrganizationID: %w", err)
	}

	user, err := u.findUserByID(ctx, sysOwner, userID)
	if err != nil {
		return mbliberrors.Errorf("findUserByID: %w", err)
	}

	fn1 := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SpaceID, error) {
		spaceManager, err := mbrf.NewSpaceManager(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewSpaceManager: %w", err)
		}

		param := mbuserservice.AddPersonalSpaceParameter{
			UserID:  userID,
			KeyName: libdomain.NewPersonalSpaceKey(user.GetUserID().Int()),
			Name:    libdomain.NewPersonalSpaceName(user.LoginID),
		}
		spaceID, err := spaceManager.AddPersonalSpace(ctx, sysOwner, &param)
		if err != nil {
			return nil, mbliberrors.Errorf("AddPersonalSpace: %w", err)
		}
		return spaceID, nil
	}
	spaceID, err := mblibservice.Do1(ctx, u.mbTxManager, fn1)
	if err != nil {
		return err //nolint:wrapcheck
	}

	subject := userID.GetRBACSubject()
	object := spaceID.GetRBACObject()
	effect := mbuserservice.RBACAllowEffect
	aoeList := []mbuserusecase.ActionObjectEffect{
		{Action: librbac.CreateDeckAction, Object: object, Effect: effect},
		{Action: librbac.ListDecksAction, Object: object, Effect: effect},
	}
	fn2 := func(mbrf mbuserservice.RepositoryFactory) error {
		authorizationManager, err := mbrf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewAuthorizationManager: %w", err)
		}
		for _, aoe := range aoeList {
			if err := authorizationManager.AddPolicyToUser(ctx, sysOwner, subject, aoe.Action, aoe.Object, aoe.Effect); err != nil {
				return mbliberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
			}
		}
		return nil
	}
	if err := mblibservice.Do0(ctx, u.mbTxManager, fn2); err != nil {
		return err //nolint:wrapcheck
	}

	if err := u.cocotolaCoreCallbackClient.OnAddUserSpace(ctx, organizationID, userID, spaceID); err != nil {
		return mbliberrors.Errorf("cocotolaCoreCallbackClient.OnAddUserSpace: %w", err)
	}

	return nil
}

func (u *AddPersonalSpaceCommand) callback() error {
	return nil
}

func (u *AddPersonalSpaceCommand) findSystemOwnerByOrganizationID(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationID *mbuserdomain.OrganizationID) (*mbuserdomain.SystemOwner, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) { //nolint:wrapcheck
		return findSystemOwnerByOrganizationID(ctx, mbrf, operator, organizationID)
	})
}

func (u *AddPersonalSpaceCommand) findUserByID(ctx context.Context, operator mbuserdomain.UserInterface, userID *mbuserdomain.UserID) (*mbuserdomain.User, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.User, error) { //nolint:wrapcheck
		return findUserByID(ctx, mbrf, operator, userID)
	})
}
