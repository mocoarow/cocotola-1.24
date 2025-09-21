package core

type GetMyProfileResponseSpace struct {
	SpaceID      int `json:"spaceId"`
	RootFolderID int `json:"rootFolderId"`
}
type GetMyProfileResponse struct {
	PrivateSpace GetMyProfileResponseSpace `json:"privateSpace"`
}
