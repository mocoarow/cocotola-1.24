package service

import (
	"context"
)

type Token struct {
	UID            string
	SignInProvider string
}

type UserRecord struct {
	UID         string
	Email       string
	DisplayName string
}

type FirebaseClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*Token, error)
	GetUser(ctx context.Context, uid string) (*UserRecord, error)
}
