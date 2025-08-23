package service

import (
	"context"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
)

type SystemOwnerAction struct {
	rf                   RepositoryFactory
	rsrf                 mbuserservice.RepositoryFactory
	systemAdmin          *mbuserservice.SystemAdmin
	SystemOwner          *mbuserservice.SystemOwner
	Organization         *mbuserservice.Organization
	AuthorizationManager mbuserservice.AuthorizationManager
}

type SystemOwnerActionOption func(context.Context, *SystemOwnerAction) error

// func WithSystemAdmin() SystemAdminActionOption {
// 	return func(ctx context.Context, action *systemAdminAction) error {
// 		if err := action.initSystemAdmin(ctx); err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

func (a *SystemOwnerAction) initSystemOwnerByOrganizationID(ctx context.Context, organizationID *mbuserdomain.OrganizationID) error {
	if a.SystemOwner != nil {
		return nil
	}

	systemOwner, err := a.systemAdmin.FindSystemOwnerByOrganizationID(ctx, organizationID)
	if err != nil {
		return mbliberrors.Errorf("find system owner by organization id(%d): %w", organizationID.Int(), err)
	}
	a.SystemOwner = systemOwner
	return nil
}
func (a *SystemOwnerAction) initSystemOwnerByOrganizationName(ctx context.Context, organizationName string) error {
	if a.SystemOwner != nil {
		return nil
	}

	systemOwner, err := a.systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
	if err != nil {
		return mbliberrors.Errorf("find system owner by organization name %s: %w", organizationName, err)
	}
	a.SystemOwner = systemOwner
	return nil
}

func (a *SystemOwnerAction) initOrganizationByOrganizationID(ctx context.Context, organizationID *mbuserdomain.OrganizationID) error {
	if a.Organization != nil {
		return nil
	}

	organization, err := a.systemAdmin.FindOrganizationByID(ctx, organizationID)
	if err != nil {
		return mbliberrors.Errorf("find organization by id(%d): %w", organizationID.Int(), err)
	}
	a.Organization = organization
	return nil
}

func (a *SystemOwnerAction) initOrganizationByOrganizationName(ctx context.Context, organizationName string) error {
	if a.Organization != nil {
		return nil
	}

	organization, err := a.systemAdmin.FindOrganizationByName(ctx, organizationName)
	if err != nil {
		return mbliberrors.Errorf("find organization by name %s: %w", organizationName, err)
	}
	a.Organization = organization
	return nil
}
func WithOrganizationByID(organizationID *mbuserdomain.OrganizationID) SystemOwnerActionOption {
	return func(ctx context.Context, action *SystemOwnerAction) error {
		if err := action.initSystemOwnerByOrganizationID(ctx, organizationID); err != nil {
			return err
		}
		if err := action.initOrganizationByOrganizationID(ctx, organizationID); err != nil {
			return err
		}
		return nil
	}
}

func WithOrganizationByName(organizationName string) SystemOwnerActionOption {
	return func(ctx context.Context, action *SystemOwnerAction) error {
		if err := action.initSystemOwnerByOrganizationName(ctx, organizationName); err != nil {
			return err
		}
		if err := action.initOrganizationByOrganizationName(ctx, organizationName); err != nil {
			return err
		}
		return nil
	}
}

func WithAuthorizationManager() SystemOwnerActionOption {
	return func(ctx context.Context, action *SystemOwnerAction) error {
		authorizationManager, err := action.rsrf.NewAuthorizationManager(ctx)
		if err != nil {
			return mbliberrors.Errorf("failed to NewAuthorizationManager. err: %w", err)
		}
		action.AuthorizationManager = authorizationManager
		return nil
	}
}
func NewSystemOwnerAction(ctx context.Context, systemToken libdomain.SystemToken, rf RepositoryFactory, options ...SystemOwnerActionOption) (*SystemOwnerAction, error) {
	systemAdminAction, err := NewSystemAdminAction(ctx, systemToken, rf)
	if err != nil {
		return nil, err
	}
	action := SystemOwnerAction{}
	action.rf = rf
	action.rsrf = systemAdminAction.rsrf
	action.systemAdmin = systemAdminAction.SystemAdmin
	for _, option := range options {
		if err := option(ctx, &action); err != nil {
			return nil, err
		}
	}
	if action.SystemOwner == nil {
		return nil, mbliberrors.Errorf("system owner is not initialized: %w", err)
	}
	if action.Organization == nil {
		return nil, mbliberrors.Errorf("organization is not initialized: %w", err)
	}
	return &action, nil
}
