package core

type CallbackOnAddUserSpaceRequest struct {
	OrganizationID int `json:"organizationId" binding:"required,gte=1"`
	UserID      int `json:"userId" binding:"required,gte=1"`
	SpaceID        int `json:"spaceId" binding:"required,gte=1"`
}
