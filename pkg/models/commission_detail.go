package models

import (
	"time"
)

// CommissionDetail represents a commission_detail
type CommissionDetail struct {
	PolicyId         string    `json:"policy_id,omitempty"`
	CommissionAmount *Money    `json:"commission_amount,omitempty"`
	CommissionType   string    `json:"commission_type,omitempty"`
	EarnedAt         time.Time `json:"earned_at,omitempty"`
}
