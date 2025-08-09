package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

// const SystemOwnerID = 2

// type SystemOwnerModel interface {
// 	OwnerModel
// 	IsSystemOwnerModel() bool
// }

type SystemOwnerModel struct {
	*OwnerModel
	// AppUserID AppUserID
}

func NewSystemOwnerModel(appUser *OwnerModel) (*SystemOwnerModel, error) {
	m := &SystemOwnerModel{
		OwnerModel: appUser,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

// func (m *systemOwnerModel) IsSystemOwnerModel() bool {
// 	return true
// }
