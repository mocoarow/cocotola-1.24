package domain

type ResourceKey string

var (
	SystemAdminID *AppUserID

	ResourceAppUser = ResourceKey("app_user")
	RecourceSpace   = ResourceKey("space")
)

func init() {
	systemAdminID := 1
	systemAdminIDTmp, err := NewAppUserID(systemAdminID)
	if err != nil {
		panic(err)
	}
	SystemAdminID = systemAdminIDTmp
}
