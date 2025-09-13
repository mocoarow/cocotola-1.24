package guest

import (
	"context"
	"time"

	"gorm.io/gorm"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckQueryUseCase struct {
	db *gorm.DB
}

func NewDeckQueryUsecase(db *gorm.DB) *DeckQueryUseCase {
	return &DeckQueryUseCase{
		db: db,
	}
}

func (u *DeckQueryUseCase) FindDecks(ctx context.Context, operator service.OperatorInterface) ([]*domain.DeckModel, error) {
	_, span := tracer.Start(ctx, "DeckQueryUseCase.FindDecks")
	defer span.End()

	decks := make([]*domain.DeckModel, 0)

	organizationID, err := mbuserdomain.NewOrganizationID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("new organization id(1). err: %w", err)
	}
	deckID, err := domain.NewDeckID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("new deck id(1). err: %w", err)
	}
	spaceID, err := domain.NewSpaceID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(1). err: %w", err)
	}
	folderID, err := domain.NewFolderID(0)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder id(0). err: %w", err)
	}

	templateID, err := domain.NewTemplateID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("new template id(1). err: %w", err)
	}

	ownerID, err := mbuserdomain.NewAppUserID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("new app user id(1). err: %w", err)
	}

	{
		base, err := mblibdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
		if err != nil {
			return nil, mbliberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
		}
		deck, err := domain.NewDeckModel(base, deckID, organizationID, spaceID, folderID, "Sample", templateID, libdomain.Lang2JA, "", ownerID)
		if err != nil {
			return nil, mbliberrors.Errorf("domain.NewCardModel. err: %w", err)
		}
		decks = append(decks, deck)
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
