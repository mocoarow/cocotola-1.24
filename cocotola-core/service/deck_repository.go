package service

import (
	"context"
	"errors"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrDeckAlreadyExists = errors.New("deck already exists")
var ErrDeckNotFound = errors.New("deck not found")

type DeckAddParameter struct {
	SpaceID     *domain.SpaceID
	FolderID    *domain.FolderID
	TemplateID  *domain.TemplateID
	Name        string
	Lang2       string
	Description string
}

type DeckUpdateParameter struct {
	Name        string
	Description string
}

type DeckRepository interface {
	AddDeck(ctx context.Context, operator mbuserservice.OperatorInterface, param *DeckAddParameter) (*domain.DeckID, error)

	UpdateDeck(ctx context.Context, operator OperatorInterface, deckID *domain.DeckID, version int, param *DeckUpdateParameter) error

	FindDecks(ctx context.Context, operator OperatorInterface) ([]*Deck, error)

	RetrieveDeckByID(ctx context.Context, operator OperatorInterface, deckID *domain.DeckID) (*Deck, error)
}
