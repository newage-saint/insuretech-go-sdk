package models

import (
	"time"
)

// PaymentRefundProcessedEvent represents a payment_refund_processed_event
type PaymentRefundProcessedEvent struct {
	RecipientId       string    `json:"recipient_id,omitempty"`
	Amount            *Money    `json:"amount,omitempty"`
	Reason            string    `json:"reason,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	CorrelationId     string    `json:"correlation_id,omitempty"`
	EventId           string    `json:"event_id,omitempty"`
	RefundId          string    `json:"refund_id,omitempty"`
	OriginalPaymentId string    `json:"original_payment_id,omitempty"`
}
