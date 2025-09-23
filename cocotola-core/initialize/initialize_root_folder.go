package initialize

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func initRootFolder(ctx context.Context, txManager service.TransactionManager, organizationID *mbuserdomain.OrganizationID, sysOwnerID *mbuserdomain.UserID, publicDefaultSpaceID *mbuserdomain.SpaceID) (*domain.FolderID, error) {
	ctx, span := tracer.Start(ctx, "initRootFolder")
	defer span.End()

	operator := &operator{
		organizationID: organizationID,
		userID:         sysOwnerID,
	}

	fn := func(rf service.RepositoryFactory) (*domain.FolderID, error) {
		foldeRepo := rf.NewFolderRepository(ctx)
		folderID, err := foldeRepo.AddFolder(ctx, operator, &service.AddFolderParameter{
			SpaceID:  publicDefaultSpaceID,
			FolderID: domain.EmptyFolderID,
			Name:     "Root",
		})
		if err != nil {
			return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		return folderID, nil
	}

	folderID, err := mblibservice.Do1(ctx, txManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return folderID, nil
}
