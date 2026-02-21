package models

import (
	"time"
)

// PolicyRevivedEvent represents a policy_revived_event
type PolicyRevivedEvent struct {
	EventId       string    `json:"event_id,omitempty"`
	PolicyId      string    `json:"policy_id,omitempty"`
	GracePeriodId string    `json:"grace_period_id,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}
