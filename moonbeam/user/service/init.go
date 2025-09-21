package service

import (
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var (
	SystemAdminID *userdomain.UserID
)

type OperatorInterface interface {
	GetUserID() *domain.UserID
	GetOrganizationID() *domain.OrganizationID
}
type SystemAdminInterface interface {
	GetUserID() *domain.UserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type UserInterface interface {
	GetUserID() *domain.UserID
	GetOrganizationID() *domain.OrganizationID
	// LoginID() string
	// Username() string
	// GetUserGroups() []domain.UserGroupModel
}

type OwnerModelInterface interface {
	UserInterface
	IsOwner() bool
	// GetUserGroups() []domain.UserGroupModel
}

type SystemOwnerInterface interface {
	OwnerModelInterface
	IsSystemOwner() bool
	// GetUserGroups() []domain.UserGroupModel
}

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

var RBACSetAction = domain.NewRBACAction("Set")
var RBACUnsetAction = domain.NewRBACAction("Unset")

var RBACAllowEffect = domain.NewRBACEffect("allow")
var RBACDenyEffect = domain.NewRBACEffect("deny")

func init() {
	systemAdminID, err := userdomain.NewUserID(1)
	if err != nil {
		panic(err)
	}
	SystemAdminID = systemAdminID
}
