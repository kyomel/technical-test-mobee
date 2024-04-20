package models

import "github.com/shopspring/decimal"

type DepositRequest struct {
	UserID          int             `json:"user_id"`
	TransactionType string          `json:"transaction_type"`
	Balance         decimal.Decimal `json:"balance"`
	Amount          decimal.Decimal `json:"amount"`
}

type TransactionsDetail struct {
	TransactionsID int `json:"transactions_id"`
}
