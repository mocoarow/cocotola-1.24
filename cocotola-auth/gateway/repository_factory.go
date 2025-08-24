package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type RepositoryFactory struct {
	dialect             mblibgateway.DialectRDBMS
	driverName          string
	db                  *gorm.DB
	location            *time.Location
	appUserEventHandler mblibservice.ResourceEventHandler
}

func NewRepositoryFactory(_ context.Context, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location, appUserEventHandler mblibservice.ResourceEventHandler) (*RepositoryFactory, error) {
	if db == nil {
		return nil, mbliberrors.Errorf("db is nil. err: %w", mblibdomain.ErrInvalidArgument)
	}

	return &RepositoryFactory{
		dialect:             dialect,
		driverName:          driverName,
		db:                  db,
		location:            location,
		appUserEventHandler: appUserEventHandler,
	}, nil
}

func (f *RepositoryFactory) NewMoonBeamRepositoryFactory(ctx context.Context) (mbuserservice.RepositoryFactory, error) {
	return mbusergateway.NewRepositoryFactory(ctx, f.dialect, f.driverName, f.db, f.location, f.appUserEventHandler)
}

func (f *RepositoryFactory) NewStateRepository(ctx context.Context) (service.StateRepository, error) {
	return NewStateRepository(ctx)
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
