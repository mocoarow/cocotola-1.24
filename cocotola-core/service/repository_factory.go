package service

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
)

type RepositoryFactory interface {
	// NewWorkbookRepository(ctx context.Context) (WorkbookRepository, error)
	NewFolderRepository(ctx context.Context) (FolderRepository, error)
	NewDeckRepository(ctx context.Context) (DeckRepository, error)
	NewCardRepository(ctx context.Context) (CardRepository, error)
}

type TransactionManager mblibservice.TransactionManagerT[RepositoryFactory]
