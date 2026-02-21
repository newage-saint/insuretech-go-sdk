package models

import (
	"time"
)

// UnderwritingDecision represents a underwriting_decision
type UnderwritingDecision struct {
	UnderwriterComments string                 `json:"underwriter_comments,omitempty"`
	ValidUntil          time.Time              `json:"valid_until,omitempty"`
	QuoteId             string                 `json:"quote_id"`
	Method              *DecisionMethod        `json:"method"`
	Decision            *DecisionType          `json:"decision"`
	RiskLevel           *UnderwritingRiskLevel `json:"risk_level,omitempty"`
	PremiumAdjusted     bool                   `json:"premium_adjusted,omitempty"`
	DecidedAt           time.Time              `json:"decided_at"`
	Id                  string                 `json:"id"`
	RiskScore           string                 `json:"risk_score,omitempty"`
	RiskFactors         string                 `json:"risk_factors,omitempty"`
	Reason              string                 `json:"reason,omitempty"`
	AdjustedPremium     *Money                 `json:"adjusted_premium,omitempty"`
	UnderwriterId       string                 `json:"underwriter_id,omitempty"`
	AuditInfo           interface{}            `json:"audit_info"`
	Conditions          string                 `json:"conditions,omitempty"`
	AdjustmentReason    string                 `json:"adjustment_reason,omitempty"`
}
