package models

import (
	"time"
)

// MFSTransactionFailedEvent represents a mfs_transaction_failed_event
type MFSTransactionFailedEvent struct {
	PaymentId        string    `json:"payment_id,omitempty"`
	CorrelationId    string    `json:"correlation_id,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	EventId          string    `json:"event_id,omitempty"`
	MfsTransactionId string    `json:"mfs_transaction_id,omitempty"`
	ErrorMessage     string    `json:"error_message,omitempty"`
}
