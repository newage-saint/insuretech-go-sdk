package models

import (
	"time"
)

// RolePermission represents a role_permission
type RolePermission struct {
	PermissionId string    `json:"permission_id"`
	AssignedAt   time.Time `json:"assigned_at"`
	RoleId       string    `json:"role_id"`
}
