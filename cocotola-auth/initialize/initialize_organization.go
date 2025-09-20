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

type systemAdmin struct {
	*mbuserdomain.SystemAdminModel
}

func (m *systemAdmin) GetUserID() *mbuserdomain.UserID {
	return m.UserID
}
func (m *systemAdmin) IsSystemAdmin() bool {
	return true
}

func newSystemAdmin(ctx context.Context) *systemAdmin {
	return &systemAdmin{
		mbuserdomain.NewSystemAdminModel(),
	}
}

func initOrganization(ctx context.Context, systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, organizationName, loginID, password string) (*mbuserdomain.OrganizationID, *mbuserdomain.SpaceID, error) {
	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp1"))

	sysAdmin := newSystemAdmin(ctx)

	// 1. check whether the organization already exists
	{
		organization, publicDefaultSpace, found, err := findOrganizationAndPublicDefaultSpace(ctx, sysAdmin, mbNonTxManager, organizationName)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("findOrganizationAndPublicDefaultSpace: %w", err)
		}
		if found {
			return organization.OrganizationID, publicDefaultSpace.SpaceID, nil
		}
	}

	// 2. add organization
	orgID2, err := addOrganization(ctx, sysAdmin, mbTxManager, mbNonTxManager, organizationName)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("add organization: %w", err)
	}
	logger.InfoContext(ctx, fmt.Sprintf("organizationID: %d", orgID2.Int()))

	// 3. find system owner
	sysOwner, err := findSystemOwnerByOrganizationName(ctx, sysAdmin, mbNonTxManager, organizationName)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("findSystemOwnerByOrganizationName: %w", err)
	}

	// 4. add first owner
	firstOwnerID, err := addFirstOwnerToOrganization(ctx, sysOwner, mbTxManager, mbNonTxManager, loginID, password)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("add first owner: %w", err)
	}
	logger.InfoContext(ctx, fmt.Sprintf("firstOwnerID: %d", firstOwnerID.Int()))

	// 5. find public default space
	publicDefaultSpace2, err := findPublicSpaceByKey(ctx, sysOwner, mbNonTxManager, mbuserservice.PublicDefaultSpaceKey)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("find public default space by key(%s): %w", mbuserservice.PublicDefaultSpaceKey, err)
	}

	return orgID2, publicDefaultSpace2.SpaceID, nil
}

func findOrganizationAndPublicDefaultSpace(ctx context.Context, systemAdmin mbuserservice.SystemAdminInterface, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.OrganizationModel, *mbuserdomain.SpaceModel, bool, error) {
	organization, err := findOrganizationByName(ctx, systemAdmin, mbNonTxManager, organizationName)
	if err != nil {
		if !errors.Is(err, mbuserservice.ErrOrganizationNotFound) {
			return nil, nil, false, mbliberrors.Errorf("find organization by name: %w", err)
		}
		return nil, nil, false, nil
	}

	sysOwner, err := findSystemOwnerByOrganizationName(ctx, systemAdmin, mbNonTxManager, organizationName)
	if err != nil {
		return nil, nil, false, mbliberrors.Errorf("find system owner by organization name: %w", err)
	}

	publicDefaultSpace, err := findPublicSpaceByKey(ctx, sysOwner, mbNonTxManager, mbuserservice.PublicDefaultSpaceKey)
	if err != nil {
		if !errors.Is(err, mbuserservice.ErrSpaceNotFound) {
			return nil, nil, false, mbliberrors.Errorf("find public default space by key: %w", err)
		}
		return nil, nil, false, nil
	}

	return organization, publicDefaultSpace, true, nil
}

func addOrganization(ctx context.Context, operator mbuserservice.SystemAdminInterface, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, organizationName string) (*mbuserdomain.OrganizationID, error) {
	command, err := mbuserusecase.NewAddOrganizationCommand(ctx, mbTxManager, mbNonTxManager)
	if err != nil {
		return nil, mbliberrors.Errorf("new AddOrganizationCommand: %w", err)
	}
	organizationID, err := command.Execute(ctx, operator, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("add organization: %w", err)
	}
	return organizationID, nil
}

func addFirstOwnerToOrganization(ctx context.Context, operator mbuserservice.SystemOwnerInterface, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, loginID, password string) (*mbuserdomain.UserID, error) {
	firstOwnerAddParam, err := mbuserservice.NewUserAddParameter(loginID, "Owner(cocotola)", password, "", "", "", "")
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

func newSystemAdminAction(ctx context.Context, systemToken libdomain.SystemToken, rf mbuserservice.RepositoryFactory) *service.SystemAdminAction {
	systemAdminAction, err := service.NewSystemAdminAction(ctx, systemToken, rf)
	if err != nil {
		libdomain.CheckError(err)
	}
	return systemAdminAction
}

func newSystemOwnerAction(ctx context.Context, systemToken libdomain.SystemToken, rf mbuserservice.RepositoryFactory, organizationName string) *service.SystemOwnerAction {
	systemOwnerAction, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
		service.WithOrganizationByName(organizationName),
		service.WithAuthorizationManager(),
	)
	if err != nil {
		libdomain.CheckError(err)
	}
	return systemOwnerAction
}
