package models

import (
	"time"
)

// InsurerCreatedEvent represents a insurer_created_event
type InsurerCreatedEvent struct {
	InsurerCode   string    `json:"insurer_code,omitempty"`
	InsurerName   string    `json:"insurer_name,omitempty"`
	Type          string    `json:"type,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	InsurerId     string    `json:"insurer_id,omitempty"`
}
