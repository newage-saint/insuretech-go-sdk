package models

import (
	"time"
)

// AuditEvent represents a audit_event
type AuditEvent struct {
	Category    *EventCategory `json:"category"`
	EventType   string         `json:"event_type"`
	Description string         `json:"description"`
	UserId      string         `json:"user_id,omitempty"`
	EntityType  string         `json:"entity_type,omitempty"`
	EntityId    string         `json:"entity_id,omitempty"`
	Metadata    string         `json:"metadata,omitempty"`
	Timestamp   time.Time      `json:"timestamp"`
	Id          string         `json:"id"`
	Severity    *EventSeverity `json:"severity"`
}
