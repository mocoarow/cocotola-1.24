//go:build medium

package gateway_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

func Test_userRepository_FindSystemOwnerByOrganizationID_shouldReturnSystemOwner_whenExistingOrganizationIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		sysAd := domain.NewSystemAdmin()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// when
		testSysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, sysAd, orgID)

		// then
		require.NoError(t, err)
		assert.Equal(t, service.SystemOwnerLoginID, testSysOwner.LoginID)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindSystemOwnerByOrganizationID_shouldReturnError_whenInvalidOrganizationIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		sysAd := domain.NewSystemAdmin()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// when
		_, err := userRepo.FindSystemOwnerByOrganizationID(ctx, sysAd, invalidOrgID)

		// then
		assert.ErrorIs(t, err, service.ErrSystemOwnerNotFound)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindSystemOwnerByOrganizationName_shouldReturnSystemOwner_whenExistingNameIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		org := getOrganization(t, ctx, ts, orgID)
		sysAd := domain.NewSystemAdmin()

		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// when
		testSysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, sysAd, org.Name)

		// then
		require.NoError(t, err)
		assert.Equal(t, service.SystemOwnerLoginID, testSysOwner.LoginID)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindSystemOwnerByOrganizationName_shouldReturnError_whenInvalidNameIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		sysAd := domain.NewSystemAdmin()

		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// when
		_, err := userRepo.FindSystemOwnerByOrganizationName(ctx, sysAd, "NOT_FOUND")

		// then
		assert.ErrorIs(t, err, service.ErrSystemOwnerNotFound)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindUserByID_shouldReturnUser_whenExistingIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// given
		newUser := testAddUser(t, ctx, ts, owner, "LOGIN_ID", "USERNAME", "PASSWORD")

		// when
		user, err := userRepo.FindUserByID(ctx, owner, newUser.GetUserID())

		// then
		require.NoError(t, err)
		assert.Equal(t, "LOGIN_ID", user.LoginID, "loginID should be 'LOGIN_ID'")
		assert.Equal(t, "USERNAME", user.Username, "username should be 'USERNAME'")
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindUserByID_shouldReturnError_whenInvaildIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		_, err := userRepo.FindUserByID(ctx, owner, invalidUserID)
		assert.ErrorIs(t, err, service.ErrUserNotFound)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindUserByLoginID_shouldReturnUser_whenExistingLoginIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// given
		_ = testAddUser(t, ctx, ts, owner, "LOGIN_ID", "USERNAME", "PASSWORD")

		// when
		user, err := userRepo.FindUserByLoginID(ctx, owner, "LOGIN_ID")

		// then
		require.NoError(t, err)
		assert.Equal(t, "LOGIN_ID", user.LoginID, "loginID should be 'LOGIN_ID'")
		assert.Equal(t, "USERNAME", user.Username, "username should be 'USERNAME'")
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindUserByLoginID_shouldReturnError_whenInvalidLoginIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// when
		_, err := userRepo.FindUserByLoginID(ctx, owner, "NOT_FOUND")

		// then
		assert.ErrorIs(t, err, service.ErrUserNotFound)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindOwnerByLoginID_shouldReturnOwner_whenExistingOwnerLoginIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		require.Equal(t, "OWNER_ID", owner.LoginID)
		require.Equal(t, "OWNER_NAME", owner.Username)

		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// when
		user, err := userRepo.FindOwnerByLoginID(ctx, sysOwner, owner.LoginID)

		// then
		require.NoError(t, err)
		assert.Equal(t, "OWNER_ID", user.LoginID)
		assert.Equal(t, "OWNER_NAME", user.Username)
	}
	testOrganization(t, fn)
}

func Test_userRepository_FindOwnerByLoginID_shouldReturnError_whenNotOwnerLoginIDIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		require.Equal(t, "OWNER_ID", owner.LoginID)
		require.Equal(t, "OWNER_NAME", owner.Username)

		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)

		// given
		_ = testAddUser(t, ctx, ts, owner, "LOGIN_ID", "USERNAME", "PASSWORD")

		// when
		_, err := userRepo.FindOwnerByLoginID(ctx, sysOwner, "LOGIN_ID")

		// then
		assert.ErrorIs(t, err, service.ErrUserNotFound)
	}
	testOrganization(t, fn)
}

func Test_userRepository_VerifyPassword_shouldReturnTrue_whenCorrectPasswordIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		sysAd := domain.NewSystemAdmin()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)
		testSysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, sysAd, orgID)
		require.NoError(t, err)

		// given
		_ = testAddUser(t, ctx, ts, owner, "LOGIN_ID", "USERNAME", "PASSWORD")

		// when
		verified, err := userRepo.VerifyPassword(ctx, testSysOwner, "LOGIN_ID", "PASSWORD")

		// then
		assert.True(t, verified)
		assert.NoError(t, err)
	}
	testOrganization(t, fn)
}

func Test_userRepository_VerifyPassword_shouldReturnFalse_whenWrongPasswordIsSpecified(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.Owner) {
		t.Helper()
		sysAd := domain.NewSystemAdmin()
		userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)
		testSysOwner, err := userRepo.FindSystemOwnerByOrganizationID(ctx, sysAd, orgID)
		require.NoError(t, err)

		// given
		_ = testAddUser(t, ctx, ts, owner, "LOGIN_ID", "USERNAME", "PASSWORD")

		// when
		verified, err := userRepo.VerifyPassword(ctx, testSysOwner, "LOGIN_ID", "WRONG_PASSWORD")

		// then
		assert.False(t, verified)
		assert.NoError(t, err)
	}
	testOrganization(t, fn)
}
