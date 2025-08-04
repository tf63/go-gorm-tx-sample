package infrastracture

import (
	"context"

	"github.com/tf63/go-gorm-tx-sample/internal/anti-pattern/domain"
	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return &accountRepositoryImpl{DB: db}
}

func (r *accountRepositoryImpl) WithTx(tx *gorm.DB) domain.AccountRepository {
	return &accountRepositoryImpl{DB: tx}
}

func (r *accountRepositoryImpl) Save(ctx context.Context, account domain.Account) error {
	// (口座情報を保存する)
	if err := r.DB.WithContext(ctx).Save(&account).Error; err != nil {
		return err
	}

	return nil
}

func (r *accountRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Account, error) {
	// (IDで口座を検索する)
	var account domain.Account
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

// アンチパターン: Repositoryにトランザクションを含める
func (a *accountRepositoryImpl) Transfer(
	ctx context.Context,
	fromID, toID, amount int,
) error {
	// トランザクションを開始
	return a.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// (fromIDの口座を取得)
		fromAccount, err := a.WithTx(tx).FindByID(ctx, fromID)
		if err != nil {
			return err
		}

		// (toIDの口座を取得)
		toAccount, err := a.WithTx(tx).FindByID(ctx, toID)
		if err != nil {
			return err
		}

		// (引き出し処理)
		if err := fromAccount.Withdraw(amount); err != nil {
			return err
		}

		// (入金処理)
		toAccount.Deposit(amount)

		// (口座情報を保存)
		if err := a.WithTx(tx).Save(ctx, *fromAccount); err != nil {
			return err
		}

		if err := a.WithTx(tx).Save(ctx, *toAccount); err != nil {
			return err
		}

		return nil // 成功時はnilを返す
	})
}
