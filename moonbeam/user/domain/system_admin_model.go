package domain

type SystemAdminModel struct {
	UserID *UserID
}

func NewSystemAdminModel() *SystemAdminModel {
	return &SystemAdminModel{
		UserID: SystemAdminID,
	}
}

func (m *SystemAdminModel) IsSystemAdmin() bool {
	return true
}
func (m *SystemAdminModel) GetUserID() *UserID {
	return m.UserID
}
