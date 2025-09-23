package resourcemanager

import (
	"context"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type AddDeckCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   libapi.CocotolaRBACClient
}

func NewAddDeckCommand(txManager, nonTxManager service.TransactionManager, rbacClient libapi.CocotolaRBACClient) *AddDeckCommand {
	return &AddDeckCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
	}
}

func (u *AddDeckCommand) Execute(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddDeckParameter) (*domain.DeckID, error) {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator, param); err != nil {
		return nil, mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	deckID, err := u.execute(ctx, operator, param)
	if err != nil {
		return nil, mbliberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(); err != nil {
		return nil, mbliberrors.Errorf("callback: %w", err)
	}

	return deckID, nil
}
func (u *AddDeckCommand) checkAuthorization(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddDeckParameter) error {
	// Check RBAC
	// Can "operator" "createDeck" in "space" ?
	action := librbac.CreateDeckAction
	object := param.SpaceID.GetRBACObject()
	ok, err := u.rbacClient.CheckAuthorization(ctx, &libapiauth.AuthorizeRequest{
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:         operator.GetUserID().Int(),
		Action:         action.Action(),
		Object:         object.Object(),
	})
	if err != nil {
		return mbliberrors.Errorf("check authorization: %w", err)
	} else if !ok {
		return mbliberrors.Errorf("permission denied. space(%d): %w", param.SpaceID.Int(), mblibdomain.ErrPermissionDenied)
	}
	return nil
}
func (u *AddDeckCommand) execute(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddDeckParameter) (*domain.DeckID, error) {
	deckID, err := mblibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.DeckID, error) {
		// folderRepo, err := rf.NewFolderRepository(ctx)
		// if err != nil {
		// 	return nil, mbliberrors.Errorf("NewFolderRepository:%w", err)
		// }
		// folder, err := folderRepo.RetrieveRooFolderBySpaceID(ctx, operator, param.SpaceID)
		// if err != nil {
		// 	return nil, mbliberrors.Errorf("retrieve root folder by space id(%d): %w", param.SpaceID.Int(), err)
		// }
		// param.FolderID = folder.FolderID
		deckRepo := rf.NewDeckRepository(ctx)
		deckID, err := deckRepo.AddDeck(ctx, operator, param)
		if err != nil {
			return nil, mbliberrors.Errorf("deckRepo.AddDeck: %w", err)
		}

		return deckID, nil
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	// RBAC
	deckObject := deckID.GetRBACObject()
	// - "operator "can" "ListCards" for "deck"
	if err := u.rbacClient.AddPolicyToUser(ctx, &libapiauth.AddPolicyToUserParameter{
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:         operator.GetUserID().Int(),
		ListOfActionObjectEffect: []libapiauth.ActionObjectEffect{
			{
				Action: mbuserdomain.NewRBACAction("ListCards").Action(),
				Object: deckObject.Object(),
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("GetDeck").Action(),
				Object: deckObject.Object(),
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("DeleteDeck").Action(),
				Object: deckObject.Object(),
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
			{
				Action: mbuserdomain.NewRBACAction("UpdateDeck").Action(),
				Object: deckObject.Object(),
				Effect: mbuserservice.RBACAllowEffect.Effect(),
			},
		},
	}); err != nil {
		return nil, mbliberrors.Errorf("add policy to user: %w", err)
	}

	return deckID, nil
}

func (u *AddDeckCommand) callback() error {
	return nil
}
