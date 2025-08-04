package domain

import "context"

type AccountRepository interface {
	FindByID(ctx context.Context, id int) (*Account, error)
	FindByIDWithTx(ctx context.Context, id int, tx Tx) (*Account, error)
	Save(ctx context.Context, account Account) error
	SaveWithTx(ctx context.Context, account Account, tx Tx) error
}
