package service

import (
	"context"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

type AppUser struct {
	operator   mbuserservice.OperatorInterface
	rf         RepositoryFactory
	rbacClient libapi.CocotolaRBACClient
}

func NewAppUser(operator mbuserservice.OperatorInterface, rf RepositoryFactory, rbacClient libapi.CocotolaRBACClient) *AppUser {
	return &AppUser{
		operator:   operator,
		rf:         rf,
		rbacClient: rbacClient,
	}
}

func (m *AppUser) AddDeck(ctx context.Context, param *AddDeckParameter) (*domain.DeckID, error) {
	// Check RBAC
	// Can "operator" "createDeck" in "space" ?
	action := librbac.CreateDeckAction
	object := param.SpaceID.GetRBACObject()
	ok, err := m.rbacClient.CheckAuthorization(ctx, &libapiauth.AuthorizeRequest{
		OrganizationID: m.operator.GetOrganizationID().Int(),
		UserID:         m.operator.GetUserID().Int(),
		Action:         action.Action(),
		Object:         object.Object(),
	})
	if err != nil {
		return nil, mbliberrors.Errorf("check authorization: %w", err)
	} else if !ok {
		return nil, mbliberrors.Errorf("permission denied. space(%d): %w", param.SpaceID.Int(), mblibdomain.ErrPermissionDenied)
	}

	folderRepo, err := m.rf.NewFolderRepository(ctx)
	if err != nil {
		return nil, mbliberrors.Errorf("NewFolderRepository:%w", err)
	}
	folder, err := folderRepo.RetrieveRooFolderBySpaceID(ctx, m.operator, param.SpaceID)
	if err != nil {
		return nil, mbliberrors.Errorf("retrieve root folder by space id(%d): %w", param.SpaceID.Int(), err)
	}
	param.FolderID = folder.FolderID
	deckRepo, err := m.rf.NewDeckRepository(ctx)
	if err != nil {
		return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
	}

	deckID, err := deckRepo.AddDeck(ctx, m.operator, param)
	if err != nil {
		return nil, mbliberrors.Errorf("deckRepo.AddDeck: %w", err)
	}

	// RBAC
	deckObject := deckID.GetRBACObject()
	// - "operator "can" "ListCards" for "deck"
	if err := m.rbacClient.AddPolicyToUser(ctx, &libapiauth.AddPolicyToUserParameter{
		OrganizationID: m.operator.GetOrganizationID().Int(),
		UserID:         m.operator.GetUserID().Int(),
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
