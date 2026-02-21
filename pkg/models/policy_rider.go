package models

import (
	"time"
)

// PolicyRider represents a policy_rider
type PolicyRider struct {
	CoverageAmount   *Money    `json:"coverage_amount,omitempty"`
	PremiumCurrency  string    `json:"premium_currency,omitempty"`
	CoverageCurrency string    `json:"coverage_currency,omitempty"`
	PolicyId         string    `json:"policy_id,omitempty"`
	PremiumAmount    *Money    `json:"premium_amount,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	RiderId          string    `json:"rider_id,omitempty"`
	RiderName        string    `json:"rider_name,omitempty"`
}
