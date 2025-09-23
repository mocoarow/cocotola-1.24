package guest

import (
	"context"
	"fmt"
	"log/slog"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libapiauth "github.com/mocoarow/cocotola-1.24/lib/api/auth"
	librbac "github.com/mocoarow/cocotola-1.24/lib/rbac"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckQueryUseCase struct {
	db         *gorm.DB
	rbacClient libapi.CocotolaRBACClient
	logger     *slog.Logger
}

func NewDeckQueryUsecase(db *gorm.DB, rbacClient libapi.CocotolaRBACClient) *DeckQueryUseCase {
	return &DeckQueryUseCase{
		db:         db,
		rbacClient: rbacClient,
		logger:     slog.Default().With(slog.String(mbliblog.LoggerNameKey, "DeckQueryUseCase")),
	}
}

func (u *DeckQueryUseCase) filterSpaces(ctx context.Context, operator mbuserdomain.UserInterface, action mbuserdomain.RBACAction, spaceIDs []*mbuserdomain.SpaceID) ([]*mbuserdomain.SpaceID, error) {
	filteredSpaceIDs := make([]*mbuserdomain.SpaceID, 0, len(spaceIDs))
	for _, spaceID := range spaceIDs {
		action := action
		object := spaceID.GetRBACObject()
		ok, err := u.rbacClient.CheckAuthorization(ctx, &libapiauth.AuthorizeRequest{
			OrganizationID: operator.GetOrganizationID().Int(),
			UserID:         operator.GetUserID().Int(),
			Action:         action.Action(),
			Object:         object.Object(),
		})
		if err != nil {
			return nil, mbliberrors.Errorf("authorize: %w", err)
		} else if !ok {
			continue
		}
		filteredSpaceIDs = append(filteredSpaceIDs, spaceID)
	}
	return filteredSpaceIDs, nil
}

func (u *DeckQueryUseCase) FindDecks(ctx context.Context, operator mbuserdomain.UserInterface, param *service.FindDecksParameter) ([]*domain.Deck, error) {
	_, span := tracer.Start(ctx, "DeckQueryUseCase.FindDecks")
	defer span.End()

	// Check RBAC
	filterSpaceIDs, err := u.filterSpaces(ctx, operator, librbac.ListDecksAction, param.SpaceIDs)
	if err != nil {
		return nil, mbliberrors.Errorf("filterSpaces: %w", err)
	}
	if len(filterSpaceIDs) == 0 {
		u.logger.InfoContext(ctx, "no accessible space")
		return []*domain.Deck{}, nil
	}

	deckRepo := gateway.NewDeckRepository(u.db)
	repoParam := service.FindDecksParameter{
		SpaceIDs: filterSpaceIDs,
	}
	decks, err := deckRepo.FindDecks(ctx, operator, &repoParam)
	if err != nil {
		return nil, fmt.Errorf("deckRepo.FindDecksInPublicSpace. err: %w", err)
	}

	return decks, nil
}

/*

	Problem.word(WordProblem(
		japanese: 'この問題は難しいです。',
		english: 'This problem is ___.',
		cefrLevel: 'B1',
		blanks: [
		  BlankAnswer(
			answer: 'difficult',
			hint: '「難しい」という意味の形容詞です。',
		  ),
		],
	  )),
	  Problem.word(WordProblem(
		japanese: '私は彼女に図書館で会った。',
		english: 'I ___ her ___ the library.',
		cefrLevel: 'B1',
		blanks: [
		  BlankAnswer(
			answer: 'met',
			hint: '「会う」の過去形です。',
		  ),
		  BlankAnswer(
			answer: 'at',
			hint: '場所を示す前置詞です。',
		  ),
		],
	  )),
*/
