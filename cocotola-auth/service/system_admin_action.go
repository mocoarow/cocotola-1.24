package service

import (
	"context"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"
)

type SystemAdminAction struct {
	rf          RepositoryFactory
	mbrf        mbuserservice.RepositoryFactory
	SystemAdmin *mbuserservice.SystemAdmin
}

func (a *SystemAdminAction) initMbrf(ctx context.Context) error {
	if a.mbrf != nil {
		return nil
	}

	mbrf, err := a.rf.NewMoonBeamRepositoryFactory(ctx)
	if err != nil {
		return err
	}
	a.mbrf = mbrf
	return nil
}

func (a *SystemAdminAction) initSystemAdmin(ctx context.Context) error {
	if a.SystemAdmin != nil {
		return nil
	}
	if err := a.initMbrf(ctx); err != nil {
		return err
	}

	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, a.mbrf)
	if err != nil {
		return err
	}
	a.SystemAdmin = systemAdmin
	return nil
}

func NewSystemAdminAction(ctx context.Context, systemToken libdomain.SystemToken, rf RepositoryFactory) (*SystemAdminAction, error) {
	action := SystemAdminAction{}
	action.rf = rf
	if err := action.initSystemAdmin(ctx); err != nil {
		return nil, err
	}
	return &action, nil
}
