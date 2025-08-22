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

type GoogleAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GoogleUserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GoogleUserUsecase struct {
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
	googleAuthClient GoogleAuthClient
	logger           *slog.Logger
}

func NewGoogleUser(txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager, googleAuthClient GoogleAuthClient) *GoogleUserUsecase {
	return &GoogleUserUsecase{
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
			return err
		}

		tmpState, err := stateRepo.GenerateState(ctx)
		if err != nil {
			return err
		}

		state = tmpState
		return nil
	}); err != nil {
		return "", err
	}

	return state, nil
}

func (u *GoogleUserUsecase) doesStateExist(ctx context.Context, state string) error {
	var matched bool
	if err := u.nonTxManager.Do(ctx, func(rf service.RepositoryFactory) error {
		stateRepo, err := rf.NewStateRepository(ctx)
		if err != nil {
			return err
		}
		tmpMatched, err := stateRepo.DoesStateExists(ctx, state)
		if err != nil {
			return err
		}

		matched = tmpMatched
		return nil
	}); err != nil {
		return err
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
		return nil, mbliberrors.Errorf(". err: %w", err)
	}

	var tokenSet *domain.AuthTokenSet

	var targetOorganization *organization
	var targetAppUser *appUser
	if err := u.txManager.Do(ctx, func(rf service.RepositoryFactory) error {
		tmpOrganization, tmpAppUser, err := u.registerAppUser(ctx, rf, organizationName, info.Email, info.Name, info.Email, accessToken, refreshToken)
		if err != nil && !errors.Is(err, mbuserservice.ErrAppUserAlreadyExists) {
			return mbliberrors.Errorf("s.registerAppUser. err: %w", err)
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

func (u *GoogleUserUsecase) registerAppUser(ctx context.Context, rf service.RepositoryFactory, organizationName string, loginID string, username string, providerID, providerAccessToken, providerRefreshToken string) (*mbuserdomain.OrganizationModel, *mbuserdomain.AppUserModel, error) {
	action, err := NewOrganizationAction(ctx, rf,
		WithOrganizationRepository(),
		WithOrganization(organizationName),
		WithAppUserRepository(),
	)
	if err != nil {
		return nil, nil, err
	}

	// findOrganization := func() (*mbuserdomain.OrganizationModel, error) {
	// 	organization, err := action.systemAdmin.FindOrganizationByName(ctx, organizationName)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return organization.OrganizationModel, nil
	// }

	findAppUser := func() (*mbuserdomain.AppUserModel, error) {
		appUser1, err := action.systemOwner.FindAppUserByLoginID(ctx, loginID)
		if err == nil {
			return appUser1.AppUserModel, nil
		}

		if !errors.Is(err, mbuserservice.ErrAppUserNotFound) {
			u.logger.InfoContext(ctx, fmt.Sprintf("Unsupported %v", err))
			return nil, mbliberrors.Errorf("systemOwner.FindAppUserByLoginID. err: %w", err)
		}

		u.logger.InfoContext(ctx, fmt.Sprintf("Add student. %+v", appUser1))
		parameter, err := mbuserservice.NewAppUserAddParameter(
			loginID,  //googleUserInfo.Email,
			username, //googleUserInfo.Name,
			"",
			"google",
			providerID,           // googleUserInfo.Email,
			providerAccessToken,  // googleAuthResponse.AccessToken,
			providerRefreshToken, // googleAuthResponse.RefreshToken,
		)
		if err != nil {
			return nil, mbliberrors.Errorf("invalid AppUserAddParameter. err: %w", err)
		}

		studentID, err := action.systemOwner.AddAppUser(ctx, parameter)
		if err != nil {
			return nil, mbliberrors.Errorf("failed to AddStudent. err: %w", err)
		}

		appUser2, err := action.systemOwner.FindAppUserByID(ctx, studentID)
		if err != nil {
			return nil, mbliberrors.Errorf("failed to FindStudentByID. err: %w", err)
		}

		return appUser2.AppUserModel, nil
	}

	// organization, err := findOrganization()
	// if err != nil {
	// 	return nil, nil, err
	// }

	appUser, err := findAppUser()
	if errors.Is(err, mbuserservice.ErrAppUserAlreadyExists) {
		return action.organization.OrganizationModel, appUser, nil
	} else if err != nil {
		return nil, nil, mbliberrors.Errorf("registerAppUser. err: %w", err)
	}

	return action.organization.OrganizationModel, appUser, nil
}
