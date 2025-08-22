package resourcemanager

import (
	"context"
	"fmt"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckCommandUseCase struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   service.CocotolaRBACClient
}

func NewDeckCommandUsecase(txManager, nonTxManager service.TransactionManager, rbacClient service.CocotolaRBACClient) *DeckCommandUseCase {
	return &DeckCommandUseCase{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
	}
}

func (u *DeckCommandUseCase) AddDeck(ctx context.Context, operator service.OperatorInterface, param *service.DeckAddParameter) (*domain.DeckID, error) {
	//
	deckID, err := mblibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.DeckID, error) {
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return nil, err
		}
		return deckRepo.AddDeck(ctx, operator, param)
	})
	if err != nil {
		return nil, mbliberrors.Errorf("add deck: %w", err)
	}

	// RBAC
	deckObject := fmt.Sprintf("deck:%d", deckID)
	// - "operator "can" "ListObject" "deck"
	u.rbacClient.AddPolicyToUser(ctx, &libapi.AddPolicyToUserParameter{
		OrganizationID: operator.OrganizationID().Int(),
		AppUserID:      operator.AppUserID().Int(),
		ListOfActionObjectEffect: []libapi.ActionObjectEffect{
			{
				Action: mbuserdomain.NewRBACAction("ListCards").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("GetDeck").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("DeleteDeck").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("UpdateDeck").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
		},
	})

	return deckID, nil
}

func (u *DeckCommandUseCase) UpdateDeck(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID, version int, param *service.DeckUpdateParameter) error {
	//
	err := mblibservice.Do0(ctx, u.txManager, func(rf service.RepositoryFactory) error {
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return err
		}
		return deckRepo.UpdateDeck(ctx, operator, deckID, version, param)
	})
	if err != nil {
		return mbliberrors.Errorf("add deck: %w", err)
	}

	return nil
}
