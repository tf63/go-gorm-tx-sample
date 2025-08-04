package infrastracture

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository struct {
	_db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{
		_db: db,
	}
}

// Repository内で呼ぶ用
func (b *BaseRepository) DB(ctx context.Context) *gorm.DB {
	// ctxにトランザクションが含まれている場合はそれを使用し、そうでなければ通常のDBを返す
	if tx, ok := GetTx(ctx); ok {
		return tx
	}
	return b._db.WithContext(ctx)
}
