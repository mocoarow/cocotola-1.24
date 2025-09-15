package auth

type AppUserAddRequest struct {
	LoginID  string `json:"loginId" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
