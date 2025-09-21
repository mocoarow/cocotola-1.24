package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type PairOfUserAndSpaceRepository interface {
	AddPairOfUserAndSpace(ctx context.Context, operator domain.UserInterface, userID *domain.UserID, spaceID *domain.SpaceID) error

	FindMySpaces(ctx context.Context, operator domain.UserInterface) ([]*domain.SpaceModel, error)
}
