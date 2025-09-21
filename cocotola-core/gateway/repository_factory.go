package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
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

func NewRepositoryFactory(_ context.Context, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location, rbacClient libapi.CocotolaRBACClient) (*RepositoryFactory, error) {
	if db == nil {
		return nil, mbliberrors.Errorf("new repository factory. db is nil: %w", mblibdomain.ErrInvalidArgument)
	}

	return &RepositoryFactory{
		dialect:    dialect,
		driverName: driverName,
		db:         db,
		location:   location,
		rbacClient: rbacClient,
	}, nil
}

func (f *RepositoryFactory) NewDeckRepository(_ context.Context) (service.DeckRepository, error) {
	return NewDeckRepository(f.db), nil
}
func (f *RepositoryFactory) NewCardRepository(_ context.Context) (service.CardRepository, error) {
	return NewCardRepository(f.db), nil
}
func (f *RepositoryFactory) NewFolderRepository(_ context.Context) (service.FolderRepository, error) {
	return NewFolderRepository(f.db), nil
}

// func (f *RepositoryFactory) NewAppUserRepository(_ context.Context) (service.AppUserRepository, error) {
// 	return NewAppUserRepository(f.db, f, f.rbacClient), nil
// }

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
