package service

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrDeckAlreadyExists = errors.New("deck already exists")

var ErrDeckNotFound = errors.New("deck not found")

type DeckAddParameter struct {
	Name        string
	TemplateID  int
	Lang2       string
	Description string
}

type DeckUpdateParameter struct {
	Name        string
	Description string
}

type DeckRepository interface {
	AddDeck(ctx context.Context, operator OperatorInterface, param *DeckAddParameter) (*domain.DeckID, error)

	UpdateDeck(ctx context.Context, operator OperatorInterface, deckID *domain.DeckID, version int, param *DeckUpdateParameter) error

	FindDecks(ctx context.Context, operator OperatorInterface, deckID *domain.DeckID) ([]*Deck, error)
}
