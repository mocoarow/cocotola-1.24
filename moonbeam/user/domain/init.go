package domain

type ResourceKey string

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
