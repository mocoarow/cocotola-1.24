package resourcemanager

import (
	"context"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckQueryUseCase struct {
	db *gorm.DB
}

func NewDeckQueryUsecase(db *gorm.DB) *DeckQueryUseCase {
	return &DeckQueryUseCase{
		db: db,
	}
}

func (u *DeckQueryUseCase) FindDecks(ctx context.Context, operator mbuserdomain.UserInterface) ([]*domain.DeckModel, error) {
	// TODO: filter by spaceIDs
	decks, err := gateway.NewDeckRepository(u.db).FindDecks(ctx, operator, &service.FindDecksParameter{}) //nolint:exhaustruct
	if err != nil {
		return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
	}
	deckModels := make([]*domain.DeckModel, len(decks))
	for i, deck := range decks {
		deckModels[i] = deck.DeckModel
	}

	return deckModels, nil
}

func (u *DeckQueryUseCase) RetrieveDeckByID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) (*domain.DeckModel, error) {
	deck, err := gateway.NewDeckRepository(u.db).RetrieveDeckByID(ctx, operator, deckID)
	if err != nil {
		return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
	}

	return deck.DeckModel, nil
}
