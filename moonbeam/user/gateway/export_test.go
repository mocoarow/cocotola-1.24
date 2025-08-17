package gateway

import (
	"context"

	"github.com/casbin/casbin/v2"
)

type RBACRepository = rbacRepository

var NewRBACRepository = newRBACRepository
var Conf = conf

type OrganizationEntity = organizationEntity

func (r *rbacRepository) InitEnforcer(ctx context.Context) (casbin.IEnforcer, error) {
	return r.initEnforcer(ctx)
}
