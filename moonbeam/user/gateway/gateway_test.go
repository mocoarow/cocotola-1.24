//go:build medium

package gateway_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	testlibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/testlib/gateway"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

const orgNameLength = 8

type testService struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

func outputOrganization(t *testing.T, db *gorm.DB) {
	t.Helper()
	var results []gateway.OrganizationEntity
	if result := db.Find(&results); result.Error != nil {
		assert.Fail(t, result.Error.Error())
	}
	var s string
	s += "\n   id,version,           created_at,          updated_at,created_by,updated_by,      name,"
	for i := range results {
		result := &results[i]
		s += fmt.Sprintf("\n%5d,%8d,%20s,%20s,%10d,%10d,%10s", result.ID, result.Version, result.CreatedAt.Format(time.RFC3339), result.UpdatedAt.Format(time.RFC3339), result.CreatedBy, result.UpdatedBy, result.Name)
	}
	t.Log(s)
}

func outputCasbinRule(t *testing.T, db *gorm.DB) {
	t.Helper()
	type Result struct {
		ID    int
		Ptype string
		V0    string
		V1    string
		V2    string
		V3    string
		V4    string
		V5    string
	}
	var results []Result
	if result := db.Raw("SELECT * FROM casbin_rule").Scan(&results); result.Error != nil {
		assert.Fail(t, result.Error.Error())
	}
	var s string
	s += "\n   id,ptype,                  v0,                  v1,         v2,         v3,         v4,         v5"
	for i := range results {
		result := &results[i]
		s += fmt.Sprintf("\n%5d,%5s,%20s,%20s, %10s, %10s, %10s, %10s", result.ID, result.Ptype, result.V0, result.V1, result.V2, result.V3, result.V4, result.V5)
	}
	t.Log(s)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			panic(err)
		}
		b[i] = letterRunes[val.Int64()]
	}
	return string(b)
}

func testDB(t *testing.T, fn func(t *testing.T, ctx context.Context, ts testService)) {
	t.Helper()
	ctx := context.Background()
	var resourceEventHandlers map[domain.ResourceKey]libservice.ResourceEventHandler
	for dialect, db := range testlibgateway.ListDB() {
		dialect := dialect
		db := db
		t.Run(dialect.Name(), func(t *testing.T) {
			// t.Parallel()
			sqlDB, err := db.DB()
			require.NoError(t, err)
			defer sqlDB.Close()

			rf, err := gateway.NewRepositoryFactory(ctx, dialect, dialect.Name(), db, loc, resourceEventHandlers)
			require.NoError(t, err)
			testService := testService{dialect: dialect, db: db, rf: rf}

			fn(t, ctx, testService)
		})
	}
}

func testOrganization(t *testing.T, fn func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *domain.SystemOwner, owner *domain.OwnerModel)) {
	t.Helper()
	testDB(t, func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		orgID, sysOwner, owner := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)

		fn(t, ctx, ts, orgID, sysOwner, owner)
	})
}

func setupOrganization(ctx context.Context, t *testing.T, ts testService) (*domain.OrganizationID, *domain.SystemOwner, *domain.OwnerModel) {
	t.Helper()
	orgName := RandString(orgNameLength)
	sysAd := domain.NewSystemAdmin()

	firstOwnerAddParam, err := service.NewUserAddParameter("OWNER_ID", "OWNER_NAME", "OWNER_PASSWORD", "", "", "", "")
	require.NoError(t, err)
	// orgAddParam, err := service.NewOrganizationAddParameter(orgName, firstOwnerAddParam)
	// require.NoError(t, err)

	orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)
	userRepo := gateway.NewUserRepository(ctx, ts.dialect, ts.db, ts.rf)
	userGorupRepo := gateway.NewUserGroupRepository(ctx, ts.dialect, ts.db)
	authorizationManager, err := gateway.NewAuthorizationManager(ctx, ts.dialect, ts.db, ts.rf)
	require.NoError(t, err)

	// 1. add organization
	orgID, err := orgRepo.AddOrganization(ctx, sysAd, orgName)
	if err != nil {
		outputOrganization(t, ts.db)
	}
	require.NoError(t, err)
	assert.Positive(t, orgID.Int())

	// 2. add "system-owner" user
	sysOwnerID, err := userRepo.AddSystemOwner(ctx, sysAd, orgID)
	require.NoError(t, err)
	require.Positive(t, sysOwnerID.Int())

	// TODO
	sysOwner, err := userRepo.FindSystemOwnerByOrganizationName(ctx, sysAd, orgName, service.IncludeGroups)
	require.NoError(t, err)

	// 3. add policy to "system-owner" user
	t.Log(`add policy to "system-owner" user`)
	rbacSysOwner := domain.NewRBACUserFromUser(sysOwnerID)
	rbacAllUserRolesObject := domain.NewRBACAllUserRolesObjectFromOrganization(orgID)
	// - "system-owner" "can" "set" "all-user-roles"
	err = authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAd, orgID, rbacSysOwner, service.RBACSetAction, rbacAllUserRolesObject, service.RBACAllowEffect)
	require.NoError(t, err)
	outputCasbinRule(t, ts.db)

	// - "system-owner" "can" "unset" "all-user-roles"
	err = authorizationManager.AddPolicyToUserBySystemAdmin(ctx, sysAd, orgID, rbacSysOwner, service.RBACUnsetAction, rbacAllUserRolesObject, service.RBACAllowEffect)
	require.NoError(t, err)

	// 4. add owner-group
	ownerGroupID, err := userGorupRepo.AddOwnerGroup(ctx, sysOwner, orgID)
	require.NoError(t, err)

	// 5. add policty to "owner" group
	rbacOwnerGroup := domain.NewRBACRoleFromGroup(orgID, ownerGroupID)
	// - "owner" group "can" "set" "all-user-roles"
	err = authorizationManager.AddPolicyToGroupBySystemAdmin(ctx, sysAd, orgID, rbacOwnerGroup, service.RBACSetAction, rbacAllUserRolesObject, service.RBACAllowEffect)
	require.NoError(t, err)
	// - "owner" group "can" "unset" "all-user-roles"
	err = authorizationManager.AddPolicyToGroupBySystemAdmin(ctx, sysAd, orgID, rbacOwnerGroup, service.RBACUnsetAction, rbacAllUserRolesObject, service.RBACAllowEffect)
	require.NoError(t, err)

	// 6. add first owner
	ownerID, err := userRepo.AddUser(ctx, sysOwner, firstOwnerAddParam)
	require.NoError(t, err)
	require.Positive(t, ownerID.Int())

	// - owner belongs to owner-group
	err = authorizationManager.AddUserToGroup(ctx, sysOwner, ownerID, ownerGroupID)
	require.NoError(t, err)

	owner, err := userRepo.FindOwnerByLoginID(ctx, sysOwner, firstOwnerAddParam.LoginID)
	require.NoError(t, err)

	// logger := slog.Default()
	// logger.Warn(fmt.Sprintf("orgID: %d", orgID.Int()))

	return orgID, sysOwner, owner
}

func teardownOrganization(t *testing.T, ts testService, orgID *domain.OrganizationID) {
	t.Helper()
	// delete all organizations
	// ts.db.Exec("delete from space where organization_id = ?", orgID.Int())
	ts.db.Exec("delete from mb_user where organization_id = ?", orgID.Int())
	ts.db.Exec("delete from mb_organization where id = ?", orgID.Int())
	// db.Where("true").Delete(&spaceEntity{})
	// db.Where("true").Delete(&userEntity{})
	// db.Where("true").Delete(&organizationEntity{})
}

func testAddUser(t *testing.T, ctx context.Context, ts testService, owner domain.OwnerInterface, loginID, username, password string) *domain.User {
	t.Helper()
	userRepo := ts.rf.NewUserRepository(ctx)
	userID1, err := userRepo.AddUser(ctx, owner, testNewUserAddParameter(t, loginID, username, password))
	require.NoError(t, err)
	user1, err := userRepo.FindUserByID(ctx, owner, userID1)
	require.NoError(t, err)
	require.Equal(t, loginID, user1.LoginID)

	return user1
}

func testAddUserGroup(t *testing.T, ctx context.Context, ts testService, owner domain.OwnerInterface, key, name, description string) *domain.UserGroup {
	t.Helper()
	userGorupRepo := ts.rf.NewUserGroupRepository(ctx)
	groupID1, err := userGorupRepo.AddUserGroup(ctx, owner, testNewUserGroupAddParameter(t, key, name, description))
	require.NoError(t, err)
	group1, err := userGorupRepo.FindUserGroupByID(ctx, owner, groupID1)
	require.NoError(t, err)
	require.Equal(t, key, group1.Key)
	require.Equal(t, name, group1.Name)
	require.Equal(t, description, group1.Description)

	return group1
}

// type testSystemAdmin struct {
// 	*domain.SystemAdmin
// }

// func (m *testSystemAdmin) GetUserID() *domain.UserID {
// 	return m.UserID
// }
// func (m *testSystemAdmin) IsSystemAdmin() bool {
// 	return true
// }
// func testNewSystemAdmin(systemAdminModel *domain.SystemAdmin) *testSystemAdmin {
// 	return &testSystemAdmin{
// 		systemAdminModel,
// 	}
// }

// type testUser struct {
// 	*domain.User
// }

// func (m *testUser) GetUserID() *domain.UserID {
// 	return m.UserID
// }
// func (m *testUser) GetOrganizationID() *domain.OrganizationID {
// 	return m.OrganizationID
// }

// //	func (m *testUser) LoginID() string {
// //		return m.User.LoginID
// //	}
// //
// //	func (m *testUser) Username() string {
// //		return m.User.Username
// //	}
// func testNewUser(userModel *domain.User) *testUser {
// 	return &testUser{
// 		userModel,
// 	}
// }

// type testUserGroup struct {
// 	*domain.UserGroup
// }

// func (m *testUserGroup) Key() string {
// 	return m.UserGroup.Key
// }
// func (m *testUserGroup) Name() string {
// 	return m.UserGroup.Key
// }
// func (m *testUserGroup) Description() string {
// 	return m.UserGroup.Description
// }
// func testNewUserGroup(userGroupModel *domain.UserGroup) *testUserGroup {
// 	return &testUserGroup{
// 		userGroupModel,
// 	}
// }
// func testNewUserGroups(userGroupModels []*domain.UserGroup) []*testUserGroup {
// 	groups := make([]*testUserGroup, len(userGroupModels))
// 	for i, groupModel := range userGroupModels {
// 		groups[i] = testNewUserGroup(groupModel)
// 	}

// 	return groups
// }

func testNewUserAddParameter(t *testing.T, loginID, username, password string) *service.AddUserParameter {
	t.Helper()
	p, err := service.NewUserAddParameter(loginID, username, password, "", "", "", "")
	require.NoError(t, err)

	return p
}

func testNewUserGroupAddParameter(t *testing.T, key, name, description string) *service.AddUserGroupParameter {
	t.Helper()
	p, err := service.NewUserGroupAddParameter(key, name, description)
	require.NoError(t, err)

	return p
}

func getOrganization(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID) *domain.Organization {
	t.Helper()
	orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)

	baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)
	userID, _ := domain.NewUserID(1)
	userModel, err := domain.NewUser(baseModel, userID, orgID, "login_id", "username", nil)
	require.NoError(t, err)

	org, err := orgRepo.GetOrganization(ctx, userModel)
	require.NoError(t, err)
	require.Len(t, org.Name, orgNameLength)

	return org
}
