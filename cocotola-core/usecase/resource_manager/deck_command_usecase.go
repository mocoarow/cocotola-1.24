package resourcemanager

import (
	"context"
	"fmt"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckCommandUseCase struct {
	txManager  service.TransactionManager
	rbacClient service.CocotolaRBACClient
}

func NewDeckCommandUsecase(txmanager service.TransactionManager, rbacClient service.CocotolaRBACClient) *DeckCommandUseCase {
	return &DeckCommandUseCase{
		txManager:  txmanager,
		rbacClient: rbacClient,
	}
}

func (u *DeckCommandUseCase) AddDeck(ctx context.Context, operator service.OperatorInterface, deck *service.DeckAddParameter) (*domain.DeckID, error) {
	//
	deckID, err := mblibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.DeckID, error) {
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return nil, err
		}
		return deckRepo.AddDeck(ctx, operator, deck)
	})
	if err != nil {
		return nil, fmt.Errorf("add deck: %w", err)
	}

	// RBAC
	deckObject := fmt.Sprintf("deck:%d", deckID)
	// - "operator "can" "ListObject" "deck"
	u.rbacClient.AddPolicyToUser(ctx, &libapi.AddPolicyToUserParameter{
		OrganizationID: operator.OrganizationID().Int(),
		AppUserID:      operator.AppUserID().Int(),
		ListOfActionObjectEffect: []libapi.ActionObjectEffect{
			{
				Action: mbuserdomain.NewRBACAction("ListObject").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("GetObject").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("DeleteObject").Action(),
				Object: deckObject,
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
		},
	})

	return deckID, nil
}
