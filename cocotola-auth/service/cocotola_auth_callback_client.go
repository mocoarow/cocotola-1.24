package service

import (
	"context"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CocotolaAuthCallbackClient interface {
	OnAddUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID) error
}
