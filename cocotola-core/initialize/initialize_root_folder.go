package initialize

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func initRootFolder(ctx context.Context, txManager service.TransactionManager, organizationID *mbuserdomain.OrganizationID, publicDefaultSpaceID *mbuserdomain.SpaceID) (*domain.FolderID, error) {
	operator := &operator{
		organizationID: organizationID,
		appUserID:      mbuserservice.SystemAdminID,
	}

	fn := func(rf service.RepositoryFactory) (*domain.FolderID, error) {
		foldeRepo, err := rf.NewFolderRepository(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewFolderRepository: %w", err)
		}
		folderID, err := foldeRepo.AddFolder(ctx, operator, &service.FolderAddParameter{
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
