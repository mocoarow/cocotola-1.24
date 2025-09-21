package service

import (
	"context"
	"errors"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrDeckAlreadyExists = errors.New("deck already exists")
var ErrDeckNotFound = errors.New("deck not found")

type AddDeckParameter struct {
	SpaceID     *mbuserdomain.SpaceID
	FolderID    *domain.FolderID
	TemplateID  *domain.TemplateID
	Name        string
	Lang2       *mblibdomain.Lang2
	Description string
}

type UpdateDeckParameter struct {
	Name        string
	Description string
}

type FindDecksParameter struct {
	SpaceIDs mbuserdomain.SpaceIDs
}

type DeckRepository interface {
	AddDeck(ctx context.Context, operator mbuserdomain.UserInterface, param *AddDeckParameter) (*domain.DeckID, error)

	UpdateDeck(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID, version int, param *UpdateDeckParameter) error

	FindDecks(ctx context.Context, operator mbuserdomain.UserInterface, param *FindDecksParameter) ([]*Deck, error)

	FindDecksByOwner(ctx context.Context, operator mbuserdomain.UserInterface) ([]*Deck, error)

	RetrieveDeckByID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) (*Deck, error)
}
