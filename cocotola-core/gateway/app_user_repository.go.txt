package gateway

import (
	"context"

	"gorm.io/gorm"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

var _ service.AppUserRepository = (*AppUserRepository)(nil)

type AppUserRepository struct {
	db         *gorm.DB
	rf         service.RepositoryFactory
	rbacClient libapi.CocotolaRBACClient
}

func NewAppUserRepository(db *gorm.DB, rf service.RepositoryFactory, rbacClient libapi.CocotolaRBACClient) *AppUserRepository {
	return &AppUserRepository{
		db:         db,
		rf:         rf,
		rbacClient: rbacClient,
	}
}

func (r *AppUserRepository) NewAppUser(_ context.Context, operator mbuserservice.OperatorInterface) (*service.AppUser, error) {
	return service.NewAppUser(operator, r.rf, r.rbacClient), nil
}
