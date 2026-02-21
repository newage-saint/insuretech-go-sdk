package models

import (
	"time"
)

// RoleCreatedEvent represents a role_created_event
type RoleCreatedEvent struct {
	RoleId      string    `json:"role_id,omitempty"`
	RoleName    string    `json:"role_name,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	RoleType    string    `json:"role_type,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	EventId     string    `json:"event_id,omitempty"`
}
