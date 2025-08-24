package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type PairOfUserAndSpaceRepository interface {
	AddPairOfUserAndSpace(ctx context.Context, operator mbuserservice.AppUserInterface, appUserID *mbuserdomain.AppUserID, spaceID *domain.SpaceID) error

	FindSpacesByUserID(ctx context.Context, operator OperatorInterface, appUserID *mbuserdomain.AppUserID) ([]*Space, error)
}
