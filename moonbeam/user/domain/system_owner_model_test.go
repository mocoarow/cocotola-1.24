//go:build medium

package domain_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

func TestNewSystemOwner(t *testing.T) {
	t.Parallel()
	model, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)
	userID, err := domain.NewUserID(1)
	require.NoError(t, err)
	organizationID, err := domain.NewOrganizationID(1)
	require.NoError(t, err)
	user, err := domain.NewUserModel(model, userID, organizationID, "LOGIN_ID", "USERNAME", nil)
	require.NoError(t, err)
	ower, err := domain.NewOwnerModel(user)
	require.NoError(t, err)
	systemOwner, err := domain.NewSystemOwnerModel(ower)
	require.NoError(t, err)
	log.Println(systemOwner)
}
