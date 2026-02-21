package models

import (
	"time"
)

// TigerBeetleAccount represents a tiger_beetle_account
type TigerBeetleAccount struct {
	AccountId            string    `json:"account_id"`
	TigerbeetleAccountId string    `json:"tigerbeetle_account_id"`
	UserId               string    `json:"user_id,omitempty"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	AccountType          string    `json:"account_type"`
	LedgerId             int       `json:"ledger_id"`
	Currency             string    `json:"currency"`
	Balance              string    `json:"balance"`
	BalanceUpdatedAt     time.Time `json:"balance_updated_at,omitempty"`
}
