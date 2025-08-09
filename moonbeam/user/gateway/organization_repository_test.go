//go:build medium

package gateway_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func Test_organizationRepository_GetOrganization(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		// require.Equal(t, ts.dialect.BoolDefaultValue(), "false")
		orgID, _, _ := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)

		orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)

		// get organization registered
		baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
		require.NoError(t, err)
		appUserID, _ := domain.NewAppUserID(1)
		userModel, err := domain.NewAppUserModel(baseModel, appUserID, orgID, "login_id", "username", nil)
		require.NoError(t, err)
		user := testNewAppUser(userModel)
		{
			org, err := orgRepo.GetOrganization(ctx, user)
			assert.NoError(t, err)
			assert.Equal(t, orgNameLength, len(org.Name()))
		}

		// get organization unregistered
		otherAppUserModel, err := domain.NewAppUserModel(baseModel, appUserID, invalidOrgID, "login_id", "username", nil)
		assert.NoError(t, err)
		otherAppUser := testNewAppUser(otherAppUserModel)
		{
			_, err := orgRepo.GetOrganization(ctx, otherAppUser)
			assert.ErrorIs(t, err, service.ErrOrganizationNotFound)
		}
	}
	testDB(t, fn)
}

func Test_organizationRepository_FindOrganizationByName(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		orgID, _, _ := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)
		sysAdModel := domain.NewSystemAdminModel()
		sysAd := testNewSystemAdmin(sysAdModel)

		orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)

		var orgName string

		// get organization registered
		baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
		assert.NoError(t, err)
		appUserID, err := domain.NewAppUserID(1)
		require.NoError(t, err)

		userModel, err := domain.NewAppUserModel(baseModel, appUserID, orgID, "login_id", "username", nil)
		assert.NoError(t, err)
		user := testNewAppUser(userModel)
		{
			org, err := orgRepo.GetOrganization(ctx, user)
			assert.NoError(t, err)
			assert.Equal(t, orgNameLength, len(org.Name()))
			orgName = org.Name()
		}

		// find organization registered by name
		{
			org, err := orgRepo.FindOrganizationByName(ctx, sysAd, orgName)
			assert.NoError(t, err)
			assert.Equal(t, orgName, org.Name())
		}

		// find organization unregistered by name
		{
			_, err := orgRepo.FindOrganizationByName(ctx, sysAd, "NOT_FOUND")
			assert.Equal(t, service.ErrOrganizationNotFound, err)
		}
	}
	testDB(t, fn)
}

func Test_organizationRepository_FindOrganizationByID(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		orgID, _, _ := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)
		sysAdModel := domain.NewSystemAdminModel()
		sysAd := testNewSystemAdmin(sysAdModel)

		orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)

		// get organization registered
		baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
		assert.NoError(t, err)
		appUserID, err := domain.NewAppUserID(1)
		require.NoError(t, err)

		userModel, err := domain.NewAppUserModel(baseModel, appUserID, orgID, "login_id", "username", nil)
		assert.NoError(t, err)
		user := testNewAppUser(userModel)
		{
			org, err := orgRepo.GetOrganization(ctx, user)
			assert.NoError(t, err)
			assert.Equal(t, orgNameLength, len(org.Name()))
		}

		// find organization registered by ID
		{
			org, err := orgRepo.FindOrganizationByID(ctx, sysAd, orgID)
			assert.NoError(t, err)
			assert.Equal(t, orgID.Int(), org.OrganizationID().Int())
		}

		// find organization unregistered by ID
		{
			_, err := orgRepo.FindOrganizationByID(ctx, sysAd, invalidOrgID)
			assert.Equal(t, service.ErrOrganizationNotFound, err)
		}
	}
	testDB(t, fn)
}
