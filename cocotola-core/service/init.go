package service

import (
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var (
	TemplateIDEnglishBlank *domain.TemplateID
)

func init() {
	templateIDEnglishBlank, err := domain.NewTemplateID(1)
	if err != nil {
		panic(err)
	}
	TemplateIDEnglishBlank = templateIDEnglishBlank
}

type RoleUserInterface interface {
	GetUserID() *mbuserdomain.UserID
	GetOrganizationID() *mbuserdomain.OrganizationID
	GetRole() string
	GetBearerToken() string
	// LoginID() string
	// Username() string
}
