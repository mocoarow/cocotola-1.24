package usecase

import (
	"context"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type AddFirstOwnerCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewAddFirstOwnerCommand(txManager service.TransactionManager, nonTxManager service.TransactionManager) *AddFirstOwnerCommand {
	return &AddFirstOwnerCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *AddFirstOwnerCommand) Execute(ctx context.Context, operator domain.SystemOwnerInterface, param *service.AddUserParameter) (*domain.UserID, error) {
	fn1 := func(rf service.RepositoryFactory) error {
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return liberrors.Errorf("failed to NewAuthorizationManager: %w", err)
		}
		rbacAllUserRolesObject := domain.NewRBACAllUserRolesObjectFromOrganization(operator.GetOrganizationID())

		// Can "the operator" "set" "all-user-roles" ?
		ok, err := authorizationManager.CheckAuthorization(ctx, operator, service.RBACSetAction, rbacAllUserRolesObject)
		if err != nil {
			return liberrors.Errorf("CheckAuthorization: %w", err)
		} else if !ok {
			return libdomain.ErrPermissionDenied
		}
		return nil
	}
	if err := libservice.Do0(ctx, u.txManager, fn1); err != nil {
		return nil, err //nolint:wrapcheck
	}

	fn2 := func(rf service.RepositoryFactory) (*domain.UserID, error) {
		userRepo := rf.NewUserRepository(ctx)
		userGroupRepo := rf.NewUserGroupRepository(ctx)
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return nil, liberrors.Errorf("failed to NewAuthorizationManager: %w", err)
		}
		// 1. add owner
		firstOwnerID, err := userRepo.AddUser(ctx, operator, param)
		if err != nil {
			return nil, liberrors.Errorf("AddUser: %w", err)
		}

		ownerGroup, err := userGroupRepo.FindUserGroupByKey(ctx, operator, service.OwnerGroupKey)
		if err != nil {
			return nil, liberrors.Errorf("FindUserGroupByKey: %w", err)
		}

		// 2. add owner to owner-group
		if err := authorizationManager.AddUserToGroup(ctx, operator, firstOwnerID, ownerGroup.UserGroupID); err != nil {
			return nil, liberrors.Errorf("AddUserToGroup: %w", err)
		}

		// 3. add policy to "first-owner" user
		firstOwner, err := userRepo.FindUserByID(ctx, operator, firstOwnerID)
		if err != nil {
			return nil, liberrors.Errorf("FindUserByLoginID: %w", err)
		}

		// first owner can create users
		subject := firstOwner.GetUserID().GetRBACSubject()
		action := domain.NewRBACAction("CreateUser")
		object := domain.NewRBACObject("*")
		effect := service.RBACAllowEffect

		if err := authorizationManager.AddPolicyToUserBySystemOwner(ctx, operator, subject, action, object, effect); err != nil {
			return nil, liberrors.Errorf("AddPolicyToUserBySystemOwner: %w", err)
		}
		return firstOwnerID, nil
	}
	firstOwnerID, err := libservice.Do1(ctx, u.nonTxManager, fn2)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return firstOwnerID, nil
}
