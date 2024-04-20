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

type walletUseCase struct {
	contextTimeout   time.Duration
	walletRepository repositories.WalletRepository
	db               *sqlx.DB
}

type WalletUseCase interface {
	CreateWallet(ctx context.Context, request models.Wallet) error
}

func NewWalletUseCase(contextTimeout time.Duration, walletRepository repositories.WalletRepository, db *sqlx.DB) WalletUseCase {
	return &walletUseCase{
		contextTimeout,
		walletRepository,
		db,
	}
}

func (wr *walletUseCase) CreateWallet(ctx context.Context, request models.Wallet) error {
	ctx, cancel := context.WithTimeout(ctx, wr.contextTimeout)
	defer cancel()

	tx, err := wr.db.Begin()
	if err != nil {
		log.Println("SQL error in UseCase CreateWallet => Open Transaction", err)
		return err
	}

	err = wr.walletRepository.CreateWallet(ctx, tx, request)
	if err != nil {
		log.Println("Error creating wallet: ", err)
		return err
	}

	helper.CommitOrRollback(tx, err)

	return err
}
