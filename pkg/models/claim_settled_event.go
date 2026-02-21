package models

import (
	"time"
)

// ClaimSettledEvent represents a claim_settled_event
type ClaimSettledEvent struct {
	EventId          string    `json:"event_id,omitempty"`
	ClaimNumber      string    `json:"claim_number,omitempty"`
	PaymentReference string    `json:"payment_reference,omitempty"`
	ClaimId          string    `json:"claim_id,omitempty"`
	CustomerId       string    `json:"customer_id,omitempty"`
	SettledAmount    *Money    `json:"settled_amount,omitempty"`
	PaymentMethod    string    `json:"payment_method,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	CorrelationId    string    `json:"correlation_id,omitempty"`
}
