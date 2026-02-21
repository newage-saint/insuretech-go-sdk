package models

import (
	"time"
)

// UnderwritingDecision represents a underwriting_decision
type UnderwritingDecision struct {
	RiskLevel           *UnderwritingRiskLevel `json:"risk_level,omitempty"`
	Id                  string                 `json:"id"`
	Reason              string                 `json:"reason,omitempty"`
	Conditions          string                 `json:"conditions,omitempty"`
	PremiumAdjusted     bool                   `json:"premium_adjusted,omitempty"`
	UnderwriterId       string                 `json:"underwriter_id,omitempty"`
	UnderwriterComments string                 `json:"underwriter_comments,omitempty"`
	AuditInfo           interface{}            `json:"audit_info"`
	Method              *DecisionMethod        `json:"method"`
	DecidedAt           time.Time              `json:"decided_at"`
	ValidUntil          time.Time              `json:"valid_until,omitempty"`
	QuoteId             string                 `json:"quote_id"`
	Decision            *DecisionType          `json:"decision"`
	RiskScore           string                 `json:"risk_score,omitempty"`
	RiskFactors         string                 `json:"risk_factors,omitempty"`
	AdjustedPremium     *Money                 `json:"adjusted_premium,omitempty"`
	AdjustmentReason    string                 `json:"adjustment_reason,omitempty"`
}
