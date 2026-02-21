package models

import (
	"time"
)

// UnderwritingRejectedEvent represents a underwriting_rejected_event
type UnderwritingRejectedEvent struct {
	EventId       string    `json:"event_id,omitempty"`
	QuoteId       string    `json:"quote_id,omitempty"`
	DecisionId    string    `json:"decision_id,omitempty"`
	Reason        string    `json:"reason,omitempty"`
	RiskLevel     string    `json:"risk_level,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}
