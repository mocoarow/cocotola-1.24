package gateway

import (
	"context"

	firebase "firebase.google.com/go/v4"
	firebaseauth "firebase.google.com/go/v4/auth"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type FirebaseClient struct {
	firebaseAuthClient *firebaseauth.Client
}

func NewFirebaseClient(ctx context.Context, googleProjectID string) (service.FirebaseClient, error) {
	fireBaseApp, err := firebase.NewApp(ctx, &firebase.Config{ //nolint:exhaustruct
		ProjectID: googleProjectID,
	})
	if err != nil {
		return nil, mbliberrors.Errorf("NewApp: %w", err)
	}

	firebaseAuthClient, err := fireBaseApp.Auth(ctx)
	if err != nil {
		return nil, mbliberrors.Errorf("Auth: %w", err)
	}

	return &FirebaseClient{
		firebaseAuthClient: firebaseAuthClient,
	}, nil
}

func (c *FirebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*service.Token, error) {
	token, err := c.firebaseAuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, mbliberrors.Errorf("VerifyIDToken: %w", err)
	}

	return &service.Token{
		UID:            token.UID,
		SignInProvider: token.Firebase.SignInProvider,
	}, nil
}

func (c *FirebaseClient) GetUser(ctx context.Context, uid string) (*service.UserRecord, error) {
	userRecord, err := c.firebaseAuthClient.GetUser(ctx, uid)
	if err != nil {
		return nil, mbliberrors.Errorf("GetUser: %w", err)
	}

	return &service.UserRecord{
		UID:         userRecord.UID,
		Email:       userRecord.Email,
		DisplayName: userRecord.DisplayName,
	}, nil
}
