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
		return nil, mbliberrors.Errorf("check authorization: %w", err)
	} else if !ok {
		return nil, mbliberrors.Errorf("permission denied. space(%d): %w", param.SpaceID.Int(), mblibdomain.ErrPermissionDenied)
	}

	//
	deckID, err := mblibservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.DeckID, error) {
		folderRepo, err := rf.NewFolderRepository(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewFolderRepository:%w", err)
		}
		folder, err := folderRepo.RetrieveRooFolderBySpaceID(ctx, operator, param.SpaceID)
		if err != nil {
			return nil, mbliberrors.Errorf("retrieve root folder by space id(%d): %w", param.SpaceID.Int(), err)
		}
		param.FolderID = folder.FolderID
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		return deckRepo.AddDeck(ctx, operator, param)
	})
	if err != nil {
		return nil, mbliberrors.Errorf("add deck: %w", err)
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
