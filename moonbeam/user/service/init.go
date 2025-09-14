package service

import (
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var (
	SystemAdminID *userdomain.AppUserID
)

const (
	SystemAdminLoginID = "__system_admin"
	SystemOwnerLoginID = "__system_owner"

	SystemOwnerGroupKey   = "__system_owner"
	OwnerGroupKey         = "__owner"
	PublicGroupKey        = "__public_group"
	PublicDefaultSpaceKey = "__public_default_space"

	SystemOwnerGroupName   = "System Owner"
	OwnerGroupName         = "Owner"
	PublicGroupName        = "Public Group"
	PublicDefaultSpaceName = "Public Default Space"
)

func init() {
	systemAdminID, err := userdomain.NewAppUserID(1)
	if err != nil {
		panic(err)
	}
	SystemAdminID = systemAdminID
}
