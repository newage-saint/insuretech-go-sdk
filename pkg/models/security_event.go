package models

import (
	"time"
)

// SecurityEvent represents a security_event
type SecurityEvent struct {
	ResourceId   string                 `json:"resource_id,omitempty"`
	Action       string                 `json:"action,omitempty"`
	UserAgent    string                 `json:"user_agent,omitempty"`
	Timestamp    time.Time              `json:"timestamp,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	EventId      string                 `json:"event_id,omitempty"`
	EventType    string                 `json:"event_type,omitempty"`
	UserId       string                 `json:"user_id,omitempty"`
	ResourceType string                 `json:"resource_type,omitempty"`
	Authorized   bool                   `json:"authorized,omitempty"`
	IpAddress    string                 `json:"ip_address,omitempty"`
}
