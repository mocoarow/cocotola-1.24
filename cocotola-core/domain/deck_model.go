package domain

import (
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

type DeckModel struct {
	*mblibdomain.BaseModel
	OrganizationID *mbuserdomain.OrganizationID
	Name           string `validate:"required"`
	TemplateID     int    `validate:"required,gte=1"`
	Lang2          string `validate:"required"`
	Description    string
}

func NewDeckModel(baseModel *mblibdomain.BaseModel, organizationID *mbuserdomain.OrganizationID, name string, templateID int, lang2 string, description string) (*DeckModel, error) {
	m := &DeckModel{
		BaseModel:      baseModel,
		OrganizationID: organizationID,
		Name:           name,
		TemplateID:     templateID,
		Lang2:          lang2,
		Description:    description,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}
