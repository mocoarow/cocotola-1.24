package service

import (
	"context"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CocotolaCoreCallbackClient interface {
	OnAddAppUserSpace(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.AppUserID, spaceID *mbuserdomain.SpaceID) error
}
