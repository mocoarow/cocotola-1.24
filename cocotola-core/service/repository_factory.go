package service

import (
	"context"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
)

type RepositoryFactory interface {
	// NewWorkbookRepository(ctx context.Context) (WorkbookRepository, error)
	NewFolderRepository(ctx context.Context) FolderRepository
	NewDeckRepository(ctx context.Context) DeckRepository
	NewCardRepository(ctx context.Context) CardRepository
	// NewAppUserRepository(ctx context.Context) (AppUserRepository, error)
}

type TransactionManager mblibservice.TransactionManagerT[RepositoryFactory]
