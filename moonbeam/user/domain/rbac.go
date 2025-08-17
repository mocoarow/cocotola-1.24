package domain

// type RBACUser string
// type RBACRole string
// type RBACObject string
// type RBACAction string

type RBACSubject interface {
	Subject() string
}

type RBACUser interface {
	RBACSubject
}

type rbacUser struct {
	value string
}

func NewRBACUser(value string) RBACUser {
	return &rbacUser{value: value}
}

func (r *rbacUser) Subject() string {
	return r.value
}

type RBACRole interface {
	RBACSubject
	Role() string
}

type rbacRole struct {
	value string
}

func NewRBACRole(value string) RBACRole {
	return &rbacRole{value: value}
}

func (r *rbacRole) Subject() string {
	return r.value
}
func (r *rbacRole) Role() string {
	return r.value
}

type RBACDomain interface {
	Domain() string
}

type rbacDomain struct {
	value string
}

func NewRBACDomain(value string) RBACDomain {
	return &rbacDomain{value: value}
}

func (r *rbacDomain) Domain() string {
	return r.value
}

type RBACObject interface {
	Object() string
}

type rbacObject struct {
	value string
}

func NewRBACObject(value string) RBACObject {
	return &rbacObject{value: value}
}

func (r *rbacObject) Object() string {
	return r.value
}

type RBACAction interface {
	Action() string
}

type rbacAction struct {
	value string
}

func NewRBACAction(value string) RBACAction {
	return &rbacAction{value: value}
}

func (r *rbacAction) Action() string {
	return r.value
}

type RBACEffect interface {
	Effect() string
}

type rbacEffect struct {
	value string
}

func NewRBACEffect(value string) RBACEffect {
	return &rbacEffect{value: value}
}

func (r *rbacEffect) Effect() string {
	return r.value
}
