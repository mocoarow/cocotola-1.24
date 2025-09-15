package service

import (
	"context"
	"errors"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrDeckAlreadyExists = errors.New("deck already exists")
var ErrDeckNotFound = errors.New("deck not found")

type AddDeckParameter struct {
	SpaceID     *mbuserdomain.SpaceID
	FolderID    *domain.FolderID
	TemplateID  *domain.TemplateID
	Name        string
	Lang2       *libdomain.Lang2
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
	AddDeck(ctx context.Context, operator mbuserservice.OperatorInterface, param *AddDeckParameter) (*domain.DeckID, error)

	UpdateDeck(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID, version int, param *UpdateDeckParameter) error

	FindDecks(ctx context.Context, operator mbuserservice.OperatorInterface, param *FindDecksParameter) ([]*Deck, error)

	FindDecksByOwner(ctx context.Context, operator mbuserservice.OperatorInterface) ([]*Deck, error)

	RetrieveDeckByID(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID) (*Deck, error)
}
