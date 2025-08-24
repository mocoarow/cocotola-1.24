package usecase

import (
	"context"
	"errors"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type organization struct {
	organizationID *mbuserdomain.OrganizationID
	name           string
}

func (m *organization) OrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *organization) Name() string {
	return m.name
}

type appUser struct {
	appUserID      *mbuserdomain.AppUserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

func (m *appUser) AppUserID() *mbuserdomain.AppUserID {
	return m.appUserID
}
func (m *appUser) OrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *appUser) Username() string {
	return m.username
}
func (m *appUser) LoginID() string {
	return m.loginID
}

type TokenSet struct {
	AccessToken  string
	RefreshToken string
}
type GoogleAuthClient interface {
	RetrieveAccessToken(ctx context.Context, code string) (*domain.AuthTokenSet, error)
	RetrieveUserInfo(ctx context.Context, accessToken string) (*domain.UserInfo, error)
}

// type GoogleAuthResponse struct {
// 	AccessToken  string `json:"access_token"`  // nolint:tagliatelle
// 	RefreshToken string `json:"refresh_token"` // nolint:tagliatelle
// }

type GoogleUserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GoogleUserUsecase struct {
	systemToken      libdomain.SystemToken
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
	googleAuthClient GoogleAuthClient
	logger           *slog.Logger
}

func NewGoogleUser(systemToken libdomain.SystemToken, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager, googleAuthClient GoogleAuthClient) *GoogleUserUsecase {
	return &GoogleUserUsecase{
		systemToken:      systemToken,
		txManager:        txManager,
		nonTxManager:     nonTxManager,
		authTokenManager: authTokenManager,
		googleAuthClient: googleAuthClient,
		logger:           slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"-GoogleUserUsecase")),
	}
}

func (u *GoogleUserUsecase) GenerateState(ctx context.Context) (string, error) {
	var state string
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		stateRepo, err := rf.NewStateRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("NewStateRepository: %w", err)
		}

		tmpState, err := stateRepo.GenerateState(ctx)
		if err != nil {
			return mbliberrors.Errorf("GenerateState: %w", err)
		}

		state = tmpState

		return nil
	}); err != nil {
		return "", err //nolint:wrapcheck
	}

	return state, nil
}

func (u *GoogleUserUsecase) doesStateExist(ctx context.Context, state string) error {
	var matched bool
	if err := u.nonTxManager.Do(ctx, func(rf service.RepositoryFactory) error {
		stateRepo, err := rf.NewStateRepository(ctx)
		if err != nil {
			return mbliberrors.Errorf("DoesStateExists: %w", err)
		}
		tmpMatched, err := stateRepo.DoesStateExists(ctx, state)
		if err != nil {
			return mbliberrors.Errorf("DoesStateExists: %w", err)
		}

		matched = tmpMatched

		return nil
	}); err != nil {
		return err //nolint:wrapcheck
	}

	if !matched {
		return mbliberrors.Errorf("invalid state. err: %w", domain.ErrUnauthenticated)
	}

	return nil
}

func (u *GoogleUserUsecase) getTokensAndUserInfo(ctx context.Context, code string) (string, string, *domain.UserInfo, error) {
	resp, err := u.googleAuthClient.RetrieveAccessToken(ctx, code)
	if err != nil {
		return "", "", nil, mbliberrors.Errorf(". err: %w", err)
	}

	info, err := u.googleAuthClient.RetrieveUserInfo(ctx, resp.AccessToken)
	if err != nil {
		return "", "", nil, mbliberrors.Errorf(". err: %w", err)
	}

	return resp.AccessToken, resp.RefreshToken, info, nil
}

func (u *GoogleUserUsecase) Authorize(ctx context.Context, state, code, organizationName string) (*domain.AuthTokenSet, error) {
	if err := u.doesStateExist(ctx, state); err != nil {
		return nil, err
	}

	accessToken, refreshToken, info, err := u.getTokensAndUserInfo(ctx, code)
	if err != nil {
		return nil, mbliberrors.Errorf("get tokens and user info err: %w", err)
	}

	createAppUserParameterFunc := func() (*mbuserservice.AppUserAddParameter, error) {
		return mbuserservice.NewAppUserAddParameter(
			info.Email, //googleUserInfo.Email,
			info.Name,  //googleUserInfo.Name,
			"",
			"google",
			info.Email,   // googleUserInfo.Email,
			accessToken,  // googleAuthResponse.AccessToken,
			refreshToken, // googleAuthResponse.RefreshToken,
		)
	}

	var tokenSet *domain.AuthTokenSet
	var targetOorganization *organization
	var targetAppUser *appUser
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		action, err := service.NewSystemOwnerAction(ctx, u.systemToken, rf,
			// service.WithOrganizationRepository(),
			service.WithOrganizationByName(organizationName),
			// service.WithAppUserRepository(),
		)
		if err != nil {
			return mbliberrors.Errorf("NewSystemOwnerAction: %w", err)
		}
		organizationID := action.Organization.OrganizationID()

		tmpOrganization, tmpAppUser, err := findOrRegisterAppUser(ctx, u.systemToken, rf, organizationID, info.Email, createAppUserParameterFunc)
		if err != nil && !errors.Is(err, mbuserservice.ErrAppUserAlreadyExists) {
			return mbliberrors.Errorf("s.findOrRegisterAppUser. err: %w", err)
		}

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
		return nil, mbliberrors.Errorf("RegisterAppUser. err: %w", err)
	}

	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetAppUser, targetOorganization)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}
	tokenSet = tokenSetTmp

	return tokenSet, nil
}

// func (u *GoogleUserUsecase) RetrieveAccessToken(ctx context.Context, code string) (*domain.AuthTokenSet, error) {
// 	resp, err := u.googleAuthClient.RetrieveAccessToken(ctx, code)
// 	if err != nil {
// 		return nil, mbliberrors.Errorf(". err: %w", err)
// 	}

// 	return resp, nil
// }

// func (u *GoogleUserUsecase) RetrieveUserInfo(ctx context.Context, googleAuthResponse *domain.AuthTokenSet) (*domain.UserInfo, error) {
// 	info, err := u.googleAuthClient.RetrieveUserInfo(ctx, googleAuthResponse)
// 	if err != nil {
// 		return nil, mbliberrors.Errorf(". err: %w", err)
// 	}

// 	return info, nil
// }

// func (u *GoogleUserUsecase) RegisterAppUser(ctx context.Context, googleUserInfo *domain.UserInfo, googleAuthResponse *domain.AuthTokenSet, organizationName string) (*domain.AuthTokenSet, error) {
// 	var tokenSet *domain.AuthTokenSet

// 	var targetOorganization *organization
// 	var targetAppUser *appUser
// 	if err := u.transactionManager.Do(ctx, func(rf service.RepositoryFactory) error {
// 		tmpOrganization, tmpAppUser, err := u.registerAppUser(ctx, rf, organizationName, googleUserInfo.Email, googleUserInfo.Name, googleUserInfo.Email, googleAuthResponse.AccessToken, googleAuthResponse.RefreshToken)
// 		if err != nil && !errors.Is(err, mbuserservice.ErrAppUserAlreadyExists) {
// 			return mbliberrors.Errorf("s.registerAppUser. err: %w", err)
// 		}

// 		targetAppUser = &appUser{
// 			appUserID:      tmpAppUser.AppUserID,
// 			organizationID: tmpAppUser.OrganizationID,
// 			loginID:        tmpAppUser.LoginID,
// 			username:       tmpAppUser.Username,
// 		}
// 		targetOorganization = &organization{
// 			organizationID: tmpOrganization.OrganizationID,
// 			name:           tmpOrganization.Name,
// 		}

// 		return nil
// 	}); err != nil {
// 		return nil, mbliberrors.Errorf("RegisterAppUser. err: %w", err)
// 	}

// 	// if err := s.registerAppUserCallback(ctx, organizationName, appUser); err != nil {
// 	// 	return nil, mbliberrors.Errorf("registerStudentCallback. err: %w", err)
// 	// }
// 	tokenSetTmp, err := u.authTokenManager.CreateTokenSet(ctx, targetAppUser, targetOorganization)
// 	if err != nil {
// 		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
// 	}
// 	tokenSet = tokenSetTmp
// 	return tokenSet, nil
// }
