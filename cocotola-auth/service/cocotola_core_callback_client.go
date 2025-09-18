package service

import (
	"context"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CocotolaCoreCallbackClient interface {
	OnAddUserSpace(ctx context.Context, organizationID *mbuserdomain.OrganizationID, userID *mbuserdomain.UserID, spaceID *mbuserdomain.SpaceID) error
}
