package guest

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CardQueryUseCase struct {
	db *gorm.DB
}

func NewCardQueryUsecase(db *gorm.DB) *CardQueryUseCase {
	return &CardQueryUseCase{
		db: db,
	}
}

func (u *CardQueryUseCase) FindCards(ctx context.Context, operator service.OperatorInterface) ([]*domain.CardModel, error) {
	cards := make([]*domain.CardModel, 0)

	type BlankAnswer struct {
		Answer string
		Hint   string
	}
	type WordProblem struct {
		Japanese  string
		English   string
		CEFRLevel string
		Blanks    []BlankAnswer
	}
	wordProblems := []WordProblem{
		{
			Japanese:  "私は昨日映画を見ました。",
			English:   "I ___ a movie yesterday.",
			CEFRLevel: "A2",
			Blanks: []BlankAnswer{
				{
					Answer: "watched",
					Hint:   "「見る」の過去形です。",
				},
			},
		},
	}
	organizationID, err := mbuserdomain.NewOrganizationID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
	}
	deckID, err := domain.NewDeckID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
	}

	templateID, err := domain.NewTemplateID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
	}

	ownerID, err := mbuserdomain.NewAppUserID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
	}

	for i, wp := range wordProblems {
		base, err := mblibdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
		if err != nil {
			return nil, mbliberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
		}
		cardID, err := domain.NewCardID(i + 1)
		if err != nil {
			return nil, mbliberrors.Errorf("domain.NewCardID. err: %w", err)
		}

		content, err := json.Marshal(wp)
		if err != nil {
			return nil, err
		}
		card, err := domain.NewCardModel(base, cardID, organizationID, deckID, templateID, string(content), ownerID)
		if err != nil {
			return nil, mbliberrors.Errorf("domain.NewCardModel. err: %w", err)
		}
		cards = append(cards, card)
	}
	return cards, nil
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
