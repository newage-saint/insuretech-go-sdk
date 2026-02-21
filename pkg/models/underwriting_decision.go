package models

import (
	"time"
)

// UnderwritingDecision represents a underwriting_decision
type UnderwritingDecision struct {
	QuoteId             string                 `json:"quote_id"`
	RiskScore           string                 `json:"risk_score,omitempty"`
	AdjustmentReason    string                 `json:"adjustment_reason,omitempty"`
	UnderwriterId       string                 `json:"underwriter_id,omitempty"`
	Id                  string                 `json:"id"`
	Decision            *DecisionType          `json:"decision"`
	RiskLevel           *UnderwritingRiskLevel `json:"risk_level,omitempty"`
	AdjustedPremium     *Money                 `json:"adjusted_premium,omitempty"`
	DecidedAt           time.Time              `json:"decided_at"`
	AuditInfo           interface{}            `json:"audit_info"`
	PremiumAdjusted     bool                   `json:"premium_adjusted,omitempty"`
	ValidUntil          time.Time              `json:"valid_until,omitempty"`
	UnderwriterComments string                 `json:"underwriter_comments,omitempty"`
	Method              *DecisionMethod        `json:"method"`
	RiskFactors         string                 `json:"risk_factors,omitempty"`
	Reason              string                 `json:"reason,omitempty"`
	Conditions          string                 `json:"conditions,omitempty"`
}
