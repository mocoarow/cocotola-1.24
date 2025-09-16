package auth

type CallbackOnAddUserRequest struct {
	OrganizationID int `json:"organizationId" binding:"required,gte=1"`
	UserID      int `json:"appUserId" binding:"required,gte=1"`
}
