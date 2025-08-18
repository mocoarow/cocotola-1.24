package resourcemanager

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckQueryUseCase struct {
	txManager service.TransactionManager
}

func NewDeckQueryUsecase(txManager service.TransactionManager) *DeckQueryUseCase {
	return &DeckQueryUseCase{
		txManager: txManager,
	}
}
func (u *DeckQueryUseCase) FindDecks(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID) ([]*service.Deck, error) {
	return mblibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) ([]*service.Deck, error) {
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return nil, err
		}
		return deckRepo.FindDecks(ctx, operator, deckID)
	})
}
