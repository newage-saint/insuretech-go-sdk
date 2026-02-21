package models

import (
	"time"
)

// RoleAssignedEvent represents a role_assigned_event
type RoleAssignedEvent struct {
	TenantId     string    `json:"tenant_id,omitempty"`
	AssignedBy   string    `json:"assigned_by,omitempty"`
	RoleId       string    `json:"role_id,omitempty"`
	RoleName     string    `json:"role_name,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
	EventId      string    `json:"event_id,omitempty"`
	AssignmentId string    `json:"assignment_id,omitempty"`
	UserId       string    `json:"user_id,omitempty"`
}
