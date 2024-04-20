package db

import "github.com/jmoiron/sqlx"

// DatabaseRepo interface
type DatabaseRepo interface {
	Connect(host string, port int, user string, password string, dbName string, sslMode string) (*sqlx.DB, error)
}
