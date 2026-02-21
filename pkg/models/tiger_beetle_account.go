package models

import (
	"time"
)

// TigerBeetleAccount represents a tiger_beetle_account
type TigerBeetleAccount struct {
	BalanceUpdatedAt     time.Time `json:"balance_updated_at,omitempty"`
	UpdatedAt            time.Time `json:"updated_at"`
	AccountId            string    `json:"account_id"`
	TigerbeetleAccountId string    `json:"tigerbeetle_account_id"`
	LedgerId             int       `json:"ledger_id"`
	Balance              string    `json:"balance"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
	UserId               string    `json:"user_id,omitempty"`
	AccountType          string    `json:"account_type"`
	Currency             string    `json:"currency"`
}
