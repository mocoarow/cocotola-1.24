package service

import "context"

type StateRepository interface {
	GenerateState(ctx context.Context) (string, error)

	DoesStateExists(ctx context.Context, state string) (bool, error)
}
