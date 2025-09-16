package domain

type OwnerModel struct {
	*UserModel
}

func NewOwnerModel(appUser *UserModel) (*OwnerModel, error) {
	return &OwnerModel{
		UserModel: appUser,
	}, nil
}

// func (m *ownerModel) IsOwnerModel() bool {
// 	return true
// }
