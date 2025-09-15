package service

import (
	"context"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type AddPersonalSpaceParameter struct {
	AppUserID *domain.AppUserID
	KeyName   string
	Name      string
}
type SpaceManager interface {
	AddPersonalSpace(ctx context.Context, operator OperatorInterface, param *AddPersonalSpaceParameter) (*domain.SpaceID, error)
	AddPublicDefaultSpace(ctx context.Context, operator OperatorInterface) (*domain.SpaceID, error)
	AddUserToSpace(ctx context.Context, operator SystemOwnerInterface, appuUserID domain.AppUserID, spaceID *domain.SpaceID) error
	GetPersonalSpace(ctx context.Context, operator OperatorInterface) (*Space, error)
}
