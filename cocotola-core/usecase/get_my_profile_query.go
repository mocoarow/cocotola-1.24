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

type GetMyProfileQuery struct {
	nonTxManager       service.TransactionManager
	cocotolaAuthClient libapi.CocotolaAuthClient
	logger             *slog.Logger
}

func NewGetMyProfileQuery(nonTxManager service.TransactionManager, cocotolaAuthClient libapi.CocotolaAuthClient) *GetMyProfileQuery {
	return &GetMyProfileQuery{
		nonTxManager:       nonTxManager,
		cocotolaAuthClient: cocotolaAuthClient,
		logger:             slog.Default().With(slog.String(mbliblog.LoggerNameKey, "GetMyProfileQuery")),
	}
}

func (u *GetMyProfileQuery) Execute(ctx context.Context, operator mbuserdomain.UserInterface, bearerToken string) (*domain.ProfileModel, error) {
	getMyProfileResp, err := u.cocotolaAuthClient.GetMyProfile(ctx, bearerToken)
	if err != nil {
		return nil, mbliberrors.Errorf("cocotolaAuthClient.GetMyProfile: %w", err)
	}
	spaceID, err := mbuserdomain.NewSpaceID(getMyProfileResp.PrivateSpaceID)
	if err != nil {
		return nil, mbliberrors.Errorf("mbuserdomain.NewSpaceID: %w", err)
	}

	fn := func(rf service.RepositoryFactory) (*domain.FolderID, error) {
		folderRepo := rf.NewFolderRepository(ctx)
		folder, err := folderRepo.RetrieveRooFolderBySpaceID(ctx, operator, spaceID)
		if err != nil {
			return nil, mbliberrors.Errorf("RetrieveRooFolderBySpaceID: %w", err)
		}
		return folder.FolderID, nil
	}
	rootFolerID, err := mblibservice.Do1(ctx, u.nonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return &domain.ProfileModel{
		PrivateSpaceID: spaceID,
		RootFolderID:   rootFolerID,
	}, nil
}
