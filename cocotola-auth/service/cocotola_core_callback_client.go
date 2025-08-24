package service

import (
	"context"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type CocotolaCoreCallbackClient interface {
	OnAddAppUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.AppUserID) error
}
