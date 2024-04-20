package repositories

import (
	"context"
	"database/sql"
	"log"
	"mobee-test/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, request *models.Users) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	// Initialize and return the repository instance
	return &userRepository{
		db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, tx *sql.Tx, request *models.Users) error {
	var (
		sqlError error
	)

	sql := `
		INSERT INTO users (
			username,
			password,
			email
		) VALUES ($1, $2, $3)
	`

	_, sqlError = tx.ExecContext(ctx, sql, request.Username, request.Password, request.Email)
	if sqlError != nil {
		log.Println("SQL error on Repo CreateUser => Execute Query", sqlError)
		return sqlError
	}

	return sqlError
}
