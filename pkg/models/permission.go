package models

import (
	"time"
)

// Permission represents a permission
type Permission struct {
	Description    string    `json:"description,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	PermissionId   string    `json:"permission_id"`
	PermissionName string    `json:"permission_name"`
	Resource       string    `json:"resource"`
	Action         string    `json:"action"`
}
