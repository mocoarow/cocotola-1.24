package auth

type CallbackOnAddUserRequest struct {
	OrganizationID int `json:"organizationId" binding:"required,gte=1"`
	UserID         int `json:"userId" binding:"required,gte=1"`
}
