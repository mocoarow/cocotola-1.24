package service

import (
	"context"
	"errors"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrCardAlreadyExists = errors.New("card already exists")
var ErrCardNotFound = errors.New("card not found")

type AddCardParameter struct {
	DeckID     *domain.DeckID     `validate:"required"`
	TemplateID *domain.TemplateID `validate:"required"`
	Content    string             `validate:"required"`
}

func NewAddCardParameter(deckID *domain.DeckID, templateID *domain.TemplateID, content string) (*AddCardParameter, error) {
	m := &AddCardParameter{
		DeckID:     deckID,
		TemplateID: templateID,
		Content:    content,
	}
	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, errors.New("validate add card parameter: " + err.Error())
	}
	return m, nil
}

type CardRepository interface {
	AddCard(ctx context.Context, operator mbuserdomain.UserInterface, param *AddCardParameter) (*domain.CardID, error)

	FindCardsByDeckID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*domain.Card, error)
}
