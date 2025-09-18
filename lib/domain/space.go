package domain

import "strconv"

func NewPersonalSpaceKey(userID int) string {
	return "__personal_space@@" + strconv.Itoa(userID)
}
func NewPersonalSpaceName(loginID string) string {
	return "Personal Space(" + loginID + ")"
}
