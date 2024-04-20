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

type userUseCase struct {
	contextTimeout time.Duration
	userRepository repositories.UserRepository
	db             *sqlx.DB
}

type UserUseCase interface {
	CreateUser(ctx context.Context, request models.Users) error
}

func NewUserUseCase(contextTimeout time.Duration, userRepository repositories.UserRepository, db *sqlx.DB) UserUseCase {
	return &userUseCase{
		contextTimeout,
		userRepository,
		db,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, request models.Users) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	tx, err := u.db.Begin()
	if err != nil {
		log.Println("SQL error in UseCase CreateUser => Open Transaction", err)
		return err
	}

	err = u.userRepository.CreateUser(ctx, tx, &request)
	if err != nil {
		log.Println("Error creating user: ", err)
		return err
	}

	helper.CommitOrRollback(tx, err)

	return err
}
