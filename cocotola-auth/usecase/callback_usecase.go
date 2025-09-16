package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type Callback struct {
	systemToken                libdomain.SystemToken
	txManager                  service.TransactionManager
	nonTxManager               service.TransactionManager
	cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient
	logger                     *slog.Logger
}

func NewCallback(systemToken libdomain.SystemToken, txManager, nonTxManager service.TransactionManager, cocotolaCoreCallbackClient service.CocotolaCoreCallbackClient) *Callback {
	return &Callback{
		systemToken:                systemToken,
		txManager:                  txManager,
		nonTxManager:               nonTxManager,
		cocotolaCoreCallbackClient: cocotolaCoreCallbackClient,
		logger:                     slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

func (u *Callback) OnAddUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error {
	u.logger.InfoContext(ctx, "OnAddUser", slog.Int("user_id", userID.Int()))

	fn := func(rf service.RepositoryFactory) error {
		action, err := service.NewSystemOwnerAction(ctx, u.systemToken, rf,
			service.WithOrganizationByID(organizationID),
		)
		if err != nil {
			return mbliberrors.Errorf("NewSystemOwnerAction: %w", err)
		}

		// Create personal space for the new user
		mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		}

		user, err := action.SystemOwner.FindUserByID(ctx, userID)
		if err != nil {
			return mbliberrors.Errorf("FindUserByID: %w", err)
		}

		spaceManager, err := mbrf.NewSpaceManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewSpaceManager: %w", err)
		}
		param := mbuserservice.AddPersonalSpaceParameter{
			UserID:  userID,
			KeyName: libdomain.NewPersonalSpaceKey(user.GetUserID().Int()),
			Name:    libdomain.NewPersonalSpaceName(user.LoginID),
		}
		spaceID, err := spaceManager.AddPersonalSpace(ctx, action.SystemOwner, &param)
		if err != nil {
			return mbliberrors.Errorf("AddSpace: %w", err)
		}

		authorizationManager, err := mbrf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewAuthorizationManager: %w", err)
		}

		subject := userID.GetRBACSubject()
		actions := []mbuserdomain.RBACAction{
			librbac.CreateDeckAction,
			librbac.ListDecksAction,
		}
		object := spaceID.GetRBACObject()
		effect := mbuserservice.RBACAllowEffect
		for _, a := range actions {
			if err := authorizationManager.AddPolicyToUser(ctx, action.SystemOwner, subject, a, object, effect); err != nil {
				return mbliberrors.Errorf("add policy to user. space(%d), action(%s): %w", spaceID.Int(), a, err)
			}
		}

		if err := u.cocotolaCoreCallbackClient.OnAddUserSpace(ctx, organizationID, userID, spaceID); err != nil {
			return mbliberrors.Errorf("cocotolaCoreCallbackClient.OnAddUserSpace: %w", err)
		}

		return nil
	}

	if err := mblibservice.Do0(ctx, u.nonTxManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
