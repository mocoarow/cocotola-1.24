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
	ctx, span := tracer.Start(ctx, "Initialize2")
	defer span.End()

	mbrff := func(ctx context.Context, db *gorm.DB) (mbuserservice.RepositoryFactory, error) {
		resouceEventHandlers := map[mbuserdomain.ResourceKey]mblibservice.ResourceEventHandler{}
		return mbusergateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC, resouceEventHandlers)
	}
	mbrf, err := mbrff(ctx, db)
	if err != nil {
		return mbliberrors.Errorf("rff: %w", err)
	}

	txManager := initMBTransactionManager(db, mbrff)
	mbNonTxManager := initMBNonTransactionManager(mbrf)

	sysAdmin := service.NewSystemAdmin(systemToken)

	sysOwner, err := findSystemOwnerByOrganizationID(ctx, sysAdmin, mbNonTxManager, organizationID)
	if err != nil {
		return mbliberrors.Errorf("findSystemOwnerByOrganizationID: %w", err)
	}

	fn := func(rf mbuserservice.RepositoryFactory) error {
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("new authorization manager: %w", err)
		}
		for _, po := range parentAndChildLink {
			if err := authorizationManager.AddObjectToObject(ctx, sysOwner, po.Child, po.Parent); err != nil {
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
