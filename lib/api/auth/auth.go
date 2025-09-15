package auth

type AppUserInfoResponse struct {
	AppUserID      int      `json:"appUserId"`
	OrganizationID int      `json:"organizationId"`
	LoginID        string   `json:"loginId"`
	Username       string   `json:"username"`
	UserGroups     []string `json:"groups"`
}

type PasswordAuthParameter struct {
	LoginID          string `json:"loginId" binding:"required"`
	Password         string `json:"password" binding:"required"`
	OrganizationName string `json:"organizationName" binding:"required"`
}
type GuestAuthRequest struct {
	OrganizationName string `json:"organizationName" binding:"required"`
}

type AuthResponse struct { //nolint:revive
	AccessToken  *string `json:"accessToken"`
	RefreshToken *string `json:"refreshToken"`
}

type RefreshTokenParameter struct {
	RefreshToken string `json:"refreshToken"`
}
