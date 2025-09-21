package service

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var ErrSpaceAlreadyExists = errors.New("space already exists")
var ErrSpaceNotFound = errors.New("space not found")

type AddSpaceParameter struct {
	Key       string
	Name      string
	SpaceType string
}

// type SpaceUpdateParameter struct {
// 	Name     string
// 	IsPublic bool
// }

type SpaceRepository interface {
	AddSpace(ctx context.Context, operator domain.UserInterface, param *AddSpaceParameter) (*domain.SpaceID, error)

	// UpdateSpace(ctx context.Context, operator UserInterface, deckID *domain.SpaceID, version int, param *SpaceUpdateParameter) error

	FindPublicSpaces(ctx context.Context, operator domain.UserInterface) ([]*domain.SpaceModel, error)

	FindPublicSpaceByKey(ctx context.Context, operator domain.UserInterface, key string) (*domain.SpaceModel, error)

	GetSpaceByID(ctx context.Context, operator domain.UserInterface, deckID *domain.SpaceID) (*domain.SpaceModel, error)
}
