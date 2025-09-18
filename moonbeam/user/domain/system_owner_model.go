package domain

import (
	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type SystemOwnerModel struct {
	*OwnerModel
}

func NewSystemOwnerModel(user *OwnerModel) (*SystemOwnerModel, error) {
	m := &SystemOwnerModel{
		OwnerModel: user,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate system owner model: %w", err)
	}

	return m, nil
}
