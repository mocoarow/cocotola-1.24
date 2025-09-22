package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type AddPersonalSpaceParameter struct {
	UserID  *domain.UserID
	KeyName string
	Name    string
}
type SpaceManager interface {
	AddPersonalSpace(ctx context.Context, operator domain.UserInterface, param *AddPersonalSpaceParameter) (*domain.SpaceID, error)
	AddPublicDefaultSpace(ctx context.Context, operator domain.UserInterface) (*domain.SpaceID, error)
	AddUserToSpace(ctx context.Context, operator domain.SystemOwnerInterface, userID domain.UserID, spaceID *domain.SpaceID) error
	GetPersonalSpace(ctx context.Context, operator domain.UserInterface) (*domain.Space, error)
}
