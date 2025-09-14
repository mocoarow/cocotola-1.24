package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type PairOfUserAndSpaceRepository interface {
	AddPairOfUserAndSpace(ctx context.Context, operator AppUserInterface, appUserID *domain.AppUserID, spaceID *domain.SpaceID) error

	FindSpacesByUserID(ctx context.Context, operator AppUserInterface, appUserID *domain.AppUserID) ([]*Space, error)
}
