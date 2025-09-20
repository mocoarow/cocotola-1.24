package usecase

import (
	"context"
	"log/slog"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	liblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	libservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type AddUserCommand struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	logger       *slog.Logger
}

func NewAddUserCommand(ctx context.Context, txManager service.TransactionManager, nonTxManager service.TransactionManager) (*AddUserCommand, error) {
	return &AddUserCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(liblog.LoggerNameKey, "AddUserCommand")),
	}, nil
}

func (u *AddUserCommand) Execute(ctx context.Context, operator service.OperatorInterface, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	u.logger.InfoContext(ctx, "AddStudent")
	fn := func(rf service.RepositoryFactory) (*domain.UserID, error) {
		userRepo := rf.NewUserRepository(ctx)
		userGroupRepo := rf.NewUserGroupRepository(ctx)
		authorizationManager, err := rf.NewAuthorizationManager(ctx)
		if err != nil {
			return nil, liberrors.Errorf("NewAuthorizationManager: %w", err)
		}

		// 1. add user
		userID, err := userRepo.AddUser(ctx, operator, param)
		if err != nil {
			return nil, liberrors.Errorf("m.userRepo.AddUser. err: %w", err)
		}

		// 2. add user to public-group
		publicGroup, err := userGroupRepo.FindUserGroupByKey(ctx, operator, service.PublicGroupKey)
		if err != nil {
			return nil, liberrors.Errorf("find public group(%s): %w", service.PublicGroupKey, err)
		}
		if err := authorizationManager.AddUserToGroup(ctx, operator, userID, publicGroup.UserGroupID); err != nil {
			return nil, liberrors.Errorf("AddUserToGroup: %w", err)
		}

		// 3. add policy to user
		subject := userID.GetRBACSubject()
		for _, aoe := range aoeList {
			if err := authorizationManager.AddPolicyToUser(ctx, operator, subject, aoe.Action, aoe.Object, aoe.Effect); err != nil {
				return nil, liberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
			}
		}

		return userID, nil
	}
	userID, err := libservice.Do1(ctx, u.txManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	fn2 := func(rf service.RepositoryFactory) error {
		userEventHandler := rf.NewUserEventHandler(ctx)
		userEventHandler.OnAdd(context.Background(), map[string]int{
			"organizationId": operator.GetOrganizationID().Int(),
			"userId":         userID.Int(),
		})
		return nil
	}
	if err := libservice.Do0(ctx, u.nonTxManager, fn2); err != nil {
		return nil, err //nolint:wrapcheck
	}
	return userID, nil
}
