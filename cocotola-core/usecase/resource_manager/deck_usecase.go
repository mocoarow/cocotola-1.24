package resourcemanager

import (
	"context"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckUsecase struct {
	rf           service.RepositoryFactory
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   libapi.CocotolaRBACClient
}

func NewDeckUsecase(rf service.RepositoryFactory, txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) *DeckUsecase {
	return &DeckUsecase{
		rf:           rf,
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
	}
}

func (u *DeckUsecase) FindDecks(ctx context.Context, operator mbuserdomain.UserInterface) ([]*domain.Deck, error) {
	deckRepo := u.rf.NewDeckRepository(ctx)
	// TODO: filter by spaceIDs
	decks, err := deckRepo.FindDecks(ctx, operator, &service.FindDecksParameter{}) //nolint:exhaustruct
	if err != nil {
		return nil, mbliberrors.Errorf("FindDecks: %w", err)
	}

	return decks, nil
}

func (u *DeckUsecase) RetrieveDeckByID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) (*domain.Deck, error) {
	deckRepo := u.rf.NewDeckRepository(ctx)
	deck, err := deckRepo.RetrieveDeckByID(ctx, operator, deckID)
	if err != nil {
		return nil, mbliberrors.Errorf("RetrieveDeckByID: %w", err)
	}

	return deck, nil
}

func (u *DeckUsecase) AddDeck(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddDeckParameter) (*domain.DeckID, error) {
	command := NewAddDeckCommand(u.txManager, u.nonTxManager, u.rbacClient)
	deckID, err := command.execute(ctx, operator, param)
	if err != nil {
		return nil, mbliberrors.Errorf("add deck: %w", err)
	}

	return deckID, nil
}

func (u *DeckUsecase) UpdateDeck(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID, version int, param *service.UpdateDeckParameter) error {
	if err := mblibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		deckRepo := rf.NewDeckRepository(ctx)
		if err := deckRepo.UpdateDeck(ctx, operator, deckID, version, param); err != nil {
			return mbliberrors.Errorf("update deck: %w", err)
		}

		return nil
	}); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
