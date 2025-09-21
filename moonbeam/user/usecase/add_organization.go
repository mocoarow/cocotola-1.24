package usecase

import (
	"context"
	"fmt"
	"log/slog"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type AddOrganizationCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	logger       *slog.Logger
}

func NewAddOrganizationCommand(ctx context.Context, txManager, nonTxManager service.TransactionManager) *AddOrganizationCommand {
	return &AddOrganizationCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(liblog.LoggerNameKey, "AddOrganizationCommand")),
	}
}

func (u *AddOrganizationCommand) Execute(ctx context.Context, operator domain.SystemAdminInterface, organizationName string) (*domain.OrganizationID, error) {
	fn := func(rf service.RepositoryFactory) (*domain.OrganizationID, error) {
		orgRepo := rf.NewOrganizationRepository(ctx)
		userRepo := rf.NewUserRepository(ctx)
		userGroupRepo := rf.NewUserGroupRepository(ctx)
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return nil, liberrors.Errorf("failed to NewAuthorizationManager: %w", err)
		}
		spaceManager, err := rf.NewSpaceManager(ctx)
		if err != nil {
			return nil, liberrors.Errorf("NewSpaceManager: %w", err)
		}

		// 1. add organization
		organizationID, err := orgRepo.AddOrganization(ctx, operator, organizationName)
		if err != nil {
			return nil, liberrors.Errorf("AddOrganization: %w", err)
		}

		// 2. add "system-owner" user
		// 3. add policy to "system-owner" user
		systemOwnerID, err := u.addSystemOwnerToOrganization(ctx, operator, userRepo, authorizationManager, organizationID)
		if err != nil {
			return nil, liberrors.Errorf("addSystemOwnertoOrganization: %w", err)
		}
		u.logger.InfoContext(ctx, fmt.Sprintf("systemOwnerID: %d", systemOwnerID.Int()))

		systemOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, operator, organizationName)
		if err != nil {
			return nil, liberrors.Errorf("FindSystemOwnerByOrganizationName: %w", err)
		}

		// 4. add owner-group
		// 5. add policty to "owner" group
		if _, err := u.addOwnergroupToOrganization(ctx, systemOwner, userGroupRepo, authorizationManager, organizationID); err != nil {
			return nil, liberrors.Errorf("addOwnergroupToOrganization: %w", err)
		}

		// 7. add public-group
		if _, err := userGroupRepo.AddPublicGroup(ctx, systemOwner, organizationID); err != nil {
			return nil, liberrors.Errorf("AddOwnerGroup: %w", err)
		}

		// 9. add public default space
		if _, err := spaceManager.AddPublicDefaultSpace(ctx, systemOwner); err != nil {
			return nil, liberrors.Errorf("add public space(%s): %w", service.PublicDefaultSpaceKey, err)
		}
		return organizationID, nil
	}
	organizationID, err := libservice.Do1(ctx, u.txManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return organizationID, nil
}

type ActionObjectEffect struct {
	Action domain.RBACAction
	Object domain.RBACObject
	Effect domain.RBACEffect
}

func (u *AddOrganizationCommand) addSystemOwnerToOrganization(ctx context.Context, operator domain.SystemAdminInterface, userRepo service.UserRepository, authorizationManager service.AuthorizationManager, organizationID *domain.OrganizationID) (*domain.UserID, error) {
	systemOwnerID, err := userRepo.AddSystemOwner(ctx, operator, organizationID)
	if err != nil {
		return nil, liberrors.Errorf("AddSystemOwner: %w", err)
	}

	// 3. add policy to "system-owner" user
	rbacSystemOwner := systemOwnerID.GetRBACSubject()
	rbacAllUserRolesObject := domain.NewRBACAllUserRolesObjectFromOrganization(organizationID)
	for _, aoe := range []ActionObjectEffect{
		{ // "system-owner" user "can" "set" "all-user-roles"
			Action: service.RBACSetAction,
			Object: rbacAllUserRolesObject,
			Effect: service.RBACAllowEffect,
		},
		{ //"system-owner" user "can" "unset" "all-user-roles"
			Action: service.RBACUnsetAction,
			Object: rbacAllUserRolesObject,
			Effect: service.RBACAllowEffect,
		},
	} {
		if err := authorizationManager.AddPolicyToUserBySystemAdmin(ctx, operator, organizationID, rbacSystemOwner, aoe.Action, aoe.Object, aoe.Effect); err != nil {
			return nil, liberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
		}
	}

	return systemOwnerID, nil
}

func (u *AddOrganizationCommand) addOwnergroupToOrganization(ctx context.Context, operator domain.SystemOwnerInterface, userGroupRepo service.UserGroupRepository, authorizationManager service.AuthorizationManager, organizationID *domain.OrganizationID) (*domain.UserGroupID, error) {
	// 4. add owner-group
	ownerGroupID, err := userGroupRepo.AddOwnerGroup(ctx, operator, organizationID)
	if err != nil {
		return nil, liberrors.Errorf("AddOwnerGroup: %w", err)
	}

	// 5. add policty to "owner" group
	rbacOwnerGroup := domain.NewRBACRoleFromGroup(organizationID, ownerGroupID)
	rbacAllUserRolesObject := domain.NewRBACAllUserRolesObjectFromOrganization(organizationID)

	for _, aoe := range []ActionObjectEffect{
		{ // "owner" group "can" "set" "all-user-roles"
			Action: service.RBACSetAction,
			Object: rbacAllUserRolesObject,
			Effect: service.RBACAllowEffect,
		},
		{ // "owner" group "can" "unset" "all-user-roles"
			Action: service.RBACUnsetAction,
			Object: rbacAllUserRolesObject,
			Effect: service.RBACAllowEffect,
		},
	} {
		if err := authorizationManager.AddPolicyToUserBySystemOwner(ctx, operator, rbacOwnerGroup, aoe.Action, aoe.Object, aoe.Effect); err != nil {
			return nil, liberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
		}
	}
	return ownerGroupID, nil
}
