package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type repositoryFactory struct {
	dialect              libgateway.DialectRDBMS
	driverName           string
	db                   *gorm.DB
	location             *time.Location
	apppUserEventHandler mblibservice.ResourceEventHandler
}

func NewRepositoryFactory(_ context.Context, dialect libgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location, apppUserEventHandler mblibservice.ResourceEventHandler) (service.RepositoryFactory, error) {
	if db == nil {
		return nil, liberrors.Errorf("db is nil. err: %w", libdomain.ErrInvalidArgument)
	}

	return &repositoryFactory{
		dialect:              dialect,
		driverName:           driverName,
		db:                   db,
		location:             location,
		apppUserEventHandler: apppUserEventHandler,
	}, nil
}

func (f *repositoryFactory) NewOrganizationRepository(ctx context.Context) service.OrganizationRepository {
	return NewOrganizationRepository(ctx, f.db)
}

func (f *repositoryFactory) NewAppUserRepository(ctx context.Context) service.AppUserRepository {
	return NewAppUserRepository(ctx, f.dialect, f.db, f)
}

func (f *repositoryFactory) NewUserGroupRepository(ctx context.Context) service.UserGroupRepository {
	return NewUserGroupRepository(ctx, f.dialect, f.db)
}

// func (f *repositoryFactory) NewPairOfUserAndGroupRepository(ctx context.Context) service.PairOfUserAndGroupRepository {
// 	return NewPairOfUserAndGroupRepository(ctx, f.db, f)
// }

// func (f *repositoryFactory) NewRBACRepository(ctx context.Context) service.RBACRepository {
// 	return NewRBACRepository(ctx, f.db)
// }

func (f *repositoryFactory) NewAuthorizationManager(ctx context.Context) (service.AuthorizationManager, error) {
	return NewAuthorizationManager(ctx, f.dialect, f.db, f)
}

func (f *repositoryFactory) NewAppUserEventHandler(_ context.Context) mblibservice.ResourceEventHandler {
	return f.apppUserEventHandler
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
