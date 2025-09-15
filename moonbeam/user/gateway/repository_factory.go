package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type repositoryFactory struct {
	dialect                libgateway.DialectRDBMS
	driverName             string
	db                     *gorm.DB
	location               *time.Location
	rersourceEventHandlers map[domain.ResourceKey]mblibservice.ResourceEventHandler
}

func NewRepositoryFactory(_ context.Context, dialect libgateway.DialectRDBMS, driverName string, db *gorm.DB, location *time.Location, rersourceEventHandlers map[domain.ResourceKey]mblibservice.ResourceEventHandler) (service.RepositoryFactory, error) {
	if db == nil {
		return nil, liberrors.Errorf("db is nil. err: %w", libdomain.ErrInvalidArgument)
	}

	return &repositoryFactory{
		dialect:                dialect,
		driverName:             driverName,
		db:                     db,
		location:               location,
		rersourceEventHandlers: rersourceEventHandlers,
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

func (f *repositoryFactory) NewSpaceRepository(ctx context.Context) service.SpaceRepository {
	return NewSpaceRepository(ctx, f.dialect, f.db)
}

// func (f *repositoryFactory) NewPairOfUserAndGroupRepository(ctx context.Context) service.PairOfUserAndGroupRepository {
// 	return NewPairOfUserAndGroupRepository(ctx, f.db, f)
// }

// func (f *repositoryFactory) NewRBACRepository(ctx context.Context) service.RBACRepository {
// 	return NewRBACRepository(ctx, f.db)
// }

func (f *repositoryFactory) NewSpaceManager(ctx context.Context) (service.SpaceManager, error) {
	return NewSpaceManager(ctx, f.dialect, f.db, f)
}

func (f *repositoryFactory) NewAuthorizationManager(ctx context.Context) (service.AuthorizationManager, error) {
	return NewAuthorizationManager(ctx, f.dialect, f.db, f)
}

func (f *repositoryFactory) NewAppUserEventHandler(_ context.Context) mblibservice.ResourceEventHandler {
	return f.rersourceEventHandlers[domain.ResourceAppUser]
}
func (f *repositoryFactory) NewSpaceEventHandler(_ context.Context) mblibservice.ResourceEventHandler {
	return f.rersourceEventHandlers[domain.RecourceSpace]
}

type RepositoryFactoryFunc func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error)
