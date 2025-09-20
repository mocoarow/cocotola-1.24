package resourcemanager

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/usecase"
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

func (u *DeckCommandUseCase) AddDeck(ctx context.Context, operator mbuserservice.OperatorInterface, param *service.AddDeckParameter) (*domain.DeckID, error) {
	appUser := usecase.NewAppUser(operator, u.txManager, u.nonTxManager, u.rbacClient)
	deckID, err := appUser.AddDeck(ctx, param)
	if err != nil {
		return nil, mbliberrors.Errorf("add deck: %w", err)
	}

	return deckID, nil
}

func (u *DeckCommandUseCase) UpdateDeck(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID, version int, param *service.UpdateDeckParameter) error {
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
