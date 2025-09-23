//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewSystemOwner_shouldReturnSystemOwner_whenValidOwnerIsSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "systemowner@example.com"
	username := "systemowner"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)

	systemOwner, err := userdomain.NewSystemOwner(owner)

	require.NoError(t, err)
	assert.NotNil(t, systemOwner.Owner)
	assert.Equal(t, owner, systemOwner.Owner)
	assert.True(t, systemOwner.IsOwner())
	assert.True(t, systemOwner.IsSystemOwner())
	assert.Equal(t, userID, systemOwner.GetUserID())
	assert.Equal(t, organizationID, systemOwner.GetOrganizationID())
}

func TestNewSystemOwner_shouldReturnError_whenOwnerIsNil(t *testing.T) {
	t.Parallel()

	systemOwner, err := userdomain.NewSystemOwner(nil)

	require.Error(t, err)
	assert.Nil(t, systemOwner)
	assert.Contains(t, err.Error(), "validate system owner")
}

func TestSystemOwner_IsOwner_shouldReturnTrue_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "systemowner@example.com"
	username := "systemowner"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)
	systemOwner, _ := userdomain.NewSystemOwner(owner)

	result := systemOwner.IsOwner()

	assert.True(t, result)
}

func TestSystemOwner_IsSystemOwner_shouldReturnTrue_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "systemowner@example.com"
	username := "systemowner"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)
	systemOwner, _ := userdomain.NewSystemOwner(owner)

	result := systemOwner.IsSystemOwner()

	assert.True(t, result)
}

func TestSystemOwner_GetUserID_shouldReturnUserID_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "systemowner@example.com"
	username := "systemowner"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)
	systemOwner, _ := userdomain.NewSystemOwner(owner)

	result := systemOwner.GetUserID()

	require.NotNil(t, result)
	assert.Equal(t, userID, result)
	assert.Equal(t, 123, result.Value)
}

func TestSystemOwner_GetOrganizationID_shouldReturnOrganizationID_whenCalled(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "systemowner@example.com"
	username := "systemowner"
	userGroups := []*userdomain.UserGroup{}

	user, _ := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)
	owner, _ := userdomain.NewOwner(user)
	systemOwner, _ := userdomain.NewSystemOwner(owner)

	result := systemOwner.GetOrganizationID()

	require.NotNil(t, result)
	assert.Equal(t, organizationID, result)
	assert.Equal(t, 456, result.Value)
}
