package models

import (
	"time"
)

// RoleAssignment represents a role_assignment
type RoleAssignment struct {
	AssignmentId string    `json:"assignment_id"`
	UserId       string    `json:"user_id"`
	RoleId       string    `json:"role_id"`
	TenantId     string    `json:"tenant_id,omitempty"`
	AssignedAt   time.Time `json:"assigned_at"`
	AssignedBy   string    `json:"assigned_by"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
}
