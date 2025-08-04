package infrastracture

import (
	"context"

	"github.com/tf63/go-gorm-tx-sample/internal/uow-pattern/db"
	"github.com/tf63/go-gorm-tx-sample/internal/uow-pattern/domain"
)

type accountRepositoryImpl struct {
	db db.DB
}

func NewAccountRepository(db db.DB) domain.AccountRepository {
	return &accountRepositoryImpl{
		db: db,
	}
}

func (r *accountRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Account, error) {
	// (IDで口座を検索する)
	return &domain.Account{}, nil
}

func (r *accountRepositoryImpl) Save(ctx context.Context, account domain.Account) error {
	// (永続化する)
	return nil
}
