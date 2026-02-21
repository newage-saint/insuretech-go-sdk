package models

import (
	"time"
)

// RoleAssignedEvent represents a role_assigned_event
type RoleAssignedEvent struct {
	UserId       string    `json:"user_id,omitempty"`
	RoleName     string    `json:"role_name,omitempty"`
	TenantId     string    `json:"tenant_id,omitempty"`
	AssignedBy   string    `json:"assigned_by,omitempty"`
	EventId      string    `json:"event_id,omitempty"`
	RoleId       string    `json:"role_id,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
	AssignmentId string    `json:"assignment_id,omitempty"`
}
