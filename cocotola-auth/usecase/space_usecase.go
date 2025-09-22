package usecase

import (
	"context"
	"fmt"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type SpaceUsecase struct {
	mbrf mbuserservice.RepositoryFactory
}

func NewSpaceUsecase(mbrf mbuserservice.RepositoryFactory) *SpaceUsecase {
	return &SpaceUsecase{mbrf: mbrf}
}

func (u *SpaceUsecase) FindPublicSpaces(ctx context.Context, operator mbuserdomain.UserInterface) ([]*mbuserdomain.Space, error) {
	_, span := tracer.Start(ctx, "SpaceUsecase.FindPublicSpaces")
	defer span.End()

	spaceRepo := u.mbrf.NewSpaceRepository(ctx)
	spaces, err := spaceRepo.FindPublicSpaces(ctx, operator)
	if err != nil {
		return nil, fmt.Errorf("spaceRepo.FindPublicSpaces. err: %w", err)
	}

	return spaces, nil
}
