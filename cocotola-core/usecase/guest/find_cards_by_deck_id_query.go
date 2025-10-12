package guest

import (
	"context"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type FindCardsByDeckIDQuery struct {
	rf         service.RepositoryFactory
	rbacClient libapi.CocotolaRBACClient
}

func NewFindCardsByDeckIDQuery(rf service.RepositoryFactory, rbacClient libapi.CocotolaRBACClient) *FindCardsByDeckIDQuery {
	return &FindCardsByDeckIDQuery{
		rf:         rf,
		rbacClient: rbacClient,
	}
}

func (u *FindCardsByDeckIDQuery) Execute(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*domain.Card, error) {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator, deckID); err != nil {
		return nil, mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	cards, err := u.execute(ctx, operator, deckID)
	if err != nil {
		return nil, mbliberrors.Errorf("execute: %w", err)
	}

	return cards, nil
}

func (u *FindCardsByDeckIDQuery) checkAuthorization(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) error {
	// Check RBAC
	// Can "operator" "ReadDeck" ?
	action := librbac.ReadDeckAction
	object := deckID.GetRBACObject()
	ok, err := u.rbacClient.CheckAuthorization(ctx, &libapiauth.AuthorizeRequest{
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:         operator.GetUserID().Int(),
		Action:         action.Action(),
		Object:         object.Object(),
	})
	if err != nil {
		return mbliberrors.Errorf("check authorization: %w", err)
	} else if !ok {
		return mbliberrors.Errorf("permission denied. deck(%d): %w", deckID.Int(), mblibdomain.ErrPermissionDenied)
	}
	return nil
}

func (u *FindCardsByDeckIDQuery) execute(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*domain.Card, error) {
	cardRepo := u.rf.NewCardRepository(ctx)
	cards, err := cardRepo.FindCardsByDeckID(ctx, operator, deckID)
	if err != nil {
		return nil, mbliberrors.Errorf("find cards by deck id(%d): %w", deckID.Int(), err)
	}
	return cards, nil
}
