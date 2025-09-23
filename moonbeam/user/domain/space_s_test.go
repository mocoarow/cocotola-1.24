//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewSpaceID_shouldReturnSpaceID_whenValidValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 123
	spaceID, err := userdomain.NewSpaceID(value)

	require.NoError(t, err)
	assert.Equal(t, value, spaceID.Value)
	assert.Equal(t, value, spaceID.Int())
	assert.True(t, spaceID.IsSpaceID())
}

func TestNewSpaceID_shouldReturnSpaceID_whenMinimumValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 1
	spaceID, err := userdomain.NewSpaceID(value)

	require.NoError(t, err)
	assert.Equal(t, value, spaceID.Value)
	assert.Equal(t, value, spaceID.Int())
	assert.True(t, spaceID.IsSpaceID())
}

func TestNewSpaceID_shouldReturnError_whenZeroValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := 0
	spaceID, err := userdomain.NewSpaceID(value)

	require.Error(t, err)
	assert.Nil(t, spaceID)
	assert.Contains(t, err.Error(), "validate space id")
}

func TestNewSpaceID_shouldReturnError_whenNegativeValueIsSpecified(t *testing.T) {
	t.Parallel()

	value := -1
	spaceID, err := userdomain.NewSpaceID(value)

	require.Error(t, err)
	assert.Nil(t, spaceID)
	assert.Contains(t, err.Error(), "validate space id")
}

func TestSpaceIDs_IDs_shouldReturnIntSlice_whenCalled(t *testing.T) {
	t.Parallel()

	spaceID1, _ := userdomain.NewSpaceID(1)
	spaceID2, _ := userdomain.NewSpaceID(2)
	spaceID3, _ := userdomain.NewSpaceID(3)
	spaceIDs := userdomain.SpaceIDs{spaceID1, spaceID2, spaceID3}

	ids := spaceIDs.IDs()

	require.Len(t, ids, 3)
	assert.Equal(t, []int{1, 2, 3}, ids)
}

func TestSpaceIDs_IDs_shouldReturnEmptySlice_whenEmptySpaceIDs(t *testing.T) {
	t.Parallel()

	spaceIDs := userdomain.SpaceIDs{}

	ids := spaceIDs.IDs()

	require.Empty(t, ids)
	assert.Equal(t, []int{}, ids)
}

func TestNewSpace_shouldReturnSpace_whenValidParametersAreSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "test-space"
	name := "Test Space"
	spaceType := "private"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	require.NoError(t, err)
	assert.Equal(t, baseModel, space.BaseModel)
	assert.Equal(t, spaceID, space.SpaceID)
	assert.Equal(t, organizationID, space.OrganizationID)
	assert.Equal(t, ownerID, space.OwnerID)
	assert.Equal(t, keyName, space.KeyName)
	assert.Equal(t, name, space.Name)
	assert.Equal(t, spaceType, space.SpaceType)
	assert.True(t, space.IsPrivate())
}

func TestNewSpace_shouldReturnSpace_whenPersonalSpaceTypeIsSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "personal-space"
	name := "Personal Space"
	spaceType := "personal"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	require.NoError(t, err)
	assert.Equal(t, spaceType, space.SpaceType)
	assert.False(t, space.IsPrivate())
}

func TestNewSpace_shouldReturnSpace_whenPublicSpaceTypeIsSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "public-space"
	name := "Public Space"
	spaceType := "public"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	require.NoError(t, err)
	assert.Equal(t, spaceType, space.SpaceType)
	assert.False(t, space.IsPrivate())
}

func TestNewSpace_shouldReturnError_whenSpaceIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "test-space"
	name := "Test Space"
	spaceType := "private"

	space, err := userdomain.NewSpace(baseModel, nil, organizationID, ownerID, keyName, name, spaceType)

	require.Error(t, err)
	assert.Nil(t, space)
	assert.Contains(t, err.Error(), "validate space model")
}

func TestNewSpace_shouldReturnError_whenOrganizationIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "test-space"
	name := "Test Space"
	spaceType := "private"

	space, err := userdomain.NewSpace(baseModel, spaceID, nil, ownerID, keyName, name, spaceType)

	require.Error(t, err)
	assert.Nil(t, space)
	assert.Contains(t, err.Error(), "validate space model")
}

func TestNewSpace_shouldReturnError_whenOwnerIDIsNil(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	keyName := "test-space"
	name := "Test Space"
	spaceType := "private"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, nil, keyName, name, spaceType)

	require.Error(t, err)
	assert.Nil(t, space)
	assert.Contains(t, err.Error(), "validate space model")
}

func TestNewSpace_shouldReturnError_whenKeyNameIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := ""
	name := "Test Space"
	spaceType := "private"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	require.Error(t, err)
	assert.Nil(t, space)
	assert.Contains(t, err.Error(), "validate space model")
}

func TestNewSpace_shouldReturnError_whenNameIsEmpty(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "test-space"
	name := ""
	spaceType := "private"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	require.Error(t, err)
	assert.Nil(t, space)
	assert.Contains(t, err.Error(), "validate space model")
}

func TestNewSpace_shouldReturnError_whenInvalidSpaceTypeIsSpecified(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "test-space"
	name := "Test Space"
	spaceType := "invalid"

	space, err := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	require.Error(t, err)
	assert.Nil(t, space)
	assert.Contains(t, err.Error(), "validate space model")
}

func TestSpace_IsPrivate_shouldReturnTrue_whenSpaceTypeIsPrivate(t *testing.T) {
	t.Parallel()

	baseModel := &domain.BaseModel{Version: 1}
	spaceID, _ := userdomain.NewSpaceID(123)
	organizationID, _ := userdomain.NewOrganizationID(456)
	ownerID, _ := userdomain.NewUserID(789)
	keyName := "private-space"
	name := "Private Space"
	spaceType := "private"

	space, _ := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, spaceType)

	result := space.IsPrivate()

	assert.True(t, result)
}

func TestSpace_IsPrivate_shouldReturnFalse_whenSpaceTypeIsNotPrivate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		spaceType string
	}{
		{"personal", "personal"},
		{"public", "public"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			baseModel := &domain.BaseModel{Version: 1}
			spaceID, _ := userdomain.NewSpaceID(123)
			organizationID, _ := userdomain.NewOrganizationID(456)
			ownerID, _ := userdomain.NewUserID(789)
			keyName := "test-space"
			name := "Test Space"

			space, _ := userdomain.NewSpace(baseModel, spaceID, organizationID, ownerID, keyName, name, tt.spaceType)

			result := space.IsPrivate()

			assert.False(t, result)
		})
	}
}
