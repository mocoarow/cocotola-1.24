package domain

import (
	"strconv"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type SpaceID struct {
	Value int `validate:"required,gte=1"`
}

func NewSpaceID(value int) (*SpaceID, error) {
	m := &SpaceID{
		Value: value,
	}
	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate space id(%d): %w", value, err)
	}
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
func (v *SpaceID) GetRBACObject() RBACObject {
	return NewRBACObject("space:" + strconv.Itoa(v.Value))
}

type SpaceIDs []*SpaceID

func (v *SpaceIDs) IDs() []int {
	ids := make([]int, len(*v))
	for i, id := range *v {
		ids[i] = id.Int()
	}

	return ids
}

type Space struct {
	*libdomain.BaseModel
	SpaceID        *SpaceID        `validate:"required"`
	OrganizationID *OrganizationID `validate:"required"`
	OwnerID        *UserID         `validate:"required"`
	KeyName        string          `validate:"required"`
	Name           string          `validate:"required"`
	SpaceType      string          `validate:"required,oneof=personal private public"`
}

func NewSpace(baseModel *libdomain.BaseModel, spaceID *SpaceID, organizationID *OrganizationID, owernID *UserID, keyName, name string, spaceType string) (*Space, error) {
	m := &Space{
		BaseModel:      baseModel,
		SpaceID:        spaceID,
		OrganizationID: organizationID,
		OwnerID:        owernID,
		KeyName:        keyName,
		Name:           name,
		SpaceType:      spaceType,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate space model: %w", err)
	}

	return m, nil
}
func (m *Space) IsPrivate() bool {
	return m.SpaceType == "private"
}
