//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewOwner_shouldReturnOwner_whenValidUserIsSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "owner@example.com"
	username := "owneruser"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)

	owner, err := userdomain.NewOwner(user)

	require.NoError(t, err)
	assert.NotNil(t, owner.User)
	assert.Equal(t, user, owner.User)
	assert.True(t, owner.IsOwner())
	assert.Equal(t, userID, owner.GetUserID())
	assert.Equal(t, organizationID, owner.GetOrganizationID())
}

func TestOwner_IsOwner_shouldReturnTrue_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "owner@example.com"
	username := "owneruser"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)

	result := owner.IsOwner()

	assert.True(t, result)
}

func TestOwner_GetUserID_shouldReturnUserID_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "owner@example.com"
	username := "owneruser"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)

	result := owner.GetUserID()

	require.NotNil(t, result)
	assert.Equal(t, userID, result)
	assert.Equal(t, 123, result.Value)
}

func TestOwner_GetOrganizationID_shouldReturnOrganizationID_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "owner@example.com"
	username := "owneruser"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)

	result := owner.GetOrganizationID()

	require.NotNil(t, result)
	assert.Equal(t, organizationID, result)
	assert.Equal(t, 456, result.Value)
}
