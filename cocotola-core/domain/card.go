package domain

import (
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CardID struct {
	Value int `validate:"required,gte=1"`
}

func NewCardID(value int) (*CardID, error) {
	return &CardID{
		Value: value,
	}, nil
}

func (v *CardID) Int() int {
	return v.Value
}
func (v *CardID) IsCardID() bool {
	return true
}
func (v *CardID) GetRBACObject() mbuserdomain.RBACObject {
	return mbuserdomain.NewRBACObject("card:" + fmt.Sprint(v.Value))
}

type Card struct {
	*mblibdomain.BaseModel `validate:"required"`
	CardID                 *CardID                      `validate:"required"`
	OrganizationID         *mbuserdomain.OrganizationID `validate:"required"`
	DeckID                 *DeckID                      `validate:"required"`
	TemplateID             *TemplateID                  `validate:"required"`
	Content                string                       `validate:"required"`
	OwnerID                *mbuserdomain.UserID         `validate:"required"`
}

func NewCard(baseModel *mblibdomain.BaseModel, cardID *CardID, organizationID *mbuserdomain.OrganizationID, deckID *DeckID, templateID *TemplateID, content string, owernID *mbuserdomain.UserID) (*Card, error) {
	m := &Card{
		BaseModel:      baseModel,
		CardID:         cardID,
		OrganizationID: organizationID,
		DeckID:         deckID,
		TemplateID:     templateID,
		Content:        content,
		OwnerID:        owernID,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate card model: %w", err)
	}

	return m, nil
}
