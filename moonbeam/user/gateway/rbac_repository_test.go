//go:build medium

package gateway_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type testSDOA struct {
	subject string
	domain  string
	object  string
	action  string
	want    bool
}

func (t *testSDOA) String() string {
	return fmt.Sprintf("%s,%s,%s,%s,%v", t.subject, t.domain, t.object, t.action, t.want)
}

func addPolicy(t *testing.T, ctx context.Context, rbacRepository service.RBACRepository, dom, sub, act, obj string, allowed bool) {
	t.Helper()
	effect := service.RBACAllowEffect
	if !allowed {
		effect = service.RBACDenyEffect
	}

	err := rbacRepository.AddPolicy(ctx, domain.NewRBACDomain(dom), domain.NewRBACUser(sub), domain.NewRBACAction(act), domain.NewRBACObject(obj), effect)
	require.NoError(t, err)
}

func addObjectGroupingPolicy(t *testing.T, ctx context.Context, rbacRepository service.RBACRepository, dom, child, parent string) {
	t.Helper()
	err := rbacRepository.AddObjectGroupingPolicy(ctx, domain.NewRBACDomain(dom), domain.NewRBACObject(child), domain.NewRBACObject(parent))
	require.NoError(t, err)
}

func addSubjectGroupingPolicy(t *testing.T, ctx context.Context, rbacRepository service.RBACRepository, dom, sub, obj string) {
	t.Helper()
	err := rbacRepository.AddSubjectGroupingPolicy(ctx, domain.NewRBACDomain(dom), domain.NewRBACUser(sub), domain.NewRBACRole(obj))
	require.NoError(t, err)
}

func TestA(t *testing.T) {
	t.Parallel()

	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		defer teardownCasbin(t, ts)
		// rbacRepo := gateway.RBACRepository{
		// 	DB:   ts.db,
		// 	Conf: gateway.Conf,
		// }
		rbacRepo, err := gateway.NewRBACRepository(ctx, ts.db)
		require.NoError(t, err)
		e := rbacRepo.GetEnforcer()

		// rbacRepo.Init()

		// err := initRBACRepository(t, ts.db, gateway.Conf)
		// require.NoError(t, err)
		addPolicy(t, ctx, rbacRepo, "domain1", "alice", "read", "domain:1,data:1", true)
		addPolicy(t, ctx, rbacRepo, "domain1", "bob", "write", "domain:1,data:2", true)
		// rbacRepo.AddPolicy(domain.NewRBACDomain("domain1"), domain.NewRBACUser("alice"), domain.NewRBACAction("write"), domain.NewRBACObject("data1"), service.RBACAllowEffect)
		addObjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "domain:1,child:1", "domain:1,data:1")

		tests := []testSDOA{
			{subject: "alice", domain: "domain1", object: "domain:1,data:1", action: "read", want: true},
			{subject: "alice", domain: "domain1", object: "domain:1,data:1", action: "write", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,data:2", action: "read", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,data:2", action: "write", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,child:1", action: "read", want: true},
			{subject: "alice", domain: "domain1", object: "domain:1,child:1", action: "write", want: false},

			{subject: "bob", domain: "domain1", object: "domain:1,data:1", action: "read", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,data:1", action: "write", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,data:2", action: "read", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,data:2", action: "write", want: true},
			{subject: "bob", domain: "domain1", object: "domain:1,child:1", action: "read", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,child:1", action: "write", want: false},
			// {subject: "bob", domain: "domain1", object: "domain:1_data:1", action: "write", want: false},
			// {subject: "bob", domain: "domain1", object: "domain:1_data:2", action: "read", want: false},
			// {subject: "bob", domain: "domain1", object: "domain:1_data:2", action: "write", want: true},

			// {subject: "charlie", domain: "domain1", object: "domain:1_data:2", action: "read", want: true},
			// {subject: "charlie", domain: "domain1", object: "domain:1_data_parent", action: "read", want: true},
		}
		// e := initEnforcer(t, ts.db, gateway.Conf)
		// e, err := rbacRepo.InitEnforcer(ctx)
		require.NoError(t, err)
		for _, tt := range tests {
			t.Run(tt.String(), func(t *testing.T) {
				t.Parallel()
				ok, err := e.Enforce(tt.subject, tt.object, tt.action, tt.domain)
				require.NoError(t, err)
				assert.Equal(t, tt.want, ok)
			})
		}
	}
	testDB(t, fn)
}

func TestB(t *testing.T) {
	t.Parallel()

	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		defer teardownCasbin(t, ts)
		rbacRepo, err := gateway.NewRBACRepository(ctx, ts.db)
		require.NoError(t, err)
		e := rbacRepo.GetEnforcer()

		addSubjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "alice", "domain:1,reader")
		addSubjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "bob", "domain:1,writer")

		addPolicy(t, ctx, rbacRepo, "domain1", "domain:1,reader", "read", "domain:1,data:1", true)
		addPolicy(t, ctx, rbacRepo, "domain1", "domain:1,writer", "write", "domain:1,data:2", true)

		addObjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "domain:1,child:1", "domain:1,data:1")

		tests := []testSDOA{
			{subject: "alice", domain: "domain1", object: "domain:1,data:1", action: "read", want: true},
			{subject: "alice", domain: "domain1", object: "domain:1,data:1", action: "write", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,data:2", action: "read", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,data:2", action: "write", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,child:1", action: "read", want: true},
			{subject: "alice", domain: "domain1", object: "domain:1,child:1", action: "write", want: false},

			{subject: "bob", domain: "domain1", object: "domain:1,data:1", action: "read", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,data:1", action: "write", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,data:2", action: "read", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,data:2", action: "write", want: true},
			{subject: "bob", domain: "domain1", object: "domain:1,child:1", action: "read", want: false},
			{subject: "bob", domain: "domain1", object: "domain:1,child:1", action: "write", want: false},
			// {subject: "bob", domain: "domain1", object: "domain:1_data:1", action: "write", want: false},
			// {subject: "bob", domain: "domain1", object: "domain:1_data:2", action: "read", want: false},
			// {subject: "bob", domain: "domain1", object: "domain:1_data:2", action: "write", want: true},

			// {subject: "charlie", domain: "domain1", object: "domain:1_data:2", action: "read", want: true},
			// {subject: "charlie", domain: "domain1", object: "domain:1_data_parent", action: "read", want: true},
		}
		for _, tt := range tests {
			t.Run(tt.String(), func(t *testing.T) {
				t.Parallel()
				ok, err := e.Enforce(tt.subject, tt.object, tt.action, tt.domain)
				require.NoError(t, err)
				assert.Equal(t, tt.want, ok)
			})
		}
	}
	testDB(t, fn)
}

func TestC(t *testing.T) {
	t.Parallel()

	fn := func(t *testing.T, ctx context.Context, ts testService) {
		t.Helper()
		defer teardownCasbin(t, ts)
		rbacRepo, err := gateway.NewRBACRepository(ctx, ts.db)
		require.NoError(t, err)
		e := rbacRepo.GetEnforcer()

		addSubjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "alice", "domain:1,reader")

		addPolicy(t, ctx, rbacRepo, "domain1", "domain:1,reader", "read", "domain:1,data:2", true)
		addPolicy(t, ctx, rbacRepo, "domain1", "domain:1,reader", "read", "domain:1,data:4", false)

		addObjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "domain:1,data:2", "domain:1,data:1")
		addObjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "domain:1,data:3", "domain:1,data:2")
		addObjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "domain:1,data:4", "domain:1,data:3")
		addObjectGroupingPolicy(t, ctx, rbacRepo, "domain1", "domain:1,data:5", "domain:1,data:4")
		// 1/
		// - 2/ <= alice can read
		//   - 3/ <= alice also can read
		//     - 4/ <= alice can't read
		//	     - 5/ <= alice also can't read

		tests := []testSDOA{
			{subject: "alice", domain: "domain1", object: "domain:1,data:1", action: "read", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,data:2", action: "read", want: true},
			{subject: "alice", domain: "domain1", object: "domain:1,data:3", action: "read", want: true},
			{subject: "alice", domain: "domain1", object: "domain:1,data:4", action: "read", want: false},
			{subject: "alice", domain: "domain1", object: "domain:1,data:5", action: "read", want: false},
		}
		for _, tt := range tests {
			t.Run(tt.String(), func(t *testing.T) {
				t.Parallel()
				ok, err := e.Enforce(tt.subject, tt.object, tt.action, tt.domain)
				require.NoError(t, err)
				assert.Equal(t, tt.want, ok)
			})
		}
	}
	testDB(t, fn)
}

func teardownCasbin(t *testing.T, ts testService) {
	t.Helper()
	// delete all organizations
	// ts.db.Exec("delete from space where organization_id = ?", orgID.Int())
	ts.db.Exec("delete from casbin_rule")
	// db.Where("true").Delete(&spaceEntity{})
	// db.Where("true").Delete(&appUserEntity{})
	// db.Where("true").Delete(&organizationEntity{})
}
