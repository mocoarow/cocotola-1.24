package gateway_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type organization struct {
	organizationID *mbuserdomain.OrganizationID
	name           string
}

var _ service.OrganizationInterface = (*organization)(nil)

func (m *organization) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}

func (m *organization) GetName() string {
	return m.name
}

type user struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

var _ service.UserInterface = (*user)(nil)

func (m *user) GetUserID() *mbuserdomain.UserID {
	return m.userID
}
func (m *user) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *user) GetUsername() string {
	return m.username
}
func (m *user) GetLoginID() string {
	return m.loginID
}

func organizationID(t *testing.T, organizationID int) *mbuserdomain.OrganizationID {
	t.Helper()
	id, err := mbuserdomain.NewOrganizationID(organizationID)
	require.NoError(t, err)
	return id
}

func userID(t *testing.T, userID int) *mbuserdomain.UserID {
	t.Helper()
	id, err := mbuserdomain.NewUserID(userID)
	require.NoError(t, err)
	return id
}
