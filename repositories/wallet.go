package repositories

import (
	"context"
	"database/sql"
	"log"
	"mobee-test/models"

	"github.com/jmoiron/sqlx"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, tx *sql.Tx, request models.Wallet) error
}

type walletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) WalletRepository {
	// Initialize and return the repository instance
	return &walletRepository{
		db,
	}
}

func (wr *walletRepository) CreateWallet(ctx context.Context, tx *sql.Tx, request models.Wallet) error {
	var (
		sqlError error
	)

	sql := `
		INSERT INTO wallets (
			userid,
			balance
		) VALUES ($1, $2)
	`

	_, sqlError = tx.ExecContext(ctx, sql, request.UserID, request.Balance)
	if sqlError != nil {
		log.Println("SQL error in Repo CreateWallet => Execute Query", sqlError)
		return sqlError
	}

	return sqlError
}
