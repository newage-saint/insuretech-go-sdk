package models

import (
	"time"
)

// UnderwritingDecision represents a underwriting_decision
type UnderwritingDecision struct {
	PremiumAdjusted     bool                   `json:"premium_adjusted,omitempty"`
	DecidedAt           time.Time              `json:"decided_at"`
	Method              *DecisionMethod        `json:"method"`
	RiskFactors         string                 `json:"risk_factors,omitempty"`
	UnderwriterId       string                 `json:"underwriter_id,omitempty"`
	ValidUntil          time.Time              `json:"valid_until,omitempty"`
	AuditInfo           interface{}            `json:"audit_info"`
	QuoteId             string                 `json:"quote_id"`
	RiskScore           string                 `json:"risk_score,omitempty"`
	Reason              string                 `json:"reason,omitempty"`
	AdjustedPremium     *Money                 `json:"adjusted_premium,omitempty"`
	AdjustmentReason    string                 `json:"adjustment_reason,omitempty"`
	UnderwriterComments string                 `json:"underwriter_comments,omitempty"`
	Id                  string                 `json:"id"`
	Decision            *DecisionType          `json:"decision"`
	RiskLevel           *UnderwritingRiskLevel `json:"risk_level,omitempty"`
	Conditions          string                 `json:"conditions,omitempty"`
}
