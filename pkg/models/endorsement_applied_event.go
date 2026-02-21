package models

import (
	"time"
)

// EndorsementAppliedEvent represents a endorsement_applied_event
type EndorsementAppliedEvent struct {
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	EndorsementId string    `json:"endorsement_id,omitempty"`
	PolicyId      string    `json:"policy_id,omitempty"`
	Type          string    `json:"type,omitempty"`
}
