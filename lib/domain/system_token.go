package domain

import "sync/atomic"

var (
	systemTokenCounter = atomic.Int64{}
)

type SystemToken interface {
	IsSystemToken() bool
}
type systemToken struct {
}

var _ SystemToken = (*systemToken)(nil)

func NewSystemToken() SystemToken {
	if systemTokenCounter.Load() != 0 {
		panic("system token can be created only once")
	}
	systemTokenCounter.Add(1)

	return &systemToken{}
}

func (t *systemToken) IsSystemToken() bool {
	return true
}
