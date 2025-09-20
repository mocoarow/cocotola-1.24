package guest

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
)

type CardQueryUsecase struct {
	db *gorm.DB
}

func NewCardQueryUsecase(db *gorm.DB) *CardQueryUsecase {
	return &CardQueryUsecase{
		db: db,
	}
}

func (u *CardQueryUsecase) FindCardsByDeckID(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID) ([]*domain.CardModel, error) {
	_, span := tracer.Start(ctx, "CardQueryUseCase.FindDecks")
	defer span.End()

	// check RBAC

	cardRepo := gateway.NewCardRepository(u.db)
	desks, err := cardRepo.FindCardsByDeckID(ctx, operator, deckID)
	if err != nil {
		return nil, fmt.Errorf("cardRepo.FindCardsByDeckID. err: %w", err)
	}

	cardModels := make([]*domain.CardModel, 0, len(desks))
	for _, d := range desks {
		cardModels = append(cardModels, d.CardModel)
	}

	return cardModels, nil
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
