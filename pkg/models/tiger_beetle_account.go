package models

import (
	"time"
)

// TigerBeetleAccount represents a tiger_beetle_account
type TigerBeetleAccount struct {
	LedgerId             int       `json:"ledger_id"`
	Balance              string    `json:"balance"`
	BalanceUpdatedAt     time.Time `json:"balance_updated_at,omitempty"`
	IsActive             bool      `json:"is_active"`
	UpdatedAt            time.Time `json:"updated_at"`
	AccountId            string    `json:"account_id"`
	TigerbeetleAccountId string    `json:"tigerbeetle_account_id"`
	AccountType          string    `json:"account_type"`
	Currency             string    `json:"currency"`
	CreatedAt            time.Time `json:"created_at"`
	UserId               string    `json:"user_id,omitempty"`
}
