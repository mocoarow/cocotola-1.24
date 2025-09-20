package usecase

import (
	"context"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type AddGuestCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
}

func NewAddGuestCommand(txManager service.TransactionManager, nonTxManager service.TransactionManager) *AddGuestCommand {
	return &AddGuestCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
	}
}

func (u *AddGuestCommand) Execute(ctx context.Context, operator service.SystemOwnerInterface, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	fn2 := func(rf service.RepositoryFactory) (*domain.UserID, error) {
		userRepo := rf.NewUserRepository(ctx)
		userGroupRepo := rf.NewUserGroupRepository(ctx)
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return nil, liberrors.Errorf("failed to NewAuthorizationManager: %w", err)
		}

		// 1. add guest
		guestID, err := userRepo.AddUser(ctx, operator, param)
		if err != nil {
			return nil, liberrors.Errorf("AddUser: %w", err)
		}

		// 2. add guest to public-group
		publicGroup, err := userGroupRepo.FindUserGroupByKey(ctx, operator, service.PublicGroupKey)
		if err != nil {
			return nil, liberrors.Errorf("find public group(%s): %w", service.PublicGroupKey, err)
		}
		if err := authorizationManager.AddUserToGroup(ctx, operator, guestID, publicGroup.UserGroupID); err != nil {
			return nil, liberrors.Errorf("AddUserToGroup: %w", err)
		}

		// 3. add policy to user
		subject := guestID.GetRBACSubject()
		for _, aoe := range aoeList {
			if err := authorizationManager.AddPolicyToUserBySystemOwner(ctx, operator, subject, aoe.Action, aoe.Object, aoe.Effect); err != nil {
				return nil, liberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
			}
		}

		return guestID, nil
	}
	guestID, err := libservice.Do1(ctx, u.nonTxManager, fn2)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return guestID, nil
}
