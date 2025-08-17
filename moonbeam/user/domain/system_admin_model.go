package domain

type SystemAdminModel struct {
	AppUserID *AppUserID
}

func NewSystemAdminModel() *SystemAdminModel {
	return &SystemAdminModel{
		AppUserID: SystemAdminID,
	}
}

// func (s *systemAdminModel) IsSystemAdminModel() bool {
// 	return true
// }
