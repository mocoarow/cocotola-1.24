//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewOrganizationID_shouldReturnOrganizationID_whenValidValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 123
	organizationID, err := userdomain.NewOrganizationID(value)

	require.NoError(t, err)
	assert.Equal(t, value, organizationID.Value)
	assert.Equal(t, value, organizationID.Int())
	assert.True(t, organizationID.IsOrganizationID())
}

func TestNewOrganizationID_shouldReturnOrganizationID_whenMinimumValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 1
	organizationID, err := userdomain.NewOrganizationID(value)

	require.NoError(t, err)
	assert.Equal(t, value, organizationID.Value)
	assert.Equal(t, value, organizationID.Int())
	assert.True(t, organizationID.IsOrganizationID())
}

func TestNewOrganizationModel_shouldReturnOrganization_whenValidParametersAreSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	organizationID, _ := userdomain.NewOrganizationID(123)
	name := "Test Organization"

	organization, err := userdomain.NewOrganizationModel(baseModel, organizationID, name)

	require.NoError(t, err)
	assert.Equal(t, baseModel, organization.BaseModel)
	assert.Equal(t, organizationID, organization.OrganizationID)
	assert.Equal(t, name, organization.Name)
}

func TestNewOrganizationModel_shouldReturnError_whenOrganizationIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	name := "Test Organization"

	organization, err := userdomain.NewOrganizationModel(baseModel, nil, name)

	require.Error(t, err)
	assert.Nil(t, organization)
	assert.Contains(t, err.Error(), "validate organization")
}

func TestNewOrganizationModel_shouldReturnError_whenNameIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	organizationID, _ := userdomain.NewOrganizationID(123)
	name := ""

	organization, err := userdomain.NewOrganizationModel(baseModel, organizationID, name)

	require.Error(t, err)
	assert.Nil(t, organization)
	assert.Contains(t, err.Error(), "validate organization")
}