package models

import (
	"time"
)

// UnderwritingRejectedEvent represents a underwriting_rejected_event
type UnderwritingRejectedEvent struct {
	RiskLevel     string    `json:"risk_level,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	QuoteId       string    `json:"quote_id,omitempty"`
	DecisionId    string    `json:"decision_id,omitempty"`
	Reason        string    `json:"reason,omitempty"`
}
