package domain

import (
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type DeckID struct {
	Value int `validate:"required,gte=1"`
}

func NewDeckID(value int) (*DeckID, error) {
	return &DeckID{
		Value: value,
	}, nil
}

func (v *DeckID) Int() int {
	return v.Value
}
func (v *DeckID) IsDeckID() bool {
	return true
}
func (v *DeckID) GetRBACObject() mbuserdomain.RBACObject {
	return mbuserdomain.NewRBACObject("deck:" + fmt.Sprint(v.Value))
}

type DeckModel struct {
	*mblibdomain.BaseModel
	OrganizationID *mbuserdomain.OrganizationID
	OwnerID        *mbuserdomain.AppUserID
	Name           string `validate:"required"`
	TemplateID     int    `validate:"required,gte=1"`
	Lang2          string `validate:"required"`
	Description    string
}

func NewDeckModel(baseModel *mblibdomain.BaseModel, organizationID *mbuserdomain.OrganizationID, owernID *mbuserdomain.AppUserID, name string, templateID int, lang2 string, description string) (*DeckModel, error) {
	m := &DeckModel{
		BaseModel:      baseModel,
		OrganizationID: organizationID,
		OwnerID:        owernID,
		Name:           name,
		TemplateID:     templateID,
		Lang2:          lang2,
		Description:    description,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate deck model: %w", err)
	}

	return m, nil
}
