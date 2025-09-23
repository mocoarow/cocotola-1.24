package initialize

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
	mbuserusecase "github.com/mocoarow/cocotola-1.24/moonbeam/user/usecase"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func initOrganization(ctx context.Context, systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, organizationName, loginID, password string) (*mbuserdomain.OrganizationID, *mbuserdomain.UserID, *mbuserdomain.SpaceID, error) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp1"))

	sysAdmin := service.NewSystemAdmin(systemToken)

	// 1. check whether the organization already exists
	{
		organization, sysOwner, publicDefaultSpace, found, err := findOrganizationAndSystemOwnerAndPublicDefaultSpace(ctx, sysAdmin, mbNonTxManager, organizationName)
		if err != nil {
			return nil, nil, nil, mbliberrors.Errorf("findOrganizationAndPublicDefaultSpace: %w", err)
		}
		if found {
			return organization.OrganizationID, sysOwner.UserID, publicDefaultSpace.SpaceID, nil
		}
	}

	// 2. add organization
	orgID2, err := addOrganization(ctx, sysAdmin, mbTxManager, mbNonTxManager, organizationName)
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("add organization: %w", err)
	}
	logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", orgID2.Int()))

	// 3. find system owner
	sysOwner, err := findSystemOwnerByOrganizationName(ctx, sysAdmin, mbNonTxManager, organizationName)
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("findSystemOwnerByOrganizationName: %w", err)
	}

	// 4. add first owner
	firstOwnerID, err := addFirstOwnerToOrganization(ctx, sysOwner, mbTxManager, mbNonTxManager, loginID, password)
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("add first owner: %w", err)
	}
	logger.InfoContext(ctx, fmt.Sprintf("firstOwnerID: %d", firstOwnerID.Int()))

	// 5. find public default space
	publicDefaultSpace2, err := findPublicSpaceByKey(ctx, sysOwner, mbNonTxManager, mbuserservice.PublicDefaultSpaceKey)
	if err != nil {
		return nil, nil, nil, mbliberrors.Errorf("find public default space by key(%s): %w", mbuserservice.PublicDefaultSpaceKey, err)
	}

	return orgID2, sysOwner.UserID, publicDefaultSpace2.SpaceID, nil
}

func findOrganizationAndSystemOwnerAndPublicDefaultSpace(ctx context.Context, systemAdmin mbuserdomain.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.Organization, *mbuserdomain.SystemOwner, *mbuserdomain.Space, bool, error) {
	organization, err := findOrganizationByName(ctx, systemAdmin, mbNonTxManager, organizationName)
	if err != nil {
		if !errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
			return nil, nil, nil, false, mbliberrors.Errorf("find organization by name: %w", err)
		}
		return nil, nil, nil, false, nil
	}

	sysOwner, err := findSystemOwnerByOrganizationName(ctx, systemAdmin, mbNonTxManager, organizationName)
	if err != nil {
		return nil, nil, nil, false, mbliberrors.Errorf("find system owner by organization name: %w", err)
	}

	publicDefaultSpace, err := findPublicSpaceByKey(ctx, sysOwner, mbNonTxManager, mbuserservice.PublicDefaultSpaceKey)
	if err != nil {
		if !errors.Is(err, mbuserservice.ErrSpaceNotFound) {
			return nil, nil, nil, false, mbliberrors.Errorf("find public default space by key: %w", err)
		}
		return nil, nil, nil, false, nil
	}

	return organization, sysOwner, publicDefaultSpace, true, nil
}

func addOrganization(ctx context.Context, operator mbuserdomain.SystemAdminInterface, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.OrganizationID, error) {
	command := mbuserusecase.NewAddOrganizationCommand(ctx, mbTxManager, mbNonTxManager)
	organizationID, err := command.Execute(ctx, operator, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("add organization: %w", err)
	}
	return organizationID, nil
}

func addFirstOwnerToOrganization(ctx context.Context, operator mbuserdomain.SystemOwnerInterface, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, loginID, password string) (*mbuserdomain.UserID, error) {
	firstOwnerAddParam, err := mbuserservice.NewAddUserParameter(loginID, "Owner(cocotola)", password, "", "", "", "")
	if err != nil {
		return nil, mbliberrors.Errorf("new UserAddParameter: %w", err)
	}
	addFirstOwnerCommand := mbuserusecase.NewAddFirstOwnerCommand(mbTxManager, mbNonTxManager)
	firstOwnerID, err := addFirstOwnerCommand.Execute(ctx, operator, firstOwnerAddParam)
	if err != nil {
		return nil, mbliberrors.Errorf("add first owner: %w", err)
	}
	return firstOwnerID, nil
}
