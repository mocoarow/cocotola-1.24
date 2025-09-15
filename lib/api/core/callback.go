package core

type CallbackOnAddAppUserSpaceRequest struct {
	OrganizationID int `json:"organizationId" binding:"required,gte=1"`
	AppUserID      int `json:"appUserId" binding:"required,gte=1"`
	SpaceID        int `json:"spaceId" binding:"required,gte=1"`
}
