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
	action := mbuserdomain.NewRBACAction("CreateUser")
	object := mbuserdomain.NewRBACObject("*")
	ok, err := service.CheckAuthorization(ctx, operator, action, object, u.mbNonTxManager)
	if err != nil {
		return nil, mbliberrors.Errorf("authorize: %w", err)
	} else if !ok {
		u.logger.InfoContext(ctx, fmt.Sprintf("operator(%d) is not authorized to create user", operator.GetUserID().Int()))

		return nil, domain.ErrUnauthenticated
	}

	// 1. Check if the user already exists
	existingUser, err := u.findUserbyLoginID(ctx, operator, param.LoginID)
	if err != nil && !errors.Is(err, mbuserservice.ErrUserNotFound) {
		return nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	existingOrg, err := u.getOrganization(ctx, operator)
	if err != nil {
		return nil, mbliberrors.Errorf("getOrganization: %w", err)
	}
	if existingUser != nil {
		u.logger.InfoContext(ctx, fmt.Sprintf("user with loginID(%s) already exists", param.LoginID))
		tokenSet, err := u.createTokenSet(ctx, existingUser, existingOrg)
		if err != nil {
			return nil, mbliberrors.Errorf("create token set: %w", err)
		}
		return tokenSet, nil
	}

	// 2. add user
	createUserParameterFunc := func() (*mbuserservice.AddUserParameter, error) {
		return param, nil
	}
	parameter, err := createUserParameterFunc()
	if err != nil {
		return nil, mbliberrors.Errorf("invalid UserAddParameter. err: %w", err)
	}
	command, err := mbuserusecase.NewAddUserCommand(ctx, u.mbTxManager, u.mbNonTxManager)
	if err != nil {
		return nil, mbliberrors.Errorf("NewAddUserCommand. err: %w", err)
	}
	if _, err := command.Execute(ctx, operator, parameter, []mbuserusecase.ActionObjectEffect{}); err != nil {
		return nil, mbliberrors.Errorf("AddUserCommand.Execute. err: %w", err)
	}

	// 3. create auth token set
	newUser, err := u.findUserbyLoginID(ctx, operator, param.LoginID)
	if err != nil {
		return nil, mbliberrors.Errorf("findUserbyLoginID: %w", err)
	}
	tokenSet, err := u.createTokenSet(ctx, newUser, existingOrg)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return tokenSet, nil
}

func (u *RegisterUserCommand) findUserbyLoginID(ctx context.Context, operator mbuserservice.OperatorInterface, loginID string) (*mbuserdomain.UserModel, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.UserModel, error) {
		userRepo := mbrf.NewUserRepository(ctx)

		user, err := userRepo.FindUserByLoginID(ctx, operator, loginID)
		if err != nil {
			return nil, mbliberrors.Errorf("FindUserByLoginID(%s). err: %w", loginID, err)
		}
		return user.UserModel, nil
	}
	userModel, err := mblibservice.Do1(ctx, u.mbTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return userModel, nil
}

func (u *RegisterUserCommand) getOrganization(ctx context.Context, operator mbuserservice.OperatorInterface) (*mbuserdomain.OrganizationModel, error) {
	fn := func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.OrganizationModel, error) {
		orgRepo := mbrf.NewOrganizationRepository(ctx)
		org, err := orgRepo.GetOrganization(ctx, operator)
		if err != nil {
			return nil, mbliberrors.Errorf("GetOrganization(). err: %w", err)
		}

		return org.OrganizationModel, nil
	}
	orgModel, err := mblibservice.Do1(ctx, u.mbNonTxManager, fn)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return orgModel, nil
}

func (u *RegisterUserCommand) createTokenSet(ctx context.Context, userModel *mbuserdomain.UserModel, organizationModel *mbuserdomain.OrganizationModel) (*domain.AuthTokenSet, error) {
	targetUser := &user{
		userID:         userModel.UserID,
		organizationID: userModel.OrganizationID,
		loginID:        userModel.LoginID,
		username:       userModel.Username,
	}
	targetOorganization := &organization{
		organizationID: organizationModel.OrganizationID,
		name:           organizationModel.Name,
	}

	tokenSet, err := u.authTokenManager.CreateTokenSet(ctx, targetUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}
	return tokenSet, nil
}
