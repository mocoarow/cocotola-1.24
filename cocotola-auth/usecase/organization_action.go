package usecase

import (
	"context"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type organizationAction struct {
	rf               service.RepositoryFactory
	rsrf             mbuserservice.RepositoryFactory
	organizationRepo mbuserservice.OrganizationRepository
	appUserRepo      mbuserservice.AppUserRepository
	organization     *mbuserservice.Organization
	systemAdmin      *mbuserservice.SystemAdmin
	systemOwner      *mbuserservice.SystemOwner
}

type OrganizationActionOption func(context.Context, *organizationAction) error

func (a *organizationAction) initRsrf(ctx context.Context) error {
	if a.rsrf != nil {
		return nil
	}

	rsrf, err := a.rf.NewmoonbeamRepositoryFactory(ctx)
	if err != nil {
		return err
	}
	a.rsrf = rsrf
	return nil
}

func (a *organizationAction) initSystemAdmin(ctx context.Context) error {
	if a.systemAdmin != nil {
		return nil
	}

	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, a.rsrf)
	if err != nil {
		return err
	}
	a.systemAdmin = systemAdmin
	return nil
}

func (a *organizationAction) initSystemOwnerByOrganizationName(ctx context.Context, organizationName string) error {
	if a.systemOwner != nil {
		return nil
	}

	if err := a.initSystemAdmin(ctx); err != nil {
		return err
	}
	systemOwner, err := a.systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
	if err != nil {
		return err
	}
	a.systemOwner = systemOwner
	return nil
}

func WithSystemAdmin() OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initSystemAdmin(ctx); err != nil {
			return err
		}
		return nil
	}
}

func WithOrganizationRepository() OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initRsrf(ctx); err != nil {
			return err
		}
		action.organizationRepo = action.rsrf.NewOrganizationRepository(ctx)
		return nil
	}
}

func WithAppUserRepository() OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initRsrf(ctx); err != nil {
			return err
		}
		action.appUserRepo = action.rsrf.NewAppUserRepository(ctx)
		return nil
	}
}

func WithOrganization(organizationName string) OrganizationActionOption {
	return func(ctx context.Context, action *organizationAction) error {
		if err := action.initSystemOwnerByOrganizationName(ctx, organizationName); err != nil {
			return err
		}
		return nil
	}
}

func NewOrganizationAction(ctx context.Context, rf service.RepositoryFactory, options ...OrganizationActionOption) (*organizationAction, error) {
	action := organizationAction{}
	action.rf = rf
	for _, option := range options {
		if err := option(ctx, &action); err != nil {
			return nil, err
		}
	}
	return &action, nil
}
