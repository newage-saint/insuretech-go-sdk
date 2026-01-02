package models

import (
	"time"
)

// AiFraudDetectedEvent represents a ai_fraud_detected_event
type AiFraudDetectedEvent struct {
	AgentId              string    `json:"agent_id,omitempty"`
	EntityId             string    `json:"entity_id,omitempty"`
	FraudScore           float64   `json:"fraud_score,omitempty"`
	RiskFactors          []string  `json:"risk_factors,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
	CorrelationId        string    `json:"correlation_id,omitempty"`
	EventId              string    `json:"event_id,omitempty"`
	EntityType           string    `json:"entity_type,omitempty"`
	Severity             string    `json:"severity,omitempty"`
	RequiresManualReview bool      `json:"requires_manual_review,omitempty"`
}
