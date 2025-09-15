package domain

import "strconv"

func NewPersonalSpaceKey(appUserID int) string {
	return "__personal_space@@" + strconv.Itoa(appUserID)
}
func NewPersonalSpaceName(loginID string) string {
	return "Personal Space(" + loginID + ")"
}
