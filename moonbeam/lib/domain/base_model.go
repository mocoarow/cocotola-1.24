package domain

import (
	"time"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type BaseModel struct {
	Version   int `validate:"required,gte=1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int `validate:"gte=0"`
	UpdatedBy int `validate:"gte=0"`
}

func NewBaseModel(version int, createdAt, updatedAt time.Time, createdBy, updatedBy int) (*BaseModel, error) {
	m := &BaseModel{
		Version:   version,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}

	if err := Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate base model: %w", err)
	}

	return m, nil
}
