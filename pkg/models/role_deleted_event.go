package models

import (
	"time"
)

// RoleDeletedEvent represents a role_deleted_event
type RoleDeletedEvent struct {
	RoleId    string    `json:"role_id,omitempty"`
	RoleName  string    `json:"role_name,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	DeletedBy string    `json:"deleted_by,omitempty"`
	EventId   string    `json:"event_id,omitempty"`
}
