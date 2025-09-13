package initialize

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

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
	Lang2 *libdomain.Lang2
	Cards []englishBlankCard
}

func getEnglishBlankDecks() []englishBlankDeck {
	return []englishBlankDeck{
		{
			Name:  "初心者向け基本文法",
			Lang2: libdomain.Lang2JA,
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
			Lang2: libdomain.Lang2JA,
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
func initEnglishBlankDeck(ctx context.Context, operator mbuserservice.OperatorInterface, deckRepo service.DeckRepository, cardRepo service.CardRepository, defaultPublicSpace *service.Space, nameToDecks map[string]*service.Deck) error {
	folderID, err := domain.NewFolderID(0)
	if err != nil {
		return mbliberrors.Errorf("new folder id(0). err: %w", err)
	}

	templateID, err := domain.NewTemplateID(1)
	if err != nil {
		return mbliberrors.Errorf("new template id(1). err: %w", err)
	}

	for _, englishBlankDeck := range getEnglishBlankDecks() {
		name := englishBlankDeck.Name
		if _, exists := nameToDecks[name]; exists {
			continue
		}

		deckAddParam := service.DeckAddParameter{
			SpaceID:     defaultPublicSpace.SpaceID,
			FolderID:    folderID,
			Name:        name,
			TemplateID:  templateID,
			Lang2:       libdomain.Lang2JA,
			Description: "",
		}

		deckID, err := deckRepo.AddDeck(ctx, operator, &deckAddParam)
		if err != nil {
			return mbliberrors.Errorf("add deck: %w", err)
		}

		for _, englishBlankCard := range englishBlankDeck.Cards {
			addCardParam := service.AddCardParameter{
				DeckID:     deckID,
				TemplateID: templateID,
				Content:    englishBlankCard.EnglishText,
			}
			if _, err := cardRepo.AddCard(ctx, operator, &addCardParam); err != nil {
				return mbliberrors.Errorf("add card: %w", err)
			}
		}
	}

	return nil
}

func initEnglishWord(ctx context.Context, txManager service.TransactionManager, organizationID *mbuserdomain.OrganizationID) error {
	operator := &operator{
		organizationID: organizationID,
		appUserID:      mbuserservice.SystemAdminID,
	}

	fn := func(rf service.RepositoryFactory) error {
		spaceRepo, err := rf.NewSpaceRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewSpaceRepository: %w", err)
		}

		defaultPublicSpace, err := spaceRepo.FindPublicSpaceByKey(ctx, "default-public")
		if err != nil {
			return mbliberrors.Errorf("FindPublicSpaceByKey: %w", err)
		}

		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		decks, err := deckRepo.FindDecksByOwner(ctx, operator)
		if err != nil {
			return mbliberrors.Errorf("FindDecksByOwner: %w", err)
		}

		cardRepo, err := rf.NewCardRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		nameToDecks := make(map[string]*service.Deck, len(decks))
		for _, deck := range decks {
			nameToDecks[deck.Name] = deck
		}

		if err := initEnglishBlankDeck(ctx, operator, deckRepo, cardRepo, defaultPublicSpace, nameToDecks); err != nil {
			return mbliberrors.Errorf("initEnglishBlankDeck: %w", err)
		}

		return nil
	}

	if err := mblibservice.Do0(ctx, txManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
