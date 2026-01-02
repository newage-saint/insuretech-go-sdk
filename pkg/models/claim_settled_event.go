package models

import (
	"time"
)

// ClaimSettledEvent represents a claim_settled_event
type ClaimSettledEvent struct {
	CustomerId       string    `json:"customer_id,omitempty"`
	SettledAmount    *Money    `json:"settled_amount,omitempty"`
	CorrelationId    string    `json:"correlation_id,omitempty"`
	ClaimId          string    `json:"claim_id,omitempty"`
	PaymentMethod    string    `json:"payment_method,omitempty"`
	PaymentReference string    `json:"payment_reference,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	EventId          string    `json:"event_id,omitempty"`
	ClaimNumber      string    `json:"claim_number,omitempty"`
}
