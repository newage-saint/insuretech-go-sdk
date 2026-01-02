package models

import (
	"time"
)

// PolicyRider represents a policy_rider
type PolicyRider struct {
	CoverageAmount *Money    `json:"coverage_amount"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	RiderId        string    `json:"rider_id"`
	PolicyId       string    `json:"policy_id"`
	RiderName      string    `json:"rider_name"`
	PremiumAmount  *Money    `json:"premium_amount"`
}
