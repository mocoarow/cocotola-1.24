package service

import (
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
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

// func NewSystemOwnerLoginID(organizationID *domain.OrganizationID) string {
// 	return systemOwnerLoginID + strconv.Itoa(organizationID.Int())
// }

var RBACSetAction = domain.NewRBACAction("Set")
var RBACUnsetAction = domain.NewRBACAction("Unset")

var RBACAllowEffect = domain.NewRBACEffect("allow")
var RBACDenyEffect = domain.NewRBACEffect("deny")
