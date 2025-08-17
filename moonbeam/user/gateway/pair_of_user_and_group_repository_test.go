//go:build medium

package gateway_test

import (
	"context"
	"testing"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_pairOfUserAndGroupRepository_FindUserGroupsByUserID(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		orgID, _, owner := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)

		// given
		user1 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_1", "USERNAME_1", "PASSWORD_1")
		user2 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_2", "USERNAME_2", "PASSWORD_2")
		group1 := testAddUserGroup(t, ctx, ts, owner, "GROUP_KEY_1", "GROUP_NAME_1", "GROUP_DESC_1")
		group2 := testAddUserGroup(t, ctx, ts, owner, "GROUP_KEY_2", "GROUP_NAME_2", "GROUP_DESC_2")
		group3 := testAddUserGroup(t, ctx, ts, owner, "GROUP_KEY_3", "GROUP_NAME_3", "GROUP_DESC_3")

		pairOfUserAndGroupRepo := gateway.NewPairOfUserAndGroupRepository(ctx, ts.dialect, ts.db, ts.rf)

		// - user1 belongs to group1, group2, group3
		for _, group := range []*service.UserGroup{group1, group2, group3} {
			err := pairOfUserAndGroupRepo.AddPairOfUserAndGroup(ctx, owner, user1.AppUserID(), group.UserGroupID())
			require.NoError(t, err)
		}
		// - user2 belongs to group1
		for _, group := range []*service.UserGroup{group1} {
			err := pairOfUserAndGroupRepo.AddPairOfUserAndGroup(ctx, owner, user2.AppUserID(), group.UserGroupID())
			require.NoError(t, err)
		}

		// when
		groupModels1, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, owner, user1.AppUserID())
		require.NoError(t, err)
		groupModels2, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, owner, user2.AppUserID())
		require.NoError(t, err)
		groups1 := testNewUserGroups(groupModels1)
		groups2 := testNewUserGroups(groupModels2)

		// then
		// - user1 belongs to group1, group2, group3
		assert.Len(t, groups1, 3)
		assert.Equal(t, "GROUP_KEY_1", groups1[0].Key())
		assert.Equal(t, "GROUP_KEY_2", groups1[1].Key())
		assert.Equal(t, "GROUP_KEY_3", groups1[2].Key())
		// - user2 belongs to group1
		assert.Len(t, groups2, 1)
		assert.Equal(t, "GROUP_KEY_1", groups2[0].Key())
	}
	testDB(t, fn)
}

func Test_pairOfUserAndGroupRepository_RemovePairOfUserAndGroup(t *testing.T) {
	t.Parallel()
	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		orgID, _, owner := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)

		// given
		user1 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_1", "USERNAME_1", "PASSWORD_1")
		// user2 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_2", "USERNAME_2", "PASSWORD_2")
		// user3 := testAddAppUser(t, ctx, ts, owner, "LOGIN_ID_3", "USERNAME_3", "PASSWORD_2")
		// group1 := testAddUserGroup(t, ctx, ts, owner, "GROUP_KEY_1", "GROUP_NAME_1", "GROUP_DESC_1")

		pairOfUserAndGroupRepo := gateway.NewPairOfUserAndGroupRepository(ctx, ts.dialect, ts.db, ts.rf)
		userGroupRepo := gateway.NewUserGroupRepository(ctx, ts.dialect, ts.db)
		ownerGroup, err := userGroupRepo.FindUserGroupByKey(ctx, owner, service.OwnerGroupKey)
		require.NoError(t, err)

		err = pairOfUserAndGroupRepo.AddPairOfUserAndGroup(ctx, owner, user1.AppUserID(), ownerGroup.UserGroupID())
		require.NoError(t, err)
		{
			userGroups1, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, user1, user1.AppUserID())
			require.NoError(t, err)
			assert.Len(t, userGroups1, 1)
		}

		// when
		err = pairOfUserAndGroupRepo.RemovePairOfUserAndGroup(ctx, owner, user1.AppUserID(), ownerGroup.UserGroupID())
		require.NoError(t, err)

		// then
		// - owner can not make user1 belong to owner-group

		// // when
		// err = pairOfUserAndGroupRepo.AddPairOfUserAndGroup(ctx, user2, user3.GetAppUserID(), ownerGroup.GetUerGroupID())
		// // then
		// // - user2 can NOT make user3 belong to owner-group
		// assert.ErrorIs(t, err, libdomain.ErrPermissionDenied)

		// // when
		// err = pairOfUserAndGroupRepo.AddPairOfUserAndGroup(ctx, user1, user3.GetAppUserID(), group1.GetUerGroupID())
		// // - user1 can make sure user3 belong to group1 because user1 belongs to owner-group
		// assert.NoError(t, err)
		// {
		// 	userGroups3, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, user1, user3.GetAppUserID())
		// 	require.NoError(t, err)
		// 	assert.Len(t, userGroups3, 1)
		// }
	}
	testDB(t, fn)
}
