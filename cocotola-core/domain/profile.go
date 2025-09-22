package domain

import (
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type ProfileModel struct {
	PrivateSpaceID *mbuserdomain.SpaceID `validate:"required"`
	RootFolderID   *FolderID             `validate:"required"`
}

func NewProfileModel(privateSpaceID *mbuserdomain.SpaceID, rootFolderID *FolderID) (*ProfileModel, error) {
	m := &ProfileModel{
		PrivateSpaceID: privateSpaceID,
		RootFolderID:   rootFolderID,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate profile model: %w", err)
	}

	return m, nil
}
