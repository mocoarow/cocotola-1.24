//go:build small

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	userdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewSystemAdmin_shouldReturnSystemAdmin_whenCalled(t *testing.T) {
	t.Parallel()

	systemAdmin := userdomain.NewSystemAdmin()

	require.NotNil(t, systemAdmin)
	assert.NotNil(t, systemAdmin.UserID)
	assert.Equal(t, 1, systemAdmin.UserID.Value)
	assert.True(t, systemAdmin.IsSystemAdmin())
	assert.Equal(t, systemAdmin.UserID, systemAdmin.GetUserID())
}

func TestSystemAdmin_IsSystemAdmin_shouldReturnTrue_whenCalled(t *testing.T) {
	t.Parallel()

	systemAdmin := userdomain.NewSystemAdmin()

	result := systemAdmin.IsSystemAdmin()

	assert.True(t, result)
}

func TestSystemAdmin_GetUserID_shouldReturnUserID_whenCalled(t *testing.T) {
	t.Parallel()

	systemAdmin := userdomain.NewSystemAdmin()

	userID := systemAdmin.GetUserID()

	require.NotNil(t, userID)
	assert.Equal(t, 1, userID.Value)
	assert.Equal(t, systemAdmin.UserID, userID)
}