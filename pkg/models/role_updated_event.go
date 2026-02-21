package models

import (
	"time"
)

// RoleUpdatedEvent represents a role_updated_event
type RoleUpdatedEvent struct {
	RoleId             string    `json:"role_id,omitempty"`
	RoleName           string    `json:"role_name,omitempty"`
	AddedPermissions   []string  `json:"added_permissions,omitempty"`
	RemovedPermissions []string  `json:"removed_permissions,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	UpdatedBy          string    `json:"updated_by,omitempty"`
	EventId            string    `json:"event_id,omitempty"`
}
