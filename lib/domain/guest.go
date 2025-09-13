package domain

func NewGuestLoginID(organizationName string) string {
	return "guest@@" + organizationName
}
func NewGuestUserName(organizationName string) string {
	return "Guest(" + organizationName + ")"
}
func IsGuestLoginID(loginID string) bool {
	return len(loginID) > 7 && loginID[:7] == "guest@@"
}
