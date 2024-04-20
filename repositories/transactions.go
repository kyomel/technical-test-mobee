package repositories

import "context"

type TransactionRepository interface {
	Deposit(ctx context.Context)
}
