//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewUserID_shouldReturnUserID_whenValidValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 123
	userID, err := userdomain.NewUserID(value)

	require.NoError(t, err)
	assert.Equal(t, value, userID.Value)
	assert.Equal(t, value, userID.Int())
	assert.True(t, userID.IsUserID())
}

func TestNewUserID_shouldReturnUserID_whenZeroValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 0
	userID, err := userdomain.NewUserID(value)

	require.NoError(t, err)
	assert.Equal(t, value, userID.Value)
	assert.Equal(t, value, userID.Int())
	assert.True(t, userID.IsUserID())
}

func TestNewUser_shouldReturnUser_whenValidParametersAreSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "test@example.com"
	username := "testuser"
	userGroups := []*userdomain.UserGroupModel{}

	user, err := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)

	require.NoError(t, err)
	assert.Equal(t, baseModel, user.BaseModel)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, organizationID, user.OrganizationID)
	assert.Equal(t, loginID, user.LoginID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, userGroups, user.UserGroups)
	assert.Equal(t, userID, user.GetUserID())
	assert.Equal(t, organizationID, user.GetOrganizationID())
}

func TestNewUser_shouldReturnError_whenUserIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "test@example.com"
	username := "testuser"
	userGroups := []*userdomain.UserGroupModel{}

	user, err := userdomain.NewUser(baseModel, nil, organizationID, loginID, username, userGroups)

	require.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "validate user model")
}

func TestNewUser_shouldReturnError_whenOrganizationIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	loginID := "test@example.com"
	username := "testuser"
	userGroups := []*userdomain.UserGroupModel{}

	user, err := userdomain.NewUser(baseModel, userID, nil, loginID, username, userGroups)

	require.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "validate user model")
}

func TestNewUser_shouldReturnError_whenLoginIDIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := ""
	username := "testuser"
	userGroups := []*userdomain.UserGroupModel{}

	user, err := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)

	require.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "validate user model")
}

func TestNewUser_shouldReturnError_whenUsernameIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userID, _ := userdomain.NewUserID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	loginID := "test@example.com"
	username := ""
	userGroups := []*userdomain.UserGroupModel{}

	user, err := userdomain.NewUser(baseModel, userID, organizationID, loginID, username, userGroups)

	require.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "validate user model")
}
