package gateway

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

const conf = `
[request_definition]
r = sub, obj, act, dom

[policy_definition]
p = sub, obj, act, eft, dom

[role_definition]
g = _, _, _
g2 = _, _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub, r.dom) && (keyMatch(r.obj, p.obj) || g2(r.obj, p.obj, r.dom)) && r.act == p.act
`

type rbacRepository struct {
	DB       *gorm.DB
	Conf     string
	enforcer casbin.IEnforcer
}

func newRBACRepository(_ context.Context, db *gorm.DB) (service.RBACRepository, error) {
	if db == nil {
		panic(errors.New("db is nil"))
	}

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, liberrors.Errorf("gormadapter.NewAdapterByDB. err: %w", err)
	}

	m, err := model.NewModelFromString(conf)
	if err != nil {
		return nil, liberrors.Errorf("model.NewModelFromString. err: %w", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, liberrors.Errorf("casbin.NewEnforcer. err: %w", err)
	}

	return &rbacRepository{
		DB:       db,
		Conf:     conf,
		enforcer: e,
	}, nil
}

// func (r *rbacRepository) Init() error {
// 	a, err := gormadapter.NewAdapterByDB(r.DB)
// 	if err != nil {
// 		return liberrors.Errorf("gormadapter.NewAdapterByDB. err: %w", err)
// 	}

// 	m, err := model.NewModelFromString(r.Conf)
// 	if err != nil {
// 		return liberrors.Errorf("model.NewModelFromString. err: %w", err)
// 	}

// 	if err := a.SavePolicy(m); err != nil {
// 		return liberrors.Errorf(". err: %w", err)
// 	}

// 	return nil
// }

func (r *rbacRepository) initEnforcer(_ context.Context) (casbin.IEnforcer, error) {
	return r.enforcer, nil
	// logger := log.GetLoggerFromContext(ctx, UserGatewayContextKey)
	// a, err := gormadapter.NewAdapterByDB(r.DB)
	// if err != nil {
	// 	return nil, liberrors.Errorf("gormadapter.NewAdapterByDB. err: %w", err)
	// }

	// m, err := model.NewModelFromString(r.Conf)
	// if err != nil {
	// 	return nil, liberrors.Errorf("model.NewModelFromString. err: %w", err)
	// }

	// e, err := casbin.NewEnforcer(m, a)
	// if err != nil {
	// 	return nil, liberrors.Errorf("casbin.NewEnforcer. err: %w", err)
	// }

	// return e, nil
}

// p, alice, domain:1_data:1, read, allow, domain1
// p, bob, domain:2_data:2, write, allow, domain2
// p, bob, domain:1_data:2, write, allow, domain1
// p, charlie, domain:1_data*, read, allow, domain1
// p, domain:1_data2_admin, domain:1_data:2, read, allow, domain1
// p, domain:1_data2_admin, domain:1_data:2, write, allow, domain1

// g, alice, domain:1_data2_admin, domain1
// g2, domain:1_data_child, domain:1_data_parent, domain1
// g2, domain:2_data_child, domain:2_data_parent, domain2

func (r *rbacRepository) AddPolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err := e.AddNamedPolicy("p", subject.Subject(), object.Object(), action.Action(), effect.Effect(), domain.Domain()); err != nil {
		return liberrors.Errorf("e.AddNamedPolicy. err: %w", err)
	}

	return nil
}

func (r *rbacRepository) RemovePolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACSubject, action domain.RBACAction, object domain.RBACObject, effect domain.RBACEffect) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err = e.RemoveNamedPolicy("p", subject.Subject(), object.Object(), action.Action(), effect.Effect(), domain.Domain()); err != nil {
		return liberrors.Errorf("e.AddNamedPolicy. err: %w", err)
	}

	return nil
}

func (r *rbacRepository) RemoveSubjectPolicy(ctx context.Context, _ domain.RBACDomain, subject domain.RBACSubject) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err := e.RemoveFilteredNamedPolicy("p", 0, subject.Subject()); err != nil {
		return liberrors.Errorf("e.AddNamedPolicy. err: %w", err)
	}

	return nil
}

// func (r *rbacRepository) AddSubjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACUser, object domain.RBACRole) error {
func (r *rbacRepository) AddSubjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, child domain.RBACSubject, parent domain.RBACSubject) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err := e.AddNamedGroupingPolicy("g", child.Subject(), parent.Subject(), domain.Domain()); err != nil {
		return liberrors.Errorf("e.AddNamedGroupingPolicy. err: %w", err)
	}

	return nil
}

func (r *rbacRepository) RemoveSubjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, subject domain.RBACUser, object domain.RBACRole) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err := e.RemoveNamedGroupingPolicy("g", subject.Subject(), object.Role(), domain.Domain()); err != nil {
		return liberrors.Errorf("e.AddNamedGroupingPolicy. err: %w", err)
	}

	return nil
}

func (r *rbacRepository) AddObjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, child domain.RBACObject, parent domain.RBACObject) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err := e.AddNamedGroupingPolicy("g2", child.Object(), parent.Object(), domain.Domain()); err != nil {
		return liberrors.Errorf("e.AddNamedGroupingPolicy. err: %w", err)
	}

	return nil
}

func (r *rbacRepository) RemoveObjectGroupingPolicy(ctx context.Context, domain domain.RBACDomain, child domain.RBACObject, parent domain.RBACObject) error {
	e, err := r.initEnforcer(ctx)
	if err != nil {
		return liberrors.Errorf("r.initEnforcer. err: %w", err)
	}

	if _, err := e.RemoveNamedGroupingPolicy("g2", child.Object(), parent.Object(), domain.Domain()); err != nil {
		return liberrors.Errorf("e.AddNamedGroupingPolicy. err: %w", err)
	}

	return nil
}

func (r *rbacRepository) NewEnforcerWithGroupsAndUsers(_ context.Context, groups []domain.RBACRole, users []domain.RBACUser) (casbin.IEnforcer, error) {
	subjects := make([]string, 0)
	for _, s := range groups {
		subjects = append(subjects, s.Role())
	}
	for _, s := range users {
		subjects = append(subjects, s.Subject())
	}
	if err := r.enforcer.LoadFilteredPolicy(gormadapter.Filter{V0: subjects}); err != nil {
		return nil, liberrors.Errorf("e.LoadFilteredPolicy. err: %w", err)
	}

	return r.enforcer, nil
	// e, err := r.initEnforcer(ctx)
	// if err != nil {
	// 	return nil, liberrors.Errorf("r.initEnforcer. err: %w", err)
	// }
	// if err := e.LoadFilteredPolicy(gormadapter.Filter{V0: subjects}); err != nil {
	// 	return nil, liberrors.Errorf("e.LoadFilteredPolicy. err: %w", err)
	// }
	// return e, nil
}
func (r *rbacRepository) GetEnforcer() casbin.IEnforcer {
	return r.enforcer
}

// func (r *rbacRepository) CanDo(ctx context.Context, operatorID domain.AppUserID, ticketID domain.TicketID, action domain.RBACAction) (bool, error) {
// 	rbacRepo := r.rf.NewRBACRepository(ctx)

// 	roleObjects := r.getAllRolesForTicket(ticketID)
// 	userObject := NewRBACAppUser(operatorID)
// 	e, err := rbacRepo.NewEnforcerWithRolesAndUsers(roleObjects, []domain.RBACUser{userObject})
// 	if err != nil {
// 		return false, liberrors.Errorf("failed to NewEnforcerWithRolesAndUsers. err: %w", err)
// 	}

// 	ticketObject := NewRBACTicketObject(ticketID)

// 	ok, err := e.Enforce(string(userObject), string(ticketObject), string(action))
// 	if err != nil {
// 		return false, liberrors.Errorf("e.Enforce. err: %w", err)
// 	}

// 	return ok, nil
// }
