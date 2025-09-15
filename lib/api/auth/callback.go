package auth

type CallbackOnAddAppUserRequest struct {
	OrganizationID int `json:"organizationId" binding:"required,gte=1"`
	AppUserID      int `json:"appUserId" binding:"required,gte=1"`
}
