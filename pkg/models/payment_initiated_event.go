package models

import (
	"time"
)

// PaymentInitiatedEvent represents a payment_initiated_event
type PaymentInitiatedEvent struct {
	CorrelationId string    `json:"correlation_id,omitempty"`
	PaymentId     string    `json:"payment_id,omitempty"`
	Amount        *Money    `json:"amount,omitempty"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	ReferenceType string    `json:"reference_type,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	PayerId       string    `json:"payer_id,omitempty"`
	ReferenceId   string    `json:"reference_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}
