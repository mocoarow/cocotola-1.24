package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
	mbuserusecase "github.com/mocoarow/cocotola-1.24/moonbeam/user/usecase"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type RegisterUserCommand struct {
	mbTxManager      mbuserservice.TransactionManager
	mbNonTxManager   mbuserservice.TransactionManager
	authTokenManager service.AuthTokenManager
	logger           *slog.Logger
}

func NewRegisterUserCommand(ctx context.Context, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, authTokenManager service.AuthTokenManager) (*RegisterUserCommand, error) {
	return &RegisterUserCommand{
		mbTxManager:      mbTxManager,
		mbNonTxManager:   mbNonTxManager,
		authTokenManager: authTokenManager,
		logger:           slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-RegisterUserCommand")),
	}, nil
}

func (u *RegisterUserCommand) Execute(ctx context.Context, operator mbuserservice.OperatorInterface, param *mbuserservice.AddUserParameter) (*domain.AuthTokenSet, error) {
	// 1. Check authorization
	if err := u.checkAuthorization(ctx, operator); err != nil {
		return nil, mbliberrors.Errorf("checkAuthorization: %w", err)
	}

	// 2. Execute
	newUserID, tokenSet, err := u.execute(ctx, operator, param)
	if err != nil {
		return nil, mbliberrors.Errorf("execute: %w", err)
	}

	// 3. Callback
	if err := u.callback(ctx, operator, newUserID); err != nil {
		return nil, mbliberrors.Errorf("callback: %w", err)
	}

	return tokenSet, nil
}

func (u *RegisterUserCommand) checkAuthorization(ctx context.Context, operator mbuserservice.OperatorInterface) error {
	action := mbuserdomain.NewRBACAction("CreateUser")
	object := mbuserdomain.NewRBACObject("*")
	ok, err := service.CheckAuthorization(ctx, operator, action, object, u.mbNonTxManager)
	if err != nil {
		return mbliberrors.Errorf("authorize: %w", err)
	} else if !ok {
		u.logger.InfoContext(ctx, fmt.Sprintf("operator(%d) is not authorized to create user", operator.GetUserID().Int()))

		return domain.ErrUnauthenticated
	}
	return nil
}

func (u *RegisterUserCommand) execute(ctx context.Context, operator mbuserservice.OperatorInterface, param *mbuserservice.AddUserParameter) (*mbuserdomain.UserID, *domain.AuthTokenSet, error) {
	// 1. Check if the user already exists
	existingUser, err := u.findUserbyLoginID(ctx, operator, param.LoginID)
	if err != nil && !errors.Is(err, mbuserservice.ErrUserNotFound) {
		return nil, nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	existingOrg, err := u.getOrganization(ctx, operator)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("getOrganization: %w", err)
	}
	if existingUser != nil {
		u.logger.InfoContext(ctx, fmt.Sprintf("user with loginID(%s) already exists", param.LoginID))
		tokenSet, err := u.createTokenSet(ctx, existingUser, existingOrg)
		if err != nil {
			return nil, nil, mbliberrors.Errorf("create token set: %w", err)
		}
		return existingUser.UserID, tokenSet, nil
	}

	// 2. add user
	aoeList := []mbuserusecase.ActionObjectEffect{}
	if _, err := mblibservice.Do1(ctx, u.mbTxManager, func(rf mbuserservice.RepositoryFactory) (*mbuserdomain.UserID, error) {
		return mbuserusecase.AddUser(ctx, operator, rf, param, aoeList)
	}); err != nil {
		return nil, nil, err //nolint:wrapcheck
	}

	// 3. create auth token set
	newUser, err := u.findUserbyLoginID(ctx, operator, param.LoginID)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	tokenSet, err := u.createTokenSet(ctx, newUser, existingOrg)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return newUser.UserID, tokenSet, nil
}

func (u *RegisterUserCommand) callback(ctx context.Context, operator mbuserservice.OperatorInterface, newUserID *mbuserdomain.UserID) error {
	fn := func(rf mbuserservice.RepositoryFactory) error {
		userEventHandler := rf.NewUserEventHandler(ctx)
		userEventHandler.OnAdd(context.Background(), map[string]int{
			"organizationId": operator.GetOrganizationID().Int(),
			"userId":         newUserID.Int(),
		})
		return nil
	}
	if err := mblibservice.Do0(ctx, u.mbNonTxManager, fn); err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}

func (u *RegisterUserCommand) findUserbyLoginID(ctx context.Context, operator mbuserservice.OperatorInterface, loginID string) (*mbuserdomain.UserModel, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.UserModel, error) {
		return findUserbyLoginID(ctx, mbrf, operator, loginID)
	})
}

func (u *RegisterUserCommand) getOrganization(ctx context.Context, operator mbuserservice.OperatorInterface) (*mbuserdomain.OrganizationModel, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.OrganizationModel, error) {
		return getOrganization(ctx, mbrf, operator)
	})
}

func (u *RegisterUserCommand) createTokenSet(ctx context.Context, userModel *mbuserdomain.UserModel, organizationModel *mbuserdomain.OrganizationModel) (*domain.AuthTokenSet, error) {
	return createTokenSet(ctx, u.authTokenManager, userModel, organizationModel)
}
