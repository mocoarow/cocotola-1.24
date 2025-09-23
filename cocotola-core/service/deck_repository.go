package service

import (
	"context"
	"errors"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrDeckAlreadyExists = errors.New("deck already exists")
var ErrDeckNotFound = errors.New("deck not found")

type AddDeckParameter struct {
	SpaceID     *mbuserdomain.SpaceID `validate:"required"`
	FolderID    *domain.FolderID      `validate:"required"`
	TemplateID  *domain.TemplateID    `validate:"required"`
	Name        string                `validate:"required"`
	Lang2       *mblibdomain.Lang2    `validate:"required"`
	Description string                `validate:"omitempty,max=100"`
}

func NewAddDeckParameter(spaceID *mbuserdomain.SpaceID, folderID *domain.FolderID, templateID *domain.TemplateID, name string, lang2 *mblibdomain.Lang2, description string) (*AddDeckParameter, error) {
	m := &AddDeckParameter{
		SpaceID:     spaceID,
		FolderID:    folderID,
		TemplateID:  templateID,
		Name:        name,
		Lang2:       lang2,
		Description: description,
	}
	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate add deck parameter: %w", err)
	}
	return m, nil
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

	FindDecks(ctx context.Context, operator mbuserdomain.UserInterface, param *FindDecksParameter) ([]*domain.Deck, error)

	FindDecksByOwner(ctx context.Context, operator mbuserdomain.UserInterface) ([]*domain.Deck, error)

	RetrieveDeckByID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) (*domain.Deck, error)
}
