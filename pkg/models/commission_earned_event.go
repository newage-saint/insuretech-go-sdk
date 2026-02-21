package models

import (
	"time"
)

// CommissionEarnedEvent represents a commission_earned_event
type CommissionEarnedEvent struct {
	EventId          string    `json:"event_id,omitempty"`
	CommissionId     string    `json:"commission_id,omitempty"`
	CommissionNumber string    `json:"commission_number,omitempty"`
	RecipientType    string    `json:"recipient_type,omitempty"`
	PolicyId         string    `json:"policy_id,omitempty"`
	CorrelationId    string    `json:"correlation_id,omitempty"`
	Type             string    `json:"type,omitempty"`
	RecipientId      string    `json:"recipient_id,omitempty"`
	Amount           *Money    `json:"amount,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}
