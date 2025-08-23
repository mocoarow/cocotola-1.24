package service

import (
	"context"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
)

type systemAdminAction struct {
	rf          RepositoryFactory
	rsrf        mbuserservice.RepositoryFactory
	SystemAdmin *mbuserservice.SystemAdmin
}

type SystemAdminActionOption func(context.Context, *systemAdminAction) error

func (a *systemAdminAction) initRsrf(ctx context.Context) error {
	if a.rsrf != nil {
		return nil
	}

	rsrf, err := a.rf.NewMoonBeamRepositoryFactory(ctx)
	if err != nil {
		return err
	}
	a.rsrf = rsrf
	return nil
}

func (a *systemAdminAction) initSystemAdmin(ctx context.Context) error {
	if a.SystemAdmin != nil {
		return nil
	}
	if err := a.initRsrf(ctx); err != nil {
		return err
	}

	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, a.rsrf)
	if err != nil {
		return err
	}
	a.SystemAdmin = systemAdmin
	return nil
}

// func WithSystemAdmin() SystemAdminActionOption {
// 	return func(ctx context.Context, action *systemAdminAction) error {
// 		if err := action.initSystemAdmin(ctx); err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

func NewSystemAdminAction(ctx context.Context, systemToken libdomain.SystemToken, rf RepositoryFactory, options ...SystemAdminActionOption) (*systemAdminAction, error) {
	action := systemAdminAction{}
	action.rf = rf
	if err := action.initSystemAdmin(ctx); err != nil {
		return nil, err
	}
	// for _, option := range options {
	// 	if err := option(ctx, &action); err != nil {
	// 		return nil, err
	// 	}
	// }
	return &action, nil
}
