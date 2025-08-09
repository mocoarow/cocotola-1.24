package usecase

import (
	// "context"
	// "fmt"

	"context"

	"github.com/golang-jwt/jwt/v5"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type AppUserClaims struct {
	LoginID          string `json:"loginId"`
	AppUserID        int    `json:"appUserId"`
	Username         string `json:"username"`
	OrganizationID   int    `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	TokenType        string `json:"tokenType"`
	jwt.RegisteredClaims
}

type Authentication struct {
	transactionManager            service.TransactionManager
	authTokenManager              service.AuthTokenManager
	systemOwnerByOrganizationName SystemOwnerByOrganizationName
}

func NewAuthentication(transactionManager service.TransactionManager, authTokenManager service.AuthTokenManager, systemOwnerByOrganizationName SystemOwnerByOrganizationName) *Authentication {
	return &Authentication{
		transactionManager:            transactionManager,
		authTokenManager:              authTokenManager,
		systemOwnerByOrganizationName: systemOwnerByOrganizationName,
	}
}

func (u *Authentication) SignInWithIDToken(ctx context.Context, idToken string) (*domain.AuthTokenSet, error) {
	tokenSet, err := u.authTokenManager.SignInWithIDToken(ctx, idToken)
	if err != nil {
		return nil, mbliberrors.Errorf("SignInWithIDToken. err: %w", err)
	}
	return tokenSet, nil
}

func (u *Authentication) GetUserInfo(ctx context.Context, bearerToken string) (*mbuserdomain.AppUserModel, error) {
	// TODO: Check whether the token is registered in the Database

	appUserInfo, err := u.authTokenManager.GetUserInfo(ctx, bearerToken)
	if err != nil {
		return nil, err
	}

	var targetAppUserModel *mbuserdomain.AppUserModel

	if err := u.transactionManager.Do(ctx, func(rf service.RepositoryFactory) error {
		systemOwner, err := u.systemOwnerByOrganizationName.Get(ctx, rf, appUserInfo.OrganizationName)
		if err != nil {
			return err
		}

		appUser, err := systemOwner.FindAppUserByLoginID(ctx, appUserInfo.LoginID)
		if err != nil {
			return err
		}

		targetAppUserModel = appUser.AppUserModel
		return nil
	}); err != nil {
		return nil, mbliberrors.Errorf("RegisterAppUser. err: %w", err)
	}

	return targetAppUserModel, nil
}

func (u *Authentication) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	accessToken, err := u.authTokenManager.RefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	// TODO: Save the token to the database

	return accessToken, nil
}

// func (u *Authentication) Authenticate(ctx context.Context, bearerToken string) (*mbuserdomain.AppUserModel, error) {
// 	logger := mbliblog.GetLoggerFromContext(ctx, liblog.AppUsecaseLoggerContextKey)

// 	token, err := jwt.ParseWithClaims(bearerToken, &AppUserClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return u.signingKey, nil
// 	})
// 	if err != nil {
// 		logger.InfoContext(ctx, fmt.Sprintf("invalid token. err: %v", err))
// 		return nil, domain.ErrUnauthenticated
// 	}

// 	claims, ok := token.Claims.(*AppUserClaims)
// 	if !ok || !token.Valid {
// 		// logger.InfoContext(ctx, "invalid token")
// 		return nil, domain.ErrUnauthenticated
// 	}

// 	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, u.rf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	organizationID, err := mbuserdomain.NewOrganizationID(claims.OrganizationID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	appUserID, err := mbuserdomain.NewAppUserID(claims.AppUserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// # TODO Check whether the token is registered in the Database

// 	appUserRepo := u.rf.NewAppUserRepository(ctx)
// 	systemOwner, err := appUserRepo.FindSystemOwnerByOrganizationID(ctx, systemAdmin, organizationID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	appUser, err := systemOwner.FindAppUserByID(ctx, appUserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return appUser.AppUserModel, nil
// }
