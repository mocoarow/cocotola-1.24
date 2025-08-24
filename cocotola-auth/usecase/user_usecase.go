package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type UserUsecase struct {
	systemToken      libdomain.SystemToken
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
	logger           *slog.Logger
}

func NewUserUsecase(systemToken libdomain.SystemToken, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager) *UserUsecase {
	return &UserUsecase{
		systemToken:      systemToken,
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
		logger:           slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-UserUsecase")),
	}
}

func (u *UserUsecase) RegisterAppUser(ctx context.Context, operator service.OperatorInterface, param *mbuserservice.AppUserAddParameter) (*domain.AuthTokenSet, error) {
	action := mbuserdomain.NewRBACAction("CreateAppUser")
	object := mbuserdomain.NewRBACObject("*")
	ok, err := service.CheckAuthorization(ctx, operator, action, object, u.nonTxManager)
	if err != nil {
		return nil, mbliberrors.Errorf("authorize: %w", err)
	} else if !ok {
		u.logger.InfoContext(ctx, "operator is not authorized to create app user")

		return nil, domain.ErrUnauthenticated
	}

	createAppUserParameterFunc := func() (*mbuserservice.AppUserAddParameter, error) {
		return param, nil
	}

	var targetOorganization *organization
	var targetAppUser *appUser
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		tmpOrganization, tmpAppUser, err := registerAppUser(ctx, u.systemToken, rf, operator.OrganizationID(), param.LoginID(), createAppUserParameterFunc)
		if err != nil && !errors.Is(err, mbuserservice.ErrAppUserAlreadyExists) {
			return mbliberrors.Errorf("register app user: %w", err)
		} else if errors.Is(err, mbuserservice.ErrAppUserAlreadyExists) {
			return mbuserservice.ErrAppUserAlreadyExists
		}

		u.logger.InfoContext(ctx, fmt.Sprintf("tmpOrganization: %d", tmpOrganization.OrganizationID))
		u.logger.InfoContext(ctx, fmt.Sprintf("tmpAppUser: %d", tmpAppUser.AppUserID))

		targetAppUser = &appUser{
			appUserID:      tmpAppUser.AppUserID,
			organizationID: tmpAppUser.OrganizationID,
			loginID:        tmpAppUser.LoginID,
			username:       tmpAppUser.Username,
		}
		targetOorganization = &organization{
			organizationID: tmpOrganization.OrganizationID,
			name:           tmpOrganization.Name,
		}

		return nil
	}); err != nil {
		return nil, err //nolint:wrapcheck
	}

	tokenSet, err := u.authTokenManager.CreateTokenSet(ctx, targetAppUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("create token set: %w", err)
	}

	return tokenSet, nil
}
