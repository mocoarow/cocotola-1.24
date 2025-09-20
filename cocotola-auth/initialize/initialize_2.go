package initialize

import (
	"context"
	"time"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
	"gorm.io/gorm"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type ParentAndChildLink struct {
	Parent mbuserdomain.RBACObject
	Child  mbuserdomain.RBACObject
}

func Initialize2(ctx context.Context, systemToken libdomain.SystemToken, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, organizationID *mbuserdomain.OrganizationID, parentAndChildLink []*ParentAndChildLink) error {
	rff := func(ctx context.Context, db *gorm.DB) (mbuserservice.RepositoryFactory, error) {
		resouceEventHandlers := map[mbuserdomain.ResourceKey]mblibservice.ResourceEventHandler{}
		return mbusergateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, resouceEventHandlers)
	}

	txManager := initMBTransactionManager(db, rff)

	fn := func(rf mbuserservice.RepositoryFactory) error {
		systemOwnerAction, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
			service.WithOrganizationByID(organizationID),
			service.WithAuthorizationManager(),
		)
		if err != nil {
			return mbliberrors.Errorf("new system owner action: %w", err)
		}

		// mbrf, err := rf.NewMoonBeamRepositoryFactory(ctx)
		// if err != nil {
		// 	return mbliberrors.Errorf("NewMoonBeamRepositoryFactory: %w", err)
		// }
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("new authorization manager: %w", err)
		}
		for _, po := range parentAndChildLink {
			if err := authorizationManager.AddObjectToObject(ctx, systemOwnerAction.SystemOwner, po.Child, po.Parent); err != nil {
				return mbliberrors.Errorf("AddObjectToObject: %w", err)
			}
		}

		return nil
	}
	if err := mblibservice.Do0(ctx, txManager, fn); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
