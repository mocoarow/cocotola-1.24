package guest

import (
	"context"
	"fmt"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type CardUsecase struct {
	mbrf service.RepositoryFactory
}

func NewCardUsecase(mbrf service.RepositoryFactory) *CardUsecase {
	return &CardUsecase{
		mbrf: mbrf,
	}
}

func (u *CardUsecase) FindCardsByDeckID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*domain.Card, error) {
	_, span := tracer.Start(ctx, "CardQueryUseCase.FindDecks")
	defer span.End()

	// check RBAC

	cardRepo := u.mbrf.NewCardRepository(ctx)
	cards, err := cardRepo.FindCardsByDeckID(ctx, operator, deckID)
	if err != nil {
		return nil, fmt.Errorf("cardRepo.FindCardsByDeckID: %w", err)
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
