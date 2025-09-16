package domain

type SystemAdminModel struct {
	UserID *UserID
}

func NewSystemAdminModel() *SystemAdminModel {
	return &SystemAdminModel{
		UserID: SystemAdminID,
	}
}

// func (s *systemAdminModel) IsSystemAdminModel() bool {
// 	return true
// }
