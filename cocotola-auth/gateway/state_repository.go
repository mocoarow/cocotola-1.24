package gateway

import (
	"context"

	"github.com/google/uuid"
)

type StateRepository struct {
	// cache *lru.Cache[string, bool]
}

func NewStateRepository(ctx context.Context) (*StateRepository, error) {
	// cache, err := lru.New[string, bool](100)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create cache: %w", err)
	// }
	return &StateRepository{
		// cache: cache,
	}, nil
}

func (r *StateRepository) GenerateState(ctx context.Context) (string, error) {
	state, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	// r.cache.Add(state.String(), true)
	return state.String(), nil
}

func (r *StateRepository) DoesStateExists(ctx context.Context, state string) (bool, error) {
	// fmt.Println(state)
	// if _, ok := r.cache.Get(state); !ok {
	// 	return false, nil
	// }

	// TODO: IMPLE

	return true, nil
}
