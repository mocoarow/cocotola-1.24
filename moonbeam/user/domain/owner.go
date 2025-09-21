package domain

type OwnerModel struct {
	*User
}

func NewOwnerModel(user *User) (*OwnerModel, error) {
	return &OwnerModel{
		User: user,
	}, nil
}

func (m *OwnerModel) IsOwner() bool {
	return true
}
func (m *OwnerModel) GetOrganizationID() *OrganizationID {
	return m.User.GetOrganizationID()
}
func (m *OwnerModel) GetUserID() *UserID {
	return m.User.GetUserID()
}
