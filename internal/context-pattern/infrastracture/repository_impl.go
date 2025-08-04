package infrastracture

import (
	"context"

	"github.com/tf63/go-gorm-tx-sample/internal/context-pattern/domain"
	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	BaseRepository
}

func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return &accountRepositoryImpl{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *accountRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Account, error) {
	// (IDで口座を検索する)
	var account domain.Account
	if err := r.DB(ctx).Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *accountRepositoryImpl) Save(ctx context.Context, account domain.Account) error {
	// (永続化する)
	if err := r.DB(ctx).Save(&account).Error; err != nil {
		return err
	}

	return nil
}
