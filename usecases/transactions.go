package usecases

import (
	"context"
	"log"
	"mobee-test/libs/helper"
	"mobee-test/models"
	"mobee-test/repositories"
	"time"

	"github.com/jmoiron/sqlx"
)

type transactionUseCase struct {
	contextTimeout        time.Duration
	transactionRepository repositories.TransactionRepository
	db                    *sqlx.DB
}

type TransactionUseCase interface {
	Deposit(ctx context.Context, request models.DepositRequest) error
}

func NewTransactionUseCase(contextTimeout time.Duration, transactionRepository repositories.TransactionRepository, db *sqlx.DB) TransactionUseCase {
	return &transactionUseCase{
		contextTimeout,
		transactionRepository,
		db,
	}
}

func (tr *transactionUseCase) Deposit(ctx context.Context, request models.DepositRequest) error {
	ctx, cancel := context.WithTimeout(ctx, tr.contextTimeout)
	defer cancel()

	tx, err := tr.db.Begin()
	if err != nil {
		log.Println("SQL error in UseCase Deposit => Open Transaction", err)
		return err
	}

	err = tr.transactionRepository.Deposit(ctx, tx, request)
	if err != nil {
		log.Println("Error deposit: ", err)
		return err
	}

	helper.CommitOrRollback(tx, err)

	return err
}
