package models

import (
	"time"
)

// FraudCheckResult represents a fraud_check_result
type FraudCheckResult struct {
	FraudScore   float64   `json:"fraud_score"`
	RiskFactors  []string  `json:"risk_factors,omitempty"`
	Flagged      bool      `json:"flagged"`
	ReviewedBy   string    `json:"reviewed_by,omitempty"`
	ReviewedAt   time.Time `json:"reviewed_at,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	FraudCheckId string    `json:"fraud_check_id"`
	ClaimId      string    `json:"claim_id"`
}
