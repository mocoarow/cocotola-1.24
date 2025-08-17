package gateway

import (
	"context"

	"gorm.io/gorm"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type transactionManager struct {
	db  *gorm.DB
	rff RepositoryFactoryFunc
}

func NewTransactionManager(db *gorm.DB, rff RepositoryFactoryFunc) (service.TransactionManager, error) {
	return &transactionManager{
		db:  db,
		rff: rff,
	}, nil
}

func (t *transactionManager) Do(ctx context.Context, fn func(rf service.RepositoryFactory) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error { // nolint:wrapcheck
		rf, err := t.rff(ctx, tx)
		if err != nil {
			return err // nolint:wrapcheck
		}
		return fn(rf)
	})
}

type noneTransactionManager struct {
	rf service.RepositoryFactory
}

func NewNoneTransactionManager(rf service.RepositoryFactory) (service.TransactionManager, error) {
	return &noneTransactionManager{
		rf: rf,
	}, nil
}

func (t *noneTransactionManager) Do(ctx context.Context, fn func(rf service.RepositoryFactory) error) error {
	return fn(t.rf)
}
