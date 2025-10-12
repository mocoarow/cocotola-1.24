package guest

import (
	"context"
	"errors"
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type FindDecksQuery struct {
	rf         service.RepositoryFactory
	rbacClient libapi.CocotolaRBACClient
}

func NewFindDecksQuery(rf service.RepositoryFactory, rbacClient libapi.CocotolaRBACClient) *FindDecksQuery {
	return &FindDecksQuery{
		rf:         rf,
		rbacClient: rbacClient,
	}
}

func (u *FindDecksQuery) Execute(ctx context.Context, operator mbuserdomain.UserInterface, param *service.FindDecksParameter) ([]*domain.Deck, error) {
	// 1. Check authorization
	filterSpaceIDs := make([]*mbuserdomain.SpaceID, 0, len(param.SpaceIDs))
	for _, spaceID := range param.SpaceIDs {
		if err := u.checkAuthorization(ctx, operator, spaceID); err != nil {
			if errors.Is(err, mblibdomain.ErrPermissionDenied) {
				continue
			}
			return nil, mbliberrors.Errorf("checkAuthorization: %w", err)
		} else {
			filterSpaceIDs = append(filterSpaceIDs, spaceID)
		}
	}
	if len(filterSpaceIDs) == 0 {
		return []*domain.Deck{}, nil
	}

	// 2. Execute
	decks, err := u.execute(ctx, operator, filterSpaceIDs)
	if err != nil {
		return nil, mbliberrors.Errorf("execute: %w", err)
	}

	return decks, nil
}

func (u *FindDecksQuery) checkAuthorization(ctx context.Context, operator mbuserdomain.UserInterface, spaceID *mbuserdomain.SpaceID) error {
	// Check RBAC
	// Can "operator" "ListDecks" ?
	action := librbac.ListDecksAction
	object := spaceID.GetRBACObject()
	ok, err := u.rbacClient.CheckAuthorization(ctx, &libapiauth.AuthorizeRequest{
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:         operator.GetUserID().Int(),
		Action:         action.Action(),
		Object:         object.Object(),
	})
	if err != nil {
		return mbliberrors.Errorf("check authorization: %w", err)
	} else if !ok {
		return mbliberrors.Errorf("permission denied. space(%d): %w", spaceID.Int(), mblibdomain.ErrPermissionDenied)
	}
	return nil
}

func (u *FindDecksQuery) execute(ctx context.Context, operator mbuserdomain.UserInterface, spaceIDs mbuserdomain.SpaceIDs) ([]*domain.Deck, error) {
	deckRepo := u.rf.NewDeckRepository(ctx)
	repoParam := service.FindDecksParameter{
		SpaceIDs: spaceIDs,
	}
	decks, err := deckRepo.FindDecks(ctx, operator, &repoParam)
	if err != nil {
		return nil, fmt.Errorf("deckRepo.FindDecks. err: %w", err)
	}
	return decks, nil
}
