package domain

// type OwnerModel interface {
// 	AppUserModel
// 	IsOwnerModel() bool
// }

type OwnerModel struct {
	*AppUserModel
}

func NewOwnerModel(appUser *AppUserModel) (*OwnerModel, error) {
	return &OwnerModel{
		AppUserModel: appUser,
	}, nil
}

// func (m *ownerModel) IsOwnerModel() bool {
// 	return true
// }
