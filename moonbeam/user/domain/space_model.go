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

type SpaceModel struct {
	*libdomain.BaseModel
	SpaceID        *SpaceID        `validate:"required"`
	OrganizationID *OrganizationID `validate:"required"`
	OwnerID        *AppUserID      `validate:"required"`
	Key            string          `validate:"required"`
	Name           string          `validate:"required"`
}

func NewSpaceModel(baseModel *libdomain.BaseModel, spaceID *SpaceID, organizationID *OrganizationID, owernID *AppUserID, key, name string) (*SpaceModel, error) {
	m := &SpaceModel{
		BaseModel:      baseModel,
		SpaceID:        spaceID,
		OrganizationID: organizationID,
		OwnerID:        owernID,
		Key:            key,
		Name:           name,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("validate space model: %w", err)
	}

	return m, nil
}

func (m *SpaceModel) IsPrivate() bool {
	return m.Key == "private"
}
