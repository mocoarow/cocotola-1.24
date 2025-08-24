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
	resourceEventHandlerFuncs := libservice.ResourceEventHandlerFuncs{}
	for dialect, db := range testlibgateway.ListDB() {
		dialect := dialect
		db := db
		t.Run(dialect.Name(), func(t *testing.T) {
			// t.Parallel()
			sqlDB, err := db.DB()
			require.NoError(t, err)
			defer sqlDB.Close()

			rf, err := gateway.NewRepositoryFactory(ctx, dialect, dialect.Name(), db, loc, resourceEventHandlerFuncs)
			require.NoError(t, err)
			testService := testService{dialect: dialect, db: db, rf: rf}

			fn(t, ctx, testService)
		})
	}
}

func testOrganization(t *testing.T, fn func(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID, sysOwner *service.SystemOwner, owner *service.Owner)) {
	t.Helper()
	testDB(t, func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		orgID, sysOwner, owner := setupOrganization(ctx, t, ts)
		defer teardownOrganization(t, ts, orgID)

		fn(t, ctx, ts, orgID, sysOwner, owner)
	})
}

func setupOrganization(ctx context.Context, t *testing.T, ts testService) (*domain.OrganizationID, *service.SystemOwner, *service.Owner) {
	t.Helper()
	orgName := RandString(orgNameLength)
	sysAd, err := service.NewSystemAdmin(ctx, ts.rf)
	require.NoError(t, err)

	firstOwnerAddParam, err := service.NewAppUserAddParameter("OWNER_ID", "OWNER_NAME", "OWNER_PASSWORD", "", "", "", "")
	require.NoError(t, err)
	orgAddParam, err := service.NewOrganizationAddParameter(orgName, firstOwnerAddParam)
	require.NoError(t, err)

	orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)
	appUserRepo := gateway.NewAppUserRepository(ctx, ts.dialect, ts.db, ts.rf)
	userGorupRepo := gateway.NewUserGroupRepository(ctx, ts.dialect, ts.db)
	authorizationManager, err := gateway.NewAuthorizationManager(ctx, ts.dialect, ts.db, ts.rf)
	require.NoError(t, err)

	// 1. add organization
	orgID, err := orgRepo.AddOrganization(ctx, sysAd, orgAddParam)
	if err != nil {
		outputOrganization(t, ts.db)
	}
	require.NoError(t, err)
	assert.Positive(t, orgID.Int())

	// 2. add "system-owner" user
	sysOwnerID, err := appUserRepo.AddSystemOwner(ctx, sysAd, orgID)
	require.NoError(t, err)
	require.Positive(t, sysOwnerID.Int())

	// TODO
	sysOwner, err := appUserRepo.FindSystemOwnerByOrganizationName(ctx, sysAd, orgName, service.IncludeGroups)
	require.NoError(t, err)

	// 3. add policy to "system-owner" user
	t.Log(`add policy to "system-owner" user`)
	rbacSysOwner := domain.NewRBACAppUser(sysOwnerID)
	rbacAllUserRolesObject := domain.NewRBACAllUserRolesObject(orgID)
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
	rbacOwnerGroup := domain.NewRBACUserRole(orgID, ownerGroupID)
	// - "owner" group "can" "set" "all-user-roles"
	err = authorizationManager.AddPolicyToGroupBySystemAdmin(ctx, sysAd, orgID, rbacOwnerGroup, service.RBACSetAction, rbacAllUserRolesObject, service.RBACAllowEffect)
	require.NoError(t, err)
	// - "owner" group "can" "unset" "all-user-roles"
	err = authorizationManager.AddPolicyToGroupBySystemAdmin(ctx, sysAd, orgID, rbacOwnerGroup, service.RBACUnsetAction, rbacAllUserRolesObject, service.RBACAllowEffect)
	require.NoError(t, err)

	// 6. add first owner
	ownerID, err := appUserRepo.AddAppUser(ctx, sysOwner, firstOwnerAddParam)
	require.NoError(t, err)
	require.Positive(t, ownerID.Int())

	// - owner belongs to owner-group
	err = authorizationManager.AddUserToGroup(ctx, sysOwner, ownerID, ownerGroupID)
	require.NoError(t, err)

	owner, err := appUserRepo.FindOwnerByLoginID(ctx, sysOwner, firstOwnerAddParam.LoginID())
	require.NoError(t, err)

	// logger := slog.Default()
	// logger.Warn(fmt.Sprintf("orgID: %d", orgID.Int()))

	return orgID, sysOwner, owner
}

func teardownOrganization(t *testing.T, ts testService, orgID *domain.OrganizationID) {
	t.Helper()
	// delete all organizations
	// ts.db.Exec("delete from space where organization_id = ?", orgID.Int())
	ts.db.Exec("delete from mb_app_user where organization_id = ?", orgID.Int())
	ts.db.Exec("delete from mb_organization where id = ?", orgID.Int())
	// db.Where("true").Delete(&spaceEntity{})
	// db.Where("true").Delete(&appUserEntity{})
	// db.Where("true").Delete(&organizationEntity{})
}

func testAddAppUser(t *testing.T, ctx context.Context, ts testService, owner service.OwnerModelInterface, loginID, username, password string) *service.AppUser {
	t.Helper()
	appUserRepo := ts.rf.NewAppUserRepository(ctx)
	userID1, err := appUserRepo.AddAppUser(ctx, owner, testNewAppUserAddParameter(t, loginID, username, password))
	require.NoError(t, err)
	user1, err := appUserRepo.FindAppUserByID(ctx, owner, userID1)
	require.NoError(t, err)
	require.Equal(t, loginID, user1.LoginID())

	return user1
}

func testAddUserGroup(t *testing.T, ctx context.Context, ts testService, owner service.OwnerModelInterface, key, name, description string) *service.UserGroup {
	t.Helper()
	userGorupRepo := ts.rf.NewUserGroupRepository(ctx)
	groupID1, err := userGorupRepo.AddUserGroup(ctx, owner, testNewUserGroupAddParameter(t, key, name, description))
	require.NoError(t, err)
	group1, err := userGorupRepo.FindUserGroupByID(ctx, owner, groupID1)
	require.NoError(t, err)
	require.Equal(t, key, group1.Key())
	require.Equal(t, name, group1.Name())
	require.Equal(t, description, group1.Description())

	return group1
}

type testSystemAdmin struct {
	*domain.SystemAdminModel
}

func (m *testSystemAdmin) AppUserID() *domain.AppUserID {
	return m.SystemAdminModel.AppUserID
}
func (m *testSystemAdmin) IsSystemAdmin() bool {
	return true
}
func testNewSystemAdmin(systemAdminModel *domain.SystemAdminModel) *testSystemAdmin {
	return &testSystemAdmin{
		systemAdminModel,
	}
}

type testAppUserModel struct {
	*domain.AppUserModel
}

func (m *testAppUserModel) AppUserID() *domain.AppUserID {
	return m.AppUserModel.AppUserID
}
func (m *testAppUserModel) OrganizationID() *domain.OrganizationID {
	return m.AppUserModel.OrganizationID
}
func (m *testAppUserModel) LoginID() string {
	return m.AppUserModel.LoginID
}
func (m *testAppUserModel) Username() string {
	return m.AppUserModel.Username
}
func testNewAppUser(appUserModel *domain.AppUserModel) *testAppUserModel {
	return &testAppUserModel{
		appUserModel,
	}
}

type testUserGroupModel struct {
	*domain.UserGroupModel
}

func (m *testUserGroupModel) Key() string {
	return m.UserGroupModel.Key
}
func (m *testUserGroupModel) Name() string {
	return m.UserGroupModel.Key
}
func (m *testUserGroupModel) Description() string {
	return m.UserGroupModel.Description
}
func testNewUserGroup(userGroupModel *domain.UserGroupModel) *testUserGroupModel {
	return &testUserGroupModel{
		userGroupModel,
	}
}
func testNewUserGroups(userGroupModels []*domain.UserGroupModel) []*testUserGroupModel {
	groups := make([]*testUserGroupModel, len(userGroupModels))
	for i, groupModel := range userGroupModels {
		groups[i] = testNewUserGroup(groupModel)
	}

	return groups
}

func testNewAppUserAddParameter(t *testing.T, loginID, username, password string) *service.AppUserAddParameter {
	t.Helper()
	p, err := service.NewAppUserAddParameter(loginID, username, password, "", "", "", "")
	require.NoError(t, err)

	return p
}

func testNewUserGroupAddParameter(t *testing.T, key, name, description string) *service.UserGroupAddParameter {
	t.Helper()
	p, err := service.NewUserGroupAddParameter(key, name, description)
	require.NoError(t, err)
	return p
}

func getOrganization(t *testing.T, ctx context.Context, ts testService, orgID *domain.OrganizationID) *service.Organization {
	t.Helper()
	orgRepo := gateway.NewOrganizationRepository(ctx, ts.db)

	baseModel, err := libdomain.NewBaseModel(1, time.Now(), time.Now(), 1, 1)
	require.NoError(t, err)
	appUserID, _ := domain.NewAppUserID(1)
	appUserModel, err := domain.NewAppUserModel(baseModel, appUserID, orgID, "login_id", "username", nil)
	require.NoError(t, err)
	appUser, err := service.NewAppUser(ctx, ts.rf, appUserModel)
	require.NoError(t, err)

	org, err := orgRepo.GetOrganization(ctx, appUser)
	require.NoError(t, err)
	require.Len(t, org.Name(), orgNameLength)

	return org
}
