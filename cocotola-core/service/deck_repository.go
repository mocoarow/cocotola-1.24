package service

import (
	"context"
	"errors"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrDeckAlreadyExists = errors.New("deck already exists")
var ErrDeckNotFound = errors.New("deck not found")

type DeckAddParameter struct {
	SpaceID     *domain.SpaceID
	FolderID    *domain.FolderID
	TemplateID  *domain.TemplateID
	Name        string
	Lang2       *libdomain.Lang2
	Description string
}

type DeckUpdateParameter struct {
	Name        string
	Description string
}

type FindDecksParameter struct {
	SpaceIDs domain.SpaceIDs
}

type DeckRepository interface {
	AddDeck(ctx context.Context, operator mbuserservice.OperatorInterface, param *DeckAddParameter) (*domain.DeckID, error)

	UpdateDeck(ctx context.Context, operator OperatorInterface, deckID *domain.DeckID, version int, param *DeckUpdateParameter) error

	FindDecks(ctx context.Context, operator OperatorInterface, param *FindDecksParameter) ([]*Deck, error)

	FindDecksByOwner(ctx context.Context, operator mbuserservice.OperatorInterface) ([]*Deck, error)

	FindDecksInPublicSpace(ctx context.Context, operator OperatorInterface) ([]*Deck, error)

	RetrieveDeckByID(ctx context.Context, operator OperatorInterface, deckID *domain.DeckID) (*Deck, error)
}
