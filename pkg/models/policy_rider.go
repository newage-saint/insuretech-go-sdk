package models

import (
	"time"
)

// PolicyRider represents a policy_rider
type PolicyRider struct {
	RiderId          string    `json:"rider_id,omitempty"`
	PolicyId         string    `json:"policy_id,omitempty"`
	PremiumAmount    *Money    `json:"premium_amount,omitempty"`
	CoverageAmount   *Money    `json:"coverage_amount,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	PremiumCurrency  string    `json:"premium_currency,omitempty"`
	CoverageCurrency string    `json:"coverage_currency,omitempty"`
	RiderName        string    `json:"rider_name,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
