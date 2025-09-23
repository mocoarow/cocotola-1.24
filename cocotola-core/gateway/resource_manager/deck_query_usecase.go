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

func (u *DeckQueryUseCase) FindDecks(ctx context.Context, operator mbuserdomain.UserInterface) ([]*domain.Deck, error) {
	deckRepo := gateway.NewDeckRepository(u.db)
	// TODO: filter by spaceIDs
	decks, err := deckRepo.FindDecks(ctx, operator, &service.FindDecksParameter{}) //nolint:exhaustruct
	if err != nil {
		return nil, mbliberrors.Errorf("FindDecks: %w", err)
	}

	return decks, nil
}

func (u *DeckQueryUseCase) RetrieveDeckByID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) (*domain.Deck, error) {
	deckRepo := gateway.NewDeckRepository(u.db)
	deck, err := deckRepo.RetrieveDeckByID(ctx, operator, deckID)
	if err != nil {
		return nil, mbliberrors.Errorf("RetrieveDeckByID: %w", err)
	}

	return deck, nil
}
