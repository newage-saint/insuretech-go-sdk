package models

import (
	"time"
)

// AuditEvent represents a audit_event
type AuditEvent struct {
	Id          string         `json:"id"`
	Category    *EventCategory `json:"category"`
	Severity    *EventSeverity `json:"severity"`
	Description string         `json:"description"`
	UserId      string         `json:"user_id,omitempty"`
	Metadata    string         `json:"metadata,omitempty"`
	Timestamp   time.Time      `json:"timestamp"`
	EventType   string         `json:"event_type"`
	EntityType  string         `json:"entity_type,omitempty"`
	EntityId    string         `json:"entity_id,omitempty"`
}
