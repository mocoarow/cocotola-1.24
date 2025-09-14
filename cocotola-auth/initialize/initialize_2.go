package initialize

import (
	"context"
	"time"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"gorm.io/gorm"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func Initialize2(ctx context.Context, systemToken libdomain.SystemToken, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, organizationID *mbuserdomain.OrganizationID, parentObject mbuserdomain.RBACObject, childObjects []mbuserdomain.RBACObject) error {
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		resouceEventHandlers := map[mbuserdomain.ResourceKey]mblibservice.ResourceEventHandler{}
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, resouceEventHandlers)
	}

	txManager := initTransactionManager(db, rff)

	fn := func(rf service.RepositoryFactory) error {
		systemOwnerAction, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
			service.WithOrganizationByID(organizationID),
			service.WithAuthorizationManager(),
		)
		if err != nil {
			return mbliberrors.Errorf("new system owner action: %w", err)
		}

		mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		}
		authorizationManager, err := mbrf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("new authorization manager: %w", err)
		}
		for _, child := range childObjects {
			authorizationManager.AddObjectToObject(ctx, systemOwnerAction.SystemOwner, child, parentObject)
		}

		return nil
	}
	if err := mblibservice.Do0(ctx, txManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
