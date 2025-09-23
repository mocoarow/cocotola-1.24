package service

import (
	"context"
	"errors"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrCardAlreadyExists = errors.New("card already exists")
var ErrCardNotFound = errors.New("card not found")

type AddCardParameter struct {
	DeckID     *domain.DeckID
	TemplateID *domain.TemplateID
	Content    string
}

type CardRepository interface {
	AddCard(ctx context.Context, operator mbuserdomain.UserInterface, param *AddCardParameter) (*domain.CardID, error)

	FindCardsByDeckID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*domain.Card, error)
}
