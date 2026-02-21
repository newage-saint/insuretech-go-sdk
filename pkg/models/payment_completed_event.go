package models

import (
	"time"
)

// PaymentCompletedEvent represents a payment_completed_event
type PaymentCompletedEvent struct {
	Amount        *Money    `json:"amount,omitempty"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	PaymentId     string    `json:"payment_id,omitempty"`
	TransactionId string    `json:"transaction_id,omitempty"`
}
