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

func NewAddUserCommand(ctx context.Context, txManager service.TransactionManager, nonTxManager service.TransactionManager) *AddUserCommand {
	return &AddUserCommand{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		logger:       slog.Default().With(slog.String(liblog.LoggerNameKey, "AddUserCommand")),
	}
}

func (u *AddUserCommand) Execute(ctx context.Context, operator domain.UserInterface, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	u.logger.InfoContext(ctx, "AddStudent")

	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator); err != nil {
		return nil, liberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	newUserID, err := u.execute(ctx, operator, param, aoeList)
	if err != nil {
		return nil, liberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(ctx, operator, newUserID); err != nil {
		return nil, liberrors.Errorf("callback: %w", err)
	}

	return newUserID, nil
}

func (u *AddUserCommand) checkAuthorization(ctx context.Context, operator domain.UserInterface) error {
	return nil
}

func (u *AddUserCommand) execute(ctx context.Context, operator domain.UserInterface, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	u.logger.InfoContext(ctx, "AddStudent")
	userID, err := libservice.Do1(ctx, u.txManager, func(rf service.RepositoryFactory) (*domain.UserID, error) {
		return AddUser(ctx, operator, rf, param, aoeList)
	})
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return userID, nil
}

func (u *AddUserCommand) callback(ctx context.Context, operator domain.UserInterface, newUserID *domain.UserID) error {
	fn := func(rf service.RepositoryFactory) error {
		userEventHandler := rf.NewUserEventHandler(ctx)
		userEventHandler.OnAdd(context.Background(), map[string]int{
			"organizationId": operator.GetOrganizationID().Int(),
			"userId":         newUserID.Int(),
		})
		return nil
	}
	if err := libservice.Do0(ctx, u.nonTxManager, fn); err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}

func AddUser(ctx context.Context, operator domain.UserInterface, rf service.RepositoryFactory, param *service.AddUserParameter, aoeList []ActionObjectEffect) (*domain.UserID, error) {
	userRepo := rf.NewUserRepository(ctx)
	userGroupRepo := rf.NewUserGroupRepository(ctx)
	authorizationManager, err := rf.NewAuthorizationManager(ctx)
	if err != nil {
		return nil, liberrors.Errorf("NewAuthorizationManager: %w", err)
	}

	// 1. create new user
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

	// 3. attach policy to user
	subject := userID.GetRBACSubject()
	for _, aoe := range aoeList {
		if err := authorizationManager.AddPolicyToUser(ctx, operator, subject, aoe.Action, aoe.Object, aoe.Effect); err != nil {
			return nil, liberrors.Errorf("AddPolicyToUserBySystemAdmin: %w", err)
		}
	}

	return userID, nil
}
