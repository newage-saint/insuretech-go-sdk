package models

import (
	"time"
)

// PaymentFailedEvent represents a payment_failed_event
type PaymentFailedEvent struct {
	ErrorCode     string    `json:"error_code,omitempty"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	PaymentId     string    `json:"payment_id,omitempty"`
	Amount        *Money    `json:"amount,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	PayerId       string    `json:"payer_id,omitempty"`
	PaymentMethod string    `json:"payment_method,omitempty"`
}
