package models

import (
	"time"
)

// UnderwritingDecision represents a underwriting_decision
type UnderwritingDecision struct {
	UnderwriterComments string                 `json:"underwriter_comments,omitempty"`
	AuditInfo           *AuditInfo             `json:"audit_info,omitempty"`
	RiskScore           string                 `json:"risk_score,omitempty"`
	Reason              string                 `json:"reason,omitempty"`
	UnderwriterId       string                 `json:"underwriter_id,omitempty"`
	Id                  string                 `json:"id"`
	QuoteId             string                 `json:"quote_id"`
	Method              *DecisionMethod        `json:"method"`
	RiskFactors         string                 `json:"risk_factors,omitempty"`
	PremiumAdjusted     bool                   `json:"premium_adjusted,omitempty"`
	AdjustmentReason    string                 `json:"adjustment_reason,omitempty"`
	ValidUntil          time.Time              `json:"valid_until,omitempty"`
	Decision            *DecisionType          `json:"decision"`
	RiskLevel           *UnderwritingRiskLevel `json:"risk_level,omitempty"`
	DecidedAt           time.Time              `json:"decided_at"`
	Conditions          string                 `json:"conditions,omitempty"`
	AdjustedPremium     *Money                 `json:"adjusted_premium,omitempty"`
}
