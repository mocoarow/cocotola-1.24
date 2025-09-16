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

type appUser struct {
	appUserID      *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

func (m *appUser) UserID() *mbuserdomain.UserID {
	return m.appUserID
}
func (m *appUser) OrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *appUser) Username() string {
	return m.username
}
func (m *appUser) LoginID() string {
	return m.loginID
}

func organizationID(t *testing.T, organizationID int) *mbuserdomain.OrganizationID {
	t.Helper()
	id, err := mbuserdomain.NewOrganizationID(organizationID)
	require.NoError(t, err)
	return id
}

func appUserID(t *testing.T, appUserID int) *mbuserdomain.UserID {
	t.Helper()
	id, err := mbuserdomain.NewUserID(appUserID)
	require.NoError(t, err)
	return id
}
