package models

import (
	"time"
)

// AiFraudDetectedEvent represents a ai_fraud_detected_event
type AiFraudDetectedEvent struct {
	Severity             string    `json:"severity,omitempty"`
	CorrelationId        string    `json:"correlation_id,omitempty"`
	EventId              string    `json:"event_id,omitempty"`
	FraudScore           float64   `json:"fraud_score,omitempty"`
	RequiresManualReview bool      `json:"requires_manual_review,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
	AgentId              string    `json:"agent_id,omitempty"`
	EntityType           string    `json:"entity_type,omitempty"`
	EntityId             string    `json:"entity_id,omitempty"`
	RiskFactors          []string  `json:"risk_factors,omitempty"`
}
