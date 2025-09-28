package service

import (
	"context"
	"errors"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrFolderAlreadyExists = errors.New("folder already exists")
var ErrFolderNotFound = errors.New("folder not found")

type AddFolderParameter struct {
	SpaceID  *mbuserdomain.SpaceID `validate:"required"`
	FolderID *domain.FolderID      `validate:"required"`
	Name     string                `validate:"required"`
}

func NewAddFolderParameter(spaceID *mbuserdomain.SpaceID, folderID *domain.FolderID, name string) (*AddFolderParameter, error) {
	m := &AddFolderParameter{
		SpaceID:  spaceID,
		FolderID: folderID,
		Name:     name,
	}
	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, errors.New("validate add folder parameter: " + err.Error())
	}
	return m, nil
}

type FolderRepository interface {
	AddFolder(ctx context.Context, operator mbuserdomain.UserInterface, param *AddFolderParameter) (*domain.FolderID, error)
	RetrieveRooFolderBySpaceID(ctx context.Context, operator mbuserdomain.UserInterface, spaceID *mbuserdomain.SpaceID) (*domain.Folder, error)
}
