package models

import (
	"time"
)

// AuditEvent represents a audit_event
type AuditEvent struct {
	Id          string         `json:"id"`
	EventType   string         `json:"event_type"`
	Severity    *EventSeverity `json:"severity"`
	EntityType  string         `json:"entity_type,omitempty"`
	EntityId    string         `json:"entity_id,omitempty"`
	Category    *EventCategory `json:"category"`
	Description string         `json:"description"`
	UserId      string         `json:"user_id,omitempty"`
	Metadata    string         `json:"metadata,omitempty"`
	Timestamp   time.Time      `json:"timestamp"`
}
