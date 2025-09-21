package domain

type SystemAdmin struct {
	UserID *UserID
}

func NewSystemAdmin() *SystemAdmin {
	return &SystemAdmin{
		UserID: SystemAdminID,
	}
}

func (m *SystemAdmin) IsSystemAdmin() bool {
	return true
}
func (m *SystemAdmin) GetUserID() *UserID {
	return m.UserID
}
