package guest

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
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

	deckRepo := gateway.NewDeckRepository(u.db)
	desks, err := deckRepo.FindDecksInPublicSpace(ctx, operator)
	if err != nil {
		return nil, fmt.Errorf("deckRepo.FindDecksInPublicSpace. err: %w", err)
	}

	deckModels := make([]*domain.DeckModel, 0, len(desks))
	for _, d := range desks {
		deckModels = append(deckModels, d.DeckModel)
	}

	return deckModels, nil
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
