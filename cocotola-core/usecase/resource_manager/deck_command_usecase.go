package resourcemanager

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckCommandUseCase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   libapi.CocotolaRBACClient
}

func NewDeckCommandUsecase(txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) *DeckCommandUseCase {
	return &DeckCommandUseCase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
	}
}

func (u *DeckCommandUseCase) AddDeck(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddDeckParameter) (*domain.DeckID, error) {
	command := NewAddDeckCommand(u.txManager, u.nonTxManager, u.rbacClient)
	deckID, err := command.execute(ctx, operator, param)
	if err != nil {
		return nil, mbliberrors.Errorf("add deck: %w", err)
	}

	return deckID, nil
}

func (u *DeckCommandUseCase) UpdateDeck(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID, version int, param *service.UpdateDeckParameter) error {
	if err := mblibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		if err := deckRepo.UpdateDeck(ctx, operator, deckID, version, param); err != nil {
			return mbliberrors.Errorf("update deck: %w", err)
		}

		return nil
	}); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
