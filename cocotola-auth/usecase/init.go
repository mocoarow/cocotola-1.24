package usecase

import (
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/mocoarow/cocotola-auth/usecase")
)

type usecaseOrganization struct {
	organizationID *mbuserdomain.OrganizationID
	name           string
}

func (m *usecaseOrganization) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *usecaseOrganization) GetName() string {
	return m.name
}

type usecaseUser struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

func (m *usecaseUser) GetUserID() *mbuserdomain.UserID {
	return m.userID
}
func (m *usecaseUser) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *usecaseUser) GetUsername() string {
	return m.username
}
func (m *usecaseUser) GetLoginID() string {
	return m.loginID
}
