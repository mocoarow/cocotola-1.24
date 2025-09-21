package gateway

import (
	"context"
	"fmt"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type SpaceQueryUseCase struct {
	mbrf mbuserservice.RepositoryFactory
}

func NewSpaceQueryUsecase(mbrf mbuserservice.RepositoryFactory) *SpaceQueryUseCase {
	return &SpaceQueryUseCase{mbrf: mbrf}
}

func (u *SpaceQueryUseCase) FindPublicSpaces(ctx context.Context, operator mbuserdomain.UserInterface) ([]*mbuserdomain.SpaceModel, error) {
	_, span := tracer.Start(ctx, "SpaceQueryUseCase.FindPublicSpaces")
	defer span.End()

	spaceRepo := u.mbrf.NewSpaceRepository(ctx)
	spaces, err := spaceRepo.FindPublicSpaces(ctx, operator)
	if err != nil {
		return nil, fmt.Errorf("spaceRepo.FindPublicSpaces. err: %w", err)
	}

	return spaces, nil
}
