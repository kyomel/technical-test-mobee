package repositories

import (
	"context"
	"database/sql"
	"log"
	"mobee-test/libs/utils"
	"mobee-test/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Deposit(ctx context.Context, tx *sql.Tx, request models.DepositRequest) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	// Initialize and return the repository instance
	return &transactionRepository{
		db,
	}
}

func (tr *transactionRepository) CheckBalance(ctx context.Context, tx *sql.Tx, userID int) (*models.DepositRequest, error) {
	var res models.DepositRequest

	sql := `
		select 
			balance
		from wallets
		where userid = $1
	`

	err := tx.QueryRowContext(ctx, sql, userID).Scan(
		&res.Balance,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (tr *transactionRepository) Deposit(ctx context.Context, tx *sql.Tx, request models.DepositRequest) error {
	var (
		transactionid int
	)

	wallet, err := tr.CheckBalance(ctx, tx, request.UserID)
	if err != nil {
		log.Println("error get balance wallet", err)
		return err
	}

	sqlDeposit := `
		INSERT INTO transactions(
			transactiontype,
			amount,
			senderid,
			receiverid
		) VALUES ($1, $2, $3, $4)
		RETURNING transactionid
	`

	sqlError := tx.QueryRow(sqlDeposit, request.TransactionType, request.Amount, request.UserID, request.UserID).Scan(
		&transactionid,
	)

	if sqlError != nil {
		log.Println("SQL error in Repo Deposit => Execute Query", sqlError)
		return sqlError
	}

	sqladdWallet := `
		UPDATE wallets
		SET balance = $1
		WHERE userid = $2
	`

	valueUpdate := utils.AddDeposit(wallet.Balance, request.Amount)

	_, sqlError = tx.ExecContext(ctx, sqladdWallet, valueUpdate, request.UserID)
	if sqlError != nil {
		log.Println("cannot update data wallet", err)
		return err
	}

	walletUpdate, err := tr.CheckBalance(ctx, tx, request.UserID)
	if err != nil {
		log.Println("error get update balance wallet", err)
		return err
	}

	sqlHistory := `
		INSERT INTO ledger(
			transactionid,
			senderid, 
			receiverid,
			amount,
			balancebefore,
			balanceafter
		) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, sqlError = tx.ExecContext(ctx, sqlHistory, transactionid, request.UserID, request.UserID, request.Amount, wallet.Balance, walletUpdate.Balance)
	if sqlError != nil {
		log.Println("error insert history", err)
		return err
	}

	return sqlError
}
