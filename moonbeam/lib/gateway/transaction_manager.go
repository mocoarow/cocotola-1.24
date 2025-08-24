package gateway

import (
	"context"

	"gorm.io/gorm"

	"github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
)

type transactionManagerT[RF any] struct {
	db  *gorm.DB
	rff func(ctx context.Context, db *gorm.DB) (RF, error)
}

func NewTransactionManagerT[RF any](db *gorm.DB, rff func(ctx context.Context, db *gorm.DB) (RF, error)) (service.TransactionManagerT[RF], error) {
	return &transactionManagerT[RF]{
		db:  db,
		rff: rff,
	}, nil
}

func (t *transactionManagerT[RF]) Do(ctx context.Context, fn func(rf RF) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error { // nolint:wrapcheck
		rf, err := t.rff(ctx, tx)
		if err != nil {
			return err // nolint:wrapcheck
		}
		return fn(rf)
	})
}

type nonTransactionManagerT[RF any] struct {
	rf RF
}

func NewNonTransactionManagerT[RF any](rf RF) (service.TransactionManagerT[RF], error) {
	return &nonTransactionManagerT[RF]{
		rf: rf,
	}, nil
}

func (t *nonTransactionManagerT[RF]) Do(ctx context.Context, fn func(rf RF) error) error {
	return fn(t.rf)
}
