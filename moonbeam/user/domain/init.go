package domain

type ResourceKey string

//	type UserInterface interface {
//		GetUserID() *UserID
//		GetOrganizationID() *OrganizationID
//	}
type SystemAdminInterface interface {
	GetUserID() *UserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type UserInterface interface {
	GetUserID() *UserID
	GetOrganizationID() *OrganizationID
	// LoginID() string
	// Username() string
	// GetUserGroups() []domain.UserGroupModel
}

type OwnerInterface interface {
	UserInterface
	IsOwner() bool
	// GetUserGroups() []domain.UserGroupModel
}

type SystemOwnerInterface interface {
	OwnerInterface
	IsSystemOwner() bool
	// GetUserGroups() []domain.UserGroupModel
}

var (
	SystemAdminID *UserID

	ResourceUser  = ResourceKey("user")
	RecourceSpace = ResourceKey("space")
)

func init() {
	systemAdminID := 1
	systemAdminIDTmp, err := NewUserID(systemAdminID)
	if err != nil {
		panic(err)
	}
	SystemAdminID = systemAdminIDTmp
}
