package models

import (
	"time"
)

// PaymentFailedEvent represents a payment_failed_event
type PaymentFailedEvent struct {
	EventId       string    `json:"event_id,omitempty"`
	Amount        *Money    `json:"amount,omitempty"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	ErrorCode     string    `json:"error_code,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	PaymentId     string    `json:"payment_id,omitempty"`
	PayerId       string    `json:"payer_id,omitempty"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
}
