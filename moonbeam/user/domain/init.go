package domain

var SystemAdminID *AppUserID

func init() {
	systemAdminID := 1
	systemAdminIDTmp, err := NewAppUserID(systemAdminID)
	if err != nil {
		panic(err)
	}
	SystemAdminID = systemAdminIDTmp
}
