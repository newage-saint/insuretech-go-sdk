package models

import (
	"time"
)

// ComplianceCheckPerformedEvent represents a compliance_check_performed_event
type ComplianceCheckPerformedEvent struct {
	EventId         string    `json:"event_id,omitempty"`
	ComplianceLogId string    `json:"compliance_log_id,omitempty"`
	Regulation      string    `json:"regulation,omitempty"`
	Status          string    `json:"status,omitempty"`
	CorrelationId   string    `json:"correlation_id,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	Type            string    `json:"type,omitempty"`
	EntityType      string    `json:"entity_type,omitempty"`
	EntityId        string    `json:"entity_id,omitempty"`
}
