package domain

import (
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type SpaceID struct {
	Value int `validate:"required,gte=1"`
}

func NewSpaceID(value int) (*SpaceID, error) {
	return &SpaceID{
		Value: value,
	}, nil
}

func (v *SpaceID) Int() int {
	return v.Value
}
func (v *SpaceID) IsSpaceID() bool {
	return true
}
func (v *SpaceID) GetRBACObject() mbuserdomain.RBACObject {
	return mbuserdomain.NewRBACObject("space:" + fmt.Sprint(v.Value))
}

type SpaceModel struct {
	*mblibdomain.BaseModel
	SpaceID        *SpaceID                     `validate:"required"`
	OrganizationID *mbuserdomain.OrganizationID `validate:"required"`
	OwnerID        *mbuserdomain.AppUserID      `validate:"required"`
	Key            string                       `validate:"required"`
	Name           string                       `validate:"required"`
}

func NewSpaceModel(baseModel *mblibdomain.BaseModel, spaceID *SpaceID, organizationID *mbuserdomain.OrganizationID, owernID *mbuserdomain.AppUserID, key, name string) (*SpaceModel, error) {
	m := &SpaceModel{
		BaseModel:      baseModel,
		SpaceID:        spaceID,
		OrganizationID: organizationID,
		OwnerID:        owernID,
		Key:            key,
		Name:           name,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate space model: %w", err)
	}

	return m, nil
}

func (m *SpaceModel) IsPrivate() bool {
	return m.Key == "private"
}
