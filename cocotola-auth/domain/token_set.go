package domain

type AuthTokenSet struct {
	AccessToken  string
	RefreshToken string
}

type UserInfo struct {
	Email string
	Name  string
}
