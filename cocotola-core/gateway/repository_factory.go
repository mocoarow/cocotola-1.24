package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

var _ service.RepositoryFactory = (*RepositoryFactory)(nil)

type RepositoryFactory struct {
	dialect    mblibgateway.DialectRDBMS
	driverName string
	db         *gorm.DB
	location   *time.Location
	rbacClient libapi.CocotolaRBACClient
}

func NewRepositoryFactory(_ context.Context, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location, rbacClient libapi.CocotolaRBACClient) *RepositoryFactory {
	if db == nil {
		panic("new repository factory. db is nil")
	}

	return &RepositoryFactory{
		dialect:    dialect,
		driverName: driverName,
		db:         db,
		location:   location,
		rbacClient: rbacClient,
	}
}

func (f *RepositoryFactory) NewDeckRepository(_ context.Context) service.DeckRepository {
	return NewDeckRepository(f.db)
}
func (f *RepositoryFactory) NewCardRepository(_ context.Context) service.CardRepository {
	return NewCardRepository(f.db)
}
func (f *RepositoryFactory) NewFolderRepository(_ context.Context) service.FolderRepository {
	return NewFolderRepository(f.db)
}

// func (f *RepositoryFactory) NewAppUserRepository(_ context.Context) (service.AppUserRepository, error) {
// 	return NewAppUserRepository(f.db, f, f.rbacClient), nil
// }

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
