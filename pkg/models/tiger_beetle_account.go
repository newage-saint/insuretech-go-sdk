package models

import (
	"time"
)

// TigerBeetleAccount represents a tiger_beetle_account
type TigerBeetleAccount struct {
	AccountId            string    `json:"account_id"`
	UserId               string    `json:"user_id,omitempty"`
	AccountType          string    `json:"account_type"`
	IsActive             bool      `json:"is_active"`
	TigerbeetleAccountId string    `json:"tigerbeetle_account_id"`
	LedgerId             int       `json:"ledger_id"`
	Currency             string    `json:"currency"`
	Balance              string    `json:"balance"`
	BalanceUpdatedAt     time.Time `json:"balance_updated_at,omitempty"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
