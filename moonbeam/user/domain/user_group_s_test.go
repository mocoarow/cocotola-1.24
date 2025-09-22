//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewUserGroupID_shouldReturnUserGroupID_whenValidValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 123
	userGroupID, err := userdomain.NewUserGroupID(value)

	require.NoError(t, err)
	assert.Equal(t, value, userGroupID.Value)
	assert.Equal(t, value, userGroupID.Int())
	assert.True(t, userGroupID.IsUserGroupID())
}

func TestNewUserGroupID_shouldReturnUserGroupID_whenZeroValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 0
	userGroupID, err := userdomain.NewUserGroupID(value)

	require.NoError(t, err)
	assert.Equal(t, value, userGroupID.Value)
	assert.Equal(t, value, userGroupID.Int())
	assert.True(t, userGroupID.IsUserGroupID())
}

func TestNewUserGroup_shouldReturnUserGroup_whenValidParametersAreSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userGroupID, _ := userdomain.NewUserGroupID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	key := "admin"
	name := "Administrator"
	description := "System administrator group"

	userGroup, err := userdomain.NewUserGroup(baseModel, userGroupID, organizationID, key, name, description)

	require.NoError(t, err)
	assert.Equal(t, baseModel, userGroup.BaseModel)
	assert.Equal(t, userGroupID, userGroup.UserGroupID)
	assert.Equal(t, organizationID, userGroup.OrganizationID)
	assert.Equal(t, key, userGroup.Key)
	assert.Equal(t, name, userGroup.Name)
	assert.Equal(t, description, userGroup.Description)
}

func TestNewUserGroup_shouldReturnUserGroup_whenEmptyDescriptionIsSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userGroupID, _ := userdomain.NewUserGroupID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	key := "user"
	name := "Regular User"
	description := ""

	userGroup, err := userdomain.NewUserGroup(baseModel, userGroupID, organizationID, key, name, description)

	require.NoError(t, err)
	assert.Equal(t, baseModel, userGroup.BaseModel)
	assert.Equal(t, userGroupID, userGroup.UserGroupID)
	assert.Equal(t, organizationID, userGroup.OrganizationID)
	assert.Equal(t, key, userGroup.Key)
	assert.Equal(t, name, userGroup.Name)
	assert.Equal(t, description, userGroup.Description)
}

func TestNewUserGroup_shouldReturnError_whenUserGroupIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	organizationID, _ := userdomain.NewOrganizationID(456)
	key := "admin"
	name := "Administrator"
	description := "System administrator group"

	userGroup, err := userdomain.NewUserGroup(baseModel, nil, organizationID, key, name, description)

	require.Error(t, err)
	assert.Nil(t, userGroup)
	assert.Contains(t, err.Error(), "validate user group")
}

func TestNewUserGroup_shouldReturnError_whenOrganizationIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userGroupID, _ := userdomain.NewUserGroupID(123)
	key := "admin"
	name := "Administrator"
	description := "System administrator group"

	userGroup, err := userdomain.NewUserGroup(baseModel, userGroupID, nil, key, name, description)

	require.Error(t, err)
	assert.Nil(t, userGroup)
	assert.Contains(t, err.Error(), "validate user group")
}

func TestNewUserGroup_shouldReturnError_whenKeyIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userGroupID, _ := userdomain.NewUserGroupID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	key := ""
	name := "Administrator"
	description := "System administrator group"

	userGroup, err := userdomain.NewUserGroup(baseModel, userGroupID, organizationID, key, name, description)

	require.Error(t, err)
	assert.Nil(t, userGroup)
	assert.Contains(t, err.Error(), "validate user group")
}

func TestNewUserGroup_shouldReturnError_whenNameIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	userGroupID, _ := userdomain.NewUserGroupID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	key := "admin"
	name := ""
	description := "System administrator group"

	userGroup, err := userdomain.NewUserGroup(baseModel, userGroupID, organizationID, key, name, description)

	require.Error(t, err)
	assert.Nil(t, userGroup)
	assert.Contains(t, err.Error(), "validate user group")
}