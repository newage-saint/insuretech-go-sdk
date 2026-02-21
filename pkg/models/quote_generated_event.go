package models

import (
	"time"
)

// QuoteGeneratedEvent represents a quote_generated_event
type QuoteGeneratedEvent struct {
	ValidUntil    string    `json:"valid_until,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	QuoteId       string    `json:"quote_id,omitempty"`
	QuoteNumber   string    `json:"quote_number,omitempty"`
	TotalPremium  *Money    `json:"total_premium,omitempty"`
}
