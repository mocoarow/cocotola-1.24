package domain

import (
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type ProfileModel struct {
	PrivateSpaceID *SpaceID
}

func NewProfileModel(privateSpaceID *SpaceID) (*ProfileModel, error) {
	m := &ProfileModel{
		PrivateSpaceID: privateSpaceID,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate profile model: %w", err)
	}

	return m, nil
}
