package service

import (
	"context"
	"errors"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrFolderAlreadyExists = errors.New("folder already exists")
var ErrFolderNotFound = errors.New("folder not found")

type AddFolderParameter struct {
	SpaceID  *mbuserdomain.SpaceID
	FolderID *domain.FolderID
	Name     string
}

// type FolderUpdateParameter struct {
// 	Name        string
// 	Description string
// }

// type FindFoldersParameter struct {
// 	SpaceIDs mbuserdomain.SpaceIDs
// }

type FolderRepository interface {
	AddFolder(ctx context.Context, operator mbuserdomain.UserInterface, param *AddFolderParameter) (*domain.FolderID, error)
	RetrieveRooFolderBySpaceID(ctx context.Context, operator mbuserdomain.UserInterface, spaceID *mbuserdomain.SpaceID) (*Folder, error)
}
