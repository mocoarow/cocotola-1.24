package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type RepositoryFactory struct {
	dialect    mblibgateway.DialectRDBMS
	driverName string
	db         *gorm.DB
	location   *time.Location
}

func NewRepositoryFactory(ctx context.Context, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location) (*RepositoryFactory, error) {
	if db == nil {
		return nil, mbliberrors.Errorf("new repository factory. db is nil: %w", mblibdomain.ErrInvalidArgument)
	}

	return &RepositoryFactory{
		dialect:    dialect,
		driverName: driverName,
		db:         db,
		location:   location,
	}, nil
}

func (f *RepositoryFactory) NewDeckRepository(ctx context.Context) (service.DeckRepository, error) {
	return NewDeckRepository(f.db), nil
}
func (f *RepositoryFactory) NewSpaceRepository(ctx context.Context) (service.SpaceRepository, error) {
	return NewSpaceRepository(f.db), nil
}
func (f *RepositoryFactory) NewPairOfUserAndSpaceRepository(ctx context.Context) (service.PairOfUserAndSpaceRepository, error) {
	return NewPairOfUserAndSpaceRepository(ctx, f.dialect, f.db, f), nil
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
