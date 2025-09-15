package usecase

import (
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type Operator struct {
	appUserID      *mbuserdomain.AppUserID
	organizationID *mbuserdomain.OrganizationID
}

func (o *Operator) GetAppUserID() *mbuserdomain.AppUserID {
	return o.appUserID
}
func (o *Operator) GetOrganizationID() *mbuserdomain.OrganizationID {
	return o.organizationID
}
