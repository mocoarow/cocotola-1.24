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

func initEnglishWord(ctx context.Context, txManager service.TransactionManager, organizationID *mbuserdomain.OrganizationID) error {
	operator := &operator{
		organizationID: organizationID,
		appUserID:      mbuserservice.SystemAdminID,
	}
	folderID, err := domain.NewFolderID(0)
	if err != nil {
		return mbliberrors.Errorf("new folder id(0). err: %w", err)
	}
	templateID, err := domain.NewTemplateID(1)
	if err != nil {
		return mbliberrors.Errorf("new template id(1). err: %w", err)
	}

	fn := func(rf service.RepositoryFactory) error {
		spaceRepo, err := rf.NewSpaceRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewSpaceRepository: %w", err)
		}

		// check default-public space
		defaultPublicSpace, err := spaceRepo.FindPublicSpaceByKey(ctx, "default-public")
		if err != nil {
			return nil
		}
		deckRepo, err := rf.NewDeckRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewDeckRepository: %w", err)
		}

		decks, err := deckRepo.FindDecksByOwner(ctx, operator)
		if err != nil {
			return mbliberrors.Errorf("FindDecksByOwner: %w", err)
		}

		deckNames := make(map[string]struct{}, len(decks))
		for _, d := range decks {
			deckNames[d.Name] = struct{}{}
		}

		for _, name := range []string{"初心者向け基本文法", "中級文法チャレンジ"} {
			if _, exists := deckNames[name]; exists {
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
			if _, err = deckRepo.AddDeck(ctx, operator, &deckAddParam); err != nil {
				return mbliberrors.Errorf("add deck: %w", err)
			}
		}
		return nil
	}

	if err := mblibservice.Do0(ctx, txManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
