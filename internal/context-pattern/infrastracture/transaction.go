package infrastracture

import (
	"context"

	"github.com/tf63/go-gorm-tx-sample/internal/context-pattern/domain"
	"gorm.io/gorm"
)

type txManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) domain.TxManager {
	return &txManager{
		db,
	}
}

func (tm *txManager) DoInTx(
	ctx context.Context,
	fn domain.TxFunction,
) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctxWithTx := WithTx(ctx, tx)
		if err := fn(ctxWithTx); err != nil {
			return err
		}
		return nil
	})
}
