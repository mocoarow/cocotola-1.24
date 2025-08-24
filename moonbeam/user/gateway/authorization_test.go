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

func Test_AddPairOfUserAndGroup(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1; i++ {
		fn := func(t *testing.T, ctx context.Context, ts testService) {
			t.Helper()
			orgID, _, owner := setupOrganization(ctx, t, ts)
			defer teardownOrganization(t, ts, orgID)

			// given
			user1 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_1", "USERNAME_1", "PASSWORD_1")
			user2 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_2", "USERNAME_2", "PASSWORD_2")

			authorizationManager, err := gateway.NewAuthorizationManager(ctx, ts.dialect, ts.db, ts.rf)
			require.NoError(t, err)
			userGroupRepo := gateway.NewUserGroupRepository(ctx, ts.dialect, ts.db)
			ownerGroup, err := userGroupRepo.FindUserGroupByKey(ctx, owner, service.OwnerGroupKey)
			require.NoError(t, err)

			rbacRoleObject := domain.NewRBACUserRoleObject(orgID, ownerGroup.UserGroupID())

			// when
			ok, err := authorizationManager.CheckAuthorization(ctx, owner, service.RBACSetAction, rbacRoleObject)
			require.NoError(t, err)
			// then
			assert.True(t, ok, "owner should be able to make anyone belong to owner-group")
			if !ok {
				outputCasbinRule(t, ts.db)
			}

			// when
			ok, err = authorizationManager.CheckAuthorization(ctx, user2, service.RBACSetAction, rbacRoleObject)
			require.NoError(t, err)
			// then
			assert.False(t, ok, "standard-user should not be able to make other users belong to owner-group")

			// given
			// - add user1 to owner-group
			err = authorizationManager.AddUserToGroup(ctx, owner, user1.AppUserID(), ownerGroup.UserGroupID())
			require.NoError(t, err)
			// when
			ok, err = authorizationManager.CheckAuthorization(ctx, user1, service.RBACSetAction, rbacRoleObject)
			require.NoError(t, err)
			// then
			// - user1 can make sure user3 belong to group1 because user1 belongs to owner-group
			assert.True(t, ok)
		}
		testDB(t, fn)
	}
}
