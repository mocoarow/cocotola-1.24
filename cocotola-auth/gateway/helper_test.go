package gateway_test

import (
	"testing"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/stretchr/testify/require"
)

type organization struct {
	organizationID *mbuserdomain.OrganizationID
	name           string
}

func (m *organization) OrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *organization) Name() string {
	return m.name
}

type user struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

func (m *user) UserID() *mbuserdomain.UserID {
	return m.userID
}
func (m *user) OrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *user) Username() string {
	return m.username
}
func (m *user) LoginID() string {
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
