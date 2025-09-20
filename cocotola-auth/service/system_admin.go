package service

import (
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type systemAdmin struct {
	*mbuserdomain.SystemAdminModel
}

func (m *systemAdmin) GetUserID() *mbuserdomain.UserID {
	return m.UserID
}
func (m *systemAdmin) IsSystemAdmin() bool {
	return true
}

func NewSystemAdmin(systemToken libdomain.SystemToken) mbuserservice.SystemAdminInterface {
	return &systemAdmin{
		mbuserdomain.NewSystemAdminModel(),
	}
}
