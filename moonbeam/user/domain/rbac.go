package domain

import (
	"fmt"
)

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

type RBACActionObjectEffect struct {
	Action RBACAction
	Object RBACObject
	Effect RBACEffect
}

func NewRBACOrganization(organizationID *OrganizationID) RBACDomain {
	return NewRBACDomain(fmt.Sprintf("domain:%d", organizationID.Int()))
}

func NewRBACAppUser(appUserID *AppUserID) RBACUser {
	return NewRBACUser(fmt.Sprintf("user:%d", appUserID.Int()))
}

func NewRBACUserRole(organizationID *OrganizationID, userGroupID *UserGroupID) RBACRole {
	return NewRBACRole(fmt.Sprintf("domain:%d,role:%d", organizationID.Int(), userGroupID.Int()))
}

func NewRBACUserRoleObject(organizationID *OrganizationID, userRoleID *UserGroupID) RBACObject {
	return NewRBACObject(fmt.Sprintf("domain:%d,role:%d", organizationID.Int(), userRoleID.Int()))
}

func NewRBACAllUserRolesObject(organizationID *OrganizationID) RBACObject {
	return NewRBACObject(fmt.Sprintf("domain:%d,role:*", organizationID.Int()))
}
