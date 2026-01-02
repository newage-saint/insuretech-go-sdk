package models

import (
	"time"
)

// QuoteConvertedEvent represents a quote_converted_event
type QuoteConvertedEvent struct {
	QuoteId       string    `json:"quote_id,omitempty"`
	PolicyId      string    `json:"policy_id,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
}
