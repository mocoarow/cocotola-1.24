package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type Callback struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   libapi.CocotolaRBACClient
	logger       *slog.Logger
}

func NewCallback(txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) *Callback {
	return &Callback{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

// func (u *Callback) OnAddUser(ctx context.Context, _ *mbuserdomain.OrganizationID, appUserID *mbuserdomain.UserID) error {
// 	u.logger.InfoContext(ctx, "OnAddUser", slog.Int("app_user_id", appUserID.Int()))

// 	fn := func(_ service.RepositoryFactory) error {
// 		spaceRepo, err := rf.NewSpaceRepository(ctx)
// 		if err != nil {
// 			return mbliberrors.Errorf("NewSpaceRepository: %w", err)
// 		}
// 		pairOfUserAndSpaceRep, err := rf.NewPairOfUserAndSpaceRepository(ctx)
// 		if err != nil {
// 			return mbliberrors.Errorf("NewPairOfUserAndSpaceRepository: %w", err)
// 		}

// 		operator := Operator{
// 			organizationID: organizationID,
// 			appUserID:      appUserID,
// 		}

// 		param := service.SpaceAddParameter{
// 			Key:      "private",
// 			Name:     "Private",
// 			IsPublic: false,
// 		}

// 		spaceID, err := spaceRepo.AddSpace(ctx, &operator, &param)
// 		if err != nil {
// 			return mbliberrors.Errorf("AddSpace: %w", err)
// 		}

// 		object := spaceID.GetRBACObject()

// 		if err := u.rbacClient.AddPolicyToUser(ctx, &libapi.AddPolicyToUserParameter{
// 			OrganizationID: operator.OrganizationID().Int(),
// 			UserID:      operator.UserID().Int(),
// 			ListOfActionObjectEffect: []libapi.ActionObjectEffect{
// 				{
// 					Action: librbac.CreateDeckAction.Action(),
// 					Object: object.Object(),
// 					Effect: mbuserservice.RBACAllowEffect.Effect(),
// 				},
// 			},
// 		}); err != nil {
// 			return mbliberrors.Errorf("add policy to user. space(%d): %w", spaceID.Int(), err)
// 		}

// 		if err := pairOfUserAndSpaceRep.AddPairOfUserAndSpace(ctx, &operator, appUserID, spaceID); err != nil {
// 			return mbliberrors.Errorf("AddPairOfUserAndSpace: %w", err)
// 		}

// 		u.logger.InfoContext(ctx, "OnAddUser: AddSpace", slog.Int("space_id", spaceID.Int()))
// 		return nil
// 	}

// 	if err := mblibservice.Do0(ctx, u.nonTxManager, fn); err != nil {
// 		return err //nolint:wrapcheck
// 	}

// 	return nil
// }

func (u *Callback) OnAddUserSpace(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.UserID, spaceID *mbuserdomain.SpaceID) error {
	u.logger.InfoContext(ctx, "OnAddUser", slog.Int("app_user_id", appUserID.Int()))

	operator := Operator{
		organizationID: organizationID,
		appUserID:      appUserID,
	}

	fn := func(rf service.RepositoryFactory) error {
		folderRepo, err := rf.NewFolderRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewSpaceRepository: %w", err)
		}

		param := service.AddFolderParameter{
			SpaceID:  spaceID,
			FolderID: domain.EmptyFolderID,
			Name:     "Root",
		}

		spaceID, err := folderRepo.AddFolder(ctx, &operator, &param)
		if err != nil {
			return mbliberrors.Errorf("AddFolder: %w", err)
		}

		u.logger.InfoContext(ctx, "OnAddUser: AddSpace", slog.Int("space_id", spaceID.Int()))
		return nil
	}

	if err := mblibservice.Do0(ctx, u.nonTxManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
