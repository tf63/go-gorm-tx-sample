package domain

import "context"

type AccountRepository interface {
	FindByID(ctx context.Context, id int) (*Account, error)
	Save(ctx context.Context, account Account) error
	// アンチパターン
	Transfer(ctx context.Context, fromID, toID, amount int) error
}
