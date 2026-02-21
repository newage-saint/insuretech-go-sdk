package models

import (
	"time"
)

// UnderwritingDecisionEvent represents a underwriting_decision_event
type UnderwritingDecisionEvent struct {
	CorrelationId      string    `json:"correlation_id,omitempty"`
	PolicyId           string    `json:"policy_id,omitempty"`
	AgentId            string    `json:"agent_id,omitempty"`
	RiskScore          float64   `json:"risk_score,omitempty"`
	RecommendedPremium *Money    `json:"recommended_premium,omitempty"`
	RiskFactors        []string  `json:"risk_factors,omitempty"`
	EventId            string    `json:"event_id,omitempty"`
	Decision           string    `json:"decision,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}
