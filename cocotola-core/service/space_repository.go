package service

import (
	"context"
	"errors"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

var ErrSpaceAlreadyExists = errors.New("space already exists")
var ErrSpaceNotFound = errors.New("space not found")

type SpaceAddParameter struct {
	SpaceID *domain.SpaceID
	Name    string
}

type SpaceUpdateParameter struct {
	Name string
}

type SpaceRepository interface {
	AddSpace(ctx context.Context, operator mbuserservice.OperatorInterface, param *SpaceAddParameter) (*domain.SpaceID, error)

	UpdateSpace(ctx context.Context, operator OperatorInterface, deckID *domain.SpaceID, version int, param *SpaceUpdateParameter) error

	FindSpaces(ctx context.Context, operator OperatorInterface) ([]*Space, error)

	GetSpaceByID(ctx context.Context, operator OperatorInterface, deckID *domain.SpaceID) (*Space, error)
}
