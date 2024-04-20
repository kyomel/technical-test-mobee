package models

import "github.com/shopspring/decimal"

type Wallet struct {
	UserID  int             `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}
