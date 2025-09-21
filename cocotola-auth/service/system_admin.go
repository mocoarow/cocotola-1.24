package service

import (
	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type systemAdmin struct {
	*mbuserdomain.SystemAdmin
}

func (m *systemAdmin) GetUserID() *mbuserdomain.UserID {
	return m.UserID
}
func (m *systemAdmin) IsSystemAdmin() bool {
	return true
}

func NewSystemAdmin(systemToken libdomain.SystemToken) mbuserdomain.SystemAdminInterface {
	return &systemAdmin{
		mbuserdomain.NewSystemAdmin(),
	}
}
