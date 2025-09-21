package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
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

func (m *organization) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *organization) GetName() string {
	return m.name
}

type user struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

func (m *user) GetUserID() *mbuserdomain.UserID {
	return m.userID
}
func (m *user) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *user) GetUsername() string {
	return m.username
}
func (m *user) GetLoginID() string {
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
	mbTxManager      mbuserservice.TransactionManager
	mbNonTxManager   mbuserservice.TransactionManager
	txManager        service.TransactionManager
	nonTxManager     service.TransactionManager
	authTokenManager service.AuthTokenManager
	googleAuthClient GoogleAuthClient
	logger           *slog.Logger
}

func NewGoogleUser(systemToken libdomain.SystemToken, mbTxManager, mbNonTxManager mbuserservice.TransactionManager, txManager, nonTxManager service.TransactionManager, authTokenManager service.AuthTokenManager, googleAuthClient GoogleAuthClient) *GoogleUserUsecase {
	return &GoogleUserUsecase{
		systemToken:      systemToken,
		mbTxManager:      mbTxManager,
		mbNonTxManager:   mbNonTxManager,
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
	systemAdmin := service.NewSystemAdmin(u.systemToken)
	sysOwner, err := u.findSystemOwnerByOrganizationName(ctx, systemAdmin, organizationName)
	if err != nil {
		return nil, mbliberrors.Errorf("Do1: %w", err)
	}

	command := NewRegisterUserCommand(ctx, u.mbTxManager, u.mbNonTxManager, u.authTokenManager)
	param, err := mbuserservice.NewUserAddParameter(
		info.Email, //googleUserInfo.Email,
		info.Name,  //googleUserInfo.Name,
		"",
		"google",
		info.Email,   // googleUserInfo.Email,
		accessToken,  // googleAuthResponse.AccessToken,
		refreshToken, // googleAuthResponse.RefreshToken,
	)
	if err != nil {
		return nil, mbliberrors.Errorf("NewUserAddParameter. err: %w", err)
	}
	tokenSet, err := command.Execute(ctx, sysOwner, param)
	if err != nil {
		return nil, mbliberrors.Errorf("s.authTokenManager.CreateTokenSet. err: %w", err)
	}

	return tokenSet, nil
}

func (u *GoogleUserUsecase) findSystemOwnerByOrganizationName(ctx context.Context, operator mbuserdomain.SystemAdminInterface, organizationName string) (*mbuserdomain.SystemOwner, error) {
	return mblibservice.Do1(ctx, u.mbNonTxManager, func(mbrf mbuserservice.RepositoryFactory) (*mbuserdomain.SystemOwner, error) { //nolint:wrapcheck
		return service.FindSystemOwnerByOrganizationName(ctx, mbrf, operator, organizationName)
	})
}
