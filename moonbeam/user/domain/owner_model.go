package domain

type OwnerModel struct {
	*UserModel
}

func NewOwnerModel(user *UserModel) (*OwnerModel, error) {
	return &OwnerModel{
		UserModel: user,
	}, nil
}

// func (m *ownerModel) IsOwnerModel() bool {
// 	return true
// }
