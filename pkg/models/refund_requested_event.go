package models

import (
	"time"
)

// RefundRequestedEvent represents a refund_requested_event
type RefundRequestedEvent struct {
	PolicyId      string    `json:"policy_id,omitempty"`
	Reason        string    `json:"reason,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	RefundId      string    `json:"refund_id,omitempty"`
	RefundNumber  string    `json:"refund_number,omitempty"`
}
