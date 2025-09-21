package service

import (
	"context"
	"fmt"
	"log/slog"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

var _ SystemAdminInterface = (*SystemAdmin)(nil)

type SystemAdminInterface interface {
	GetUserID() *domain.UserID
	IsSystemAdmin() bool
	// GetUserGroups() []domain.UserGroupModel
}

type SystemAdmin struct {
	*domain.SystemAdminModel
	rf       RepositoryFactory
	orgRepo  OrganizationRepository
	userRepo UserRepository
	logger   *slog.Logger
}

func NewSystemAdmin(ctx context.Context, rf RepositoryFactory) (*SystemAdmin, error) {
	if rf == nil {
		return nil, fmt.Errorf("new system admin. argument 'rf' is nil: %w", libdomain.ErrInvalidArgument)
	}
	orgRepo := rf.NewOrganizationRepository(ctx)
	userRepo := rf.NewUserRepository(ctx)

	m := &SystemAdmin{
		SystemAdminModel: domain.NewSystemAdminModel(),
		rf:               rf,
		orgRepo:          orgRepo,
		userRepo:         userRepo,
		logger:           slog.Default().With(slog.String(liblog.LoggerNameKey, "SystemAdmin")),
	}

	return m, nil
}

func (m *SystemAdmin) GetUserID() *domain.UserID {
	return m.UserID
}
func (m *SystemAdmin) IsSystemAdmin() bool {
	return true
}

// func (m *SystemAdmin) FindSystemOwnerByOrganizationID(ctx context.Context, organizationID *domain.OrganizationID) (*SystemOwner, error) {
// 	sysOwner, err := m.userRepo.FindSystemOwnerByOrganizationID(ctx, m, organizationID)
// 	if err != nil {
// 		return nil, liberrors.Errorf("m.userRepo.FindSystemOwnerByOrganizationID. error: %w", err)
// 	}

// 	return sysOwner, nil
// }

// func (m *SystemAdmin) FindSystemOwnerByOrganizationName(ctx context.Context, organizationName string) (*SystemOwner, error) {
// 	sysOwner, err := m.userRepo.FindSystemOwnerByOrganizationName(ctx, m, organizationName)
// 	if err != nil {
// 		return nil, liberrors.Errorf("find system owner by organization name: %w", err)
// 	}

// 	return sysOwner, nil
// }

func (m *SystemAdmin) FindOrganizationByID(ctx context.Context, organizationID *domain.OrganizationID) (*Organization, error) {
	org, err := m.orgRepo.FindOrganizationByID(ctx, m, organizationID)
	if err != nil {
		return nil, liberrors.Errorf("m.orgRepo.FindOrganizationByID: %w", err)
	}

	return org, nil
}

func (m *SystemAdmin) FindOrganizationByName(ctx context.Context, name string) (*Organization, error) {
	org, err := m.orgRepo.FindOrganizationByName(ctx, m, name)
	if err != nil {
		return nil, liberrors.Errorf("m.orgRepo.FindOrganizationByName. error: %w", err)
	}

	return org, nil
}

var RBACSetAction = domain.NewRBACAction("Set")
var RBACUnsetAction = domain.NewRBACAction("Unset")

var RBACAllowEffect = domain.NewRBACEffect("allow")
var RBACDenyEffect = domain.NewRBACEffect("deny")
