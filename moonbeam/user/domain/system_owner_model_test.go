//go:build medium

package domain_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewSystemOwner(t *testing.T) {
	t.Parallel()
	model, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)
	appUserID, err := domain.NewAppUserID(1)
	require.NoError(t, err)
	organizationID, err := domain.NewOrganizationID(1)
	require.NoError(t, err)
	appUser, err := domain.NewAppUserModel(model, appUserID, organizationID, "LOGIN_ID", "USERNAME", nil)
	assert.NoError(t, err)
	ower, err := domain.NewOwnerModel(appUser)
	assert.NoError(t, err)
	systemOwner, err := domain.NewSystemOwnerModel(ower)
	assert.NoError(t, err)
	log.Println(systemOwner)
}
