package domain

type OwnerModel struct {
	*UserModel
}

func NewOwnerModel(user *UserModel) (*OwnerModel, error) {
	return &OwnerModel{
		UserModel: user,
	}, nil
}

func (m *OwnerModel) IsOwner() bool {
	return true
}
func (m *OwnerModel) GetOrganizationID() *OrganizationID {
	return m.UserModel.GetOrganizationID()
}
func (m *OwnerModel) GetUserID() *UserID {
	return m.UserModel.GetUserID()
}
