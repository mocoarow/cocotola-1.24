package initialize

import (
	"context"
	"encoding/json"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type englishBlankAnswer struct {
	Answer string
}

type englishBlankCard struct {
	SourceText   string
	EnglishText  string
	Level        string
	BlankAnswers []englishBlankAnswer
}
type englishBlankDeck struct {
	Name  string
	Lang2 *mblibdomain.Lang2
	Cards []englishBlankCard
}

func getEnglishBlankDecks() []englishBlankDeck {
	return []englishBlankDeck{
		{
			Name:  "初心者向け基本文法",
			Lang2: mblibdomain.Lang2JA,
			Cards: []englishBlankCard{
				{
					SourceText:  "私は毎日英語を勉強します。",
					EnglishText: "I ___ English every day.",
					Level:       "easyA1",
					BlankAnswers: []englishBlankAnswer{
						{Answer: "astudym"},
					},
				},
				{
					SourceText:  "彼は英語を上手に話します。",
					EnglishText: "He speaks English ___.",
					Level:       "A2",
					BlankAnswers: []englishBlankAnswer{
						{Answer: "well"},
					},
				},
			},
		},
		{
			Name:  "中級文法チャレンジ",
			Lang2: mblibdomain.Lang2JA,
			Cards: []englishBlankCard{
				{
					SourceText:  "私は彼女に図書館で会った。",
					EnglishText: "I ___ her ___ the library.",
					Level:       "B1",
					BlankAnswers: []englishBlankAnswer{
						{Answer: "met"},
						{Answer: "at"},
					},
				},
				{
					SourceText:  "彼女は毎週ピアノを練習しています。",
					EnglishText: "She ___ the piano every week.",
					Level:       "A1",
					BlankAnswers: []englishBlankAnswer{
						{Answer: "practices"},
					},
				},
			},
		},
	}
}

func initEnglishBlankDeck(ctx context.Context, operator mbuserdomain.UserInterface, deckRepo service.DeckRepository, cardRepo service.CardRepository, publicDefaultSpaceID *mbuserdomain.SpaceID, rootFolderID *domain.FolderID, nameToDecks map[string]*service.Deck) ([]*domain.DeckID, error) {
	deckIDs := make([]*domain.DeckID, 0)
	for _, englishBlankDeck := range getEnglishBlankDecks() {
		name := englishBlankDeck.Name
		if _, exists := nameToDecks[name]; exists {
			continue
		}

		deckAddParam := service.AddDeckParameter{
			SpaceID:     publicDefaultSpaceID,
			FolderID:    rootFolderID,
			Name:        name,
			TemplateID:  service.TemplateIDEnglishBlank,
			Lang2:       mblibdomain.Lang2JA,
			Description: "",
		}

		deckID, err := deckRepo.AddDeck(ctx, operator, &deckAddParam)
		if err != nil {
			return nil, mbliberrors.Errorf("add deck: %w", err)
		}

		for _, englishBlankCard := range englishBlankDeck.Cards {
			content, err := json.Marshal(englishBlankCard)
			if err != nil {
				return nil, mbliberrors.Errorf("json.Marshal (%+v): %w", err)
			}

			addCardParam := service.AddCardParameter{
				DeckID:     deckID,
				TemplateID: service.TemplateIDEnglishBlank,
				Content:    string(content),
			}
			if _, err := cardRepo.AddCard(ctx, operator, &addCardParam); err != nil {
				return nil, mbliberrors.Errorf("add card: %w", err)
			}
		}
		deckIDs = append(deckIDs, deckID)
	}

	return deckIDs, nil
}

func initEnglishWord(ctx context.Context, txManager service.TransactionManager, organizationID *mbuserdomain.OrganizationID, systemOwnerID *mbuserdomain.UserID, publicDefaultSpaceID *mbuserdomain.SpaceID, rootFolderID *domain.FolderID) ([]*domain.DeckID, error) {
	operator := &operator{
		organizationID: organizationID,
		userID:         systemOwnerID,
	}

	fn := func(rf service.RepositoryFactory) ([]*domain.DeckID, error) {
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		decks, err := deckRepo.FindDecksByOwner(ctx, operator)
		if err != nil {
			return nil, mbliberrors.Errorf("FindDecksByOwner: %w", err)
		}

		cardRepo, err := rf.NewCardRepository(ctx)
		if err != nil {
			return nil, mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		nameToDecks := make(map[string]*service.Deck, len(decks))
		for _, deck := range decks {
			nameToDecks[deck.Name] = deck
		}

		deckIDs, err := initEnglishBlankDeck(ctx, operator, deckRepo, cardRepo, publicDefaultSpaceID, rootFolderID, nameToDecks)
		if err != nil {
			return nil, mbliberrors.Errorf("initEnglishBlankDeck: %w", err)
		}

		return deckIDs, nil
	}

	deckIDs, err := mblibservice.Do1(ctx, txManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return deckIDs, nil
}
