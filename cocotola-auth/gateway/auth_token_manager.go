package gateway

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

type UserClaims struct {
	LoginID string `json:"loginId"`
	// TODO: Check if UserID is needed in the token
	UserID           int    `json:"userId"`
	Username         string `json:"username"`
	OrganizationID   int    `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	// Role             string `json:"role"`
	TokenType string `json:"tokenType"`
	jwt.RegisteredClaims
}

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
	// userID      *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
	loginID        string
	username       string
}

//	func (m *user) UserID() *mbuserdomain.UserID {
//		return m.userID
//	}
func (m *user) GetOrganizationID() *mbuserdomain.OrganizationID {
	return m.organizationID
}
func (m *user) GetUsername() string {
	return m.username
}
func (m *user) GetLoginID() string {
	return m.loginID
}

type AuthTokenManager struct {
	firebaseAuthClient service.FirebaseClient
	SigningKey         []byte
	SigningMethod      jwt.SigningMethod
	TokenTimeout       time.Duration
	RefreshTimeout     time.Duration
	logger             *slog.Logger
}

func NewAuthTokenManager(_ context.Context, firebaseAuthClient service.FirebaseClient, signingKey []byte, signingMethod jwt.SigningMethod, tokenTimeout, refreshTimeout time.Duration) service.AuthTokenManager {
	return &AuthTokenManager{
		firebaseAuthClient: firebaseAuthClient,
		SigningKey:         signingKey,
		SigningMethod:      signingMethod,
		TokenTimeout:       tokenTimeout,
		RefreshTimeout:     refreshTimeout,
		logger:             slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"AuthTokenManager")),
	}
}

func (m *AuthTokenManager) SignInWithIDToken(ctx context.Context, idToken string) (*domain.AuthTokenSet, error) {
	token, err := m.firebaseAuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, mbliberrors.Errorf("m.firebaseAuthClient.VerifyIDToken. err: %w", err)
	}
	userRecord, err := m.firebaseAuthClient.GetUser(ctx, token.UID)
	if err != nil {
		return nil, mbliberrors.Errorf("m.firebaseAuthClient.GetUser. err: %w", err)
	}
	loginID := userRecord.UID
	username := "Anonymous"
	if token.SignInProvider != "anonymous" {
		loginID = userRecord.Email
		username = userRecord.DisplayName
	}

	organizationID, err := mbuserdomain.NewOrganizationID(1)
	if err != nil {
		return nil, mbliberrors.Errorf("mbuserdomain.NewOrganizationID. err: %w", err)
	}

	user := user{
		// UserID:        userRecord.UserID,
		loginID:        loginID,
		username:       username,
		organizationID: organizationID,
	}

	organization := organization{
		organizationID: organizationID,
		name:           "cocotola",
	}

	tokenSet, err := m.CreateTokenSet(ctx, &user, &organization)
	if err != nil {
		return nil, mbliberrors.Errorf("m.CreateTokenSet. err: %w", err)
	}

	return tokenSet, nil
}

func (m *AuthTokenManager) CreateTokenSet(ctx context.Context, user service.UserInterface, organization service.OrganizationInterface) (*domain.AuthTokenSet, error) {
	if user == nil {
		return nil, mbliberrors.Errorf("user is nil. err: %w", mblibdomain.ErrInvalidArgument)
	}
	accessToken, err := m.createJWT(ctx, user, organization, m.TokenTimeout, "access")
	if err != nil {
		return nil, err
	}

	refreshToken, err := m.createJWT(ctx, user, organization, m.RefreshTimeout, "refresh")
	if err != nil {
		return nil, err
	}

	return &domain.AuthTokenSet{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *AuthTokenManager) createJWT(ctx context.Context, user service.UserInterface, organization service.OrganizationInterface, duration time.Duration, tokenType string) (string, error) {
	if len(m.SigningKey) == 0 {
		return "", mbliberrors.Errorf("m.SigningKey is not set")
	}

	now := time.Now()
	claims := UserClaims{ //nolint:exhaustruct
		// UserID:        user.UserID().Int(),
		LoginID:          user.GetLoginID(),
		Username:         user.GetUsername(),
		OrganizationID:   organization.GetOrganizationID().Int(),
		OrganizationName: organization.GetName(),
		TokenType:        tokenType,
		RegisteredClaims: jwt.RegisteredClaims{ //nolint:exhaustruct
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}

	m.logger.DebugContext(ctx, fmt.Sprintf("claims: %+v", claims))

	token := jwt.NewWithClaims(m.SigningMethod, claims)
	signed, err := token.SignedString(m.SigningKey)
	if err != nil {
		return "", mbliberrors.Errorf(". err: %w", err)
	}

	return signed, nil
}

func (m *AuthTokenManager) GetUserInfo(ctx context.Context, tokenString string) (*service.UserInfo, error) {
	currentClaims, err := m.parseToken(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("parseToken(%s). err: %w", err.Error(), domain.ErrUnauthenticated)
	}

	return &service.UserInfo{
		// UserID:        currentClaims.UserID,
		LoginID:          currentClaims.LoginID,
		Username:         currentClaims.Username,
		OrganizationID:   currentClaims.OrganizationID,
		OrganizationName: currentClaims.OrganizationName,
	}, nil
}

func (m *AuthTokenManager) parseToken(ctx context.Context, tokenString string) (*UserClaims, error) {
	keyFunc := func(_ *jwt.Token) (interface{}, error) {
		return m.SigningKey, nil
	}

	currentToken, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, keyFunc) //nolint:exhaustruct
	if err != nil {
		m.logger.InfoContext(ctx, fmt.Sprintf("%v", err))
		// return nil, fmt.Errorf("jwt.ParseWithClaims. err: %w", domain.ErrUnauthenticated)
		return nil, mbliberrors.Errorf("ParseWithClaims: %w", err)
	}
	if !currentToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	currentClaims, ok := currentToken.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	v := jwt.NewValidator()
	if err := v.Validate(currentClaims); err != nil {
		return nil, mbliberrors.Errorf("Validate: %w", err)
	}

	return currentClaims, nil
}

func (m *AuthTokenManager) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	currentClaims, err := m.parseToken(ctx, tokenString)
	if err != nil {
		return "", fmt.Errorf("parseToken(%s). err: %w", err.Error(), domain.ErrUnauthenticated)
	}

	if currentClaims.TokenType != "refresh" {
		return "", fmt.Errorf("invalid token type. err: %w", domain.ErrUnauthenticated)
	}

	// userID, err := mbuserdomain.NewUserID(currentClaims.UserID)
	// if err != nil {
	// 	return "", err
	// }

	// TODO: check
	user := &user{ //nolint:exhaustruct
		// userID: userID,
		loginID:  currentClaims.LoginID,
		username: currentClaims.Username,
	}

	organizationID, err := mbuserdomain.NewOrganizationID(currentClaims.OrganizationID)
	if err != nil {
		return "", mbliberrors.Errorf("NewOrganizationID: %w", err)
	}

	organization := &organization{
		organizationID: organizationID,
		name:           currentClaims.OrganizationName,
	}

	accessToken, err := m.createJWT(ctx, user, organization, m.TokenTimeout, "access")
	if err != nil {
		return "", mbliberrors.Errorf("m.createJWT. err: %w", err)
	}

	return accessToken, nil
}
