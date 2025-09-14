package initialize

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func initOrganization(ctx context.Context, systemToken libdomain.SystemToken, _, nonTxManager service.TransactionManager, organizationName, loginID, password string) (*mbuserdomain.OrganizationID, *mbuserdomain.SpaceID, error) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp1"))

	fn := func(rf service.RepositoryFactory) (*mbuserdomain.OrganizationID, *mbuserdomain.SpaceID, error) {
		systemAdminAction, err := service.NewSystemAdminAction(ctx, systemToken, rf)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("new system admin action: %w", err)
		}
		// 1. check whether the organization already exists
		organization, err := systemAdminAction.SystemAdmin.FindOrganizationByName(ctx, organizationName)
		if err == nil {
			logger.InfoContext(ctx, fmt.Sprintf("organization: %d", organization.OrganizationID().Int()))

			systemOwnerAction, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
				service.WithOrganizationByName(organizationName),
				service.WithAuthorizationManager(),
			)
			if err != nil {
				return nil, nil, mbliberrors.Errorf("new system owner action: %w", err)
			}

			publicDefaultSpace, err := systemOwnerAction.SystemOwner.GetPublidDefaultSpace(ctx)
			if err != nil {
				return nil, nil, mbliberrors.Errorf("GetPublidDefaultSpace: %w", err)
			}
			return organization.OrganizationID(), publicDefaultSpace.SpaceID, nil
		} else if !errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
			return nil, nil, mbliberrors.Errorf("find organization by name(%s): %w", organizationName, err)
		}

		// 2. add organization
		organizationID, err := addOrganization(ctx, systemAdminAction, organizationName, loginID, password)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("add organization: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", organizationID.Int()))

		systemOwnerAction, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
			service.WithOrganizationByName(organizationName),
			service.WithAuthorizationManager(),
		)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("new system owner action: %w", err)
		}

		// 3. add policy to "first-owner" user

		firstOwner, err := systemOwnerAction.SystemOwner.FindAppUserByLoginID(ctx, loginID)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("FindAppUserByLoginID: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("firstOwner: %d", firstOwner.AppUserID().Int()))

		// first owner can create app users
		subject := firstOwner.AppUserID().GetRBACSubject()
		action := mbuserdomain.NewRBACAction("CreateAppUser")
		object := mbuserdomain.NewRBACObject("*")
		effect := mbuserservice.RBACAllowEffect

		if err := systemOwnerAction.AuthorizationManager.AddPolicyToUserBySystemOwner(ctx, systemOwnerAction.SystemOwner, subject, action, object, effect); err != nil {
			return nil, nil, mbliberrors.Errorf("AddPolicyToUserBySystemOwner: %w", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", organizationID.Int()))

		publicDefaultSpace, err := systemOwnerAction.SystemOwner.GetPublidDefaultSpace(ctx)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("GetPublidDefaultSpace: %w", err)
		}
		logger.InfoContext(ctx, fmt.Sprintf("publicDefaultSpace: %d", publicDefaultSpace.SpaceID.Int()))

		return organizationID, publicDefaultSpace.SpaceID, nil
	}

	organizationID, publicDefaultSpaceID, err := mblibservice.Do2(ctx, nonTxManager, fn)
	if err != nil {
		return nil, nil, err //nolint:wrapcheck
	}

	return organizationID, publicDefaultSpaceID, nil
}

func addOrganization(ctx context.Context, systemAdminAction *service.SystemAdminAction, organizationName, loginID, password string) (*mbuserdomain.OrganizationID, error) {
	firstOwnerAddParam, err := mbuserservice.NewAppUserAddParameter(loginID, "Owner(cocotola)", password, "", "", "", "")
	if err != nil {
		return nil, mbliberrors.Errorf("new AppUserAddParameter: %w", err)
	}

	organizationAddParameter, err := mbuserservice.NewOrganizationAddParameter(organizationName, firstOwnerAddParam)
	if err != nil {
		return nil, mbliberrors.Errorf("new OrganizationAddParameter: %w", err)
	}

	organizationID, err := systemAdminAction.SystemAdmin.AddOrganization(ctx, organizationAddParameter)
	if err != nil {
		return nil, mbliberrors.Errorf("add organization: %w", err)
	}

	return organizationID, nil
}
