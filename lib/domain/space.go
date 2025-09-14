package domain

func NewPrivateSpaceKey(loginID string) string {
	return "__private_space@@" + loginID
}
func NewPrivateSpaceName(loginID string) string {
	return "Private Space(" + loginID + ")"
}
