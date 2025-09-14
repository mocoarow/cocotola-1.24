package service

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var ErrSpaceAlreadyExists = errors.New("space already exists")
var ErrSpaceNotFound = errors.New("space not found")

type SpaceAddParameter struct {
	Key      string
	Name     string
	IsPublic bool
}

type SpaceUpdateParameter struct {
	Name     string
	IsPublic bool
}

type SpaceRepository interface {
	AddSpace(ctx context.Context, operator OperatorInterface, param *SpaceAddParameter) (*domain.SpaceID, error)

	UpdateSpace(ctx context.Context, operator OperatorInterface, deckID *domain.SpaceID, version int, param *SpaceUpdateParameter) error

	FindPublicSpaces(ctx context.Context, operator OperatorInterface) ([]*Space, error)

	FindPublicSpaceByKey(ctx context.Context, key string) (*Space, error)

	GetSpaceByID(ctx context.Context, operator OperatorInterface, deckID *domain.SpaceID) (*Space, error)
}
