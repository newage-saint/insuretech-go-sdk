package models

import (
	"time"
)

// Role represents a role
type Role struct {
	Permissions []string    `json:"permissions,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	IsActive    bool        `json:"is_active"`
	RoleId      string      `json:"role_id"`
	Description string      `json:"description,omitempty"`
	RoleName    string      `json:"role_name"`
	DisplayName string      `json:"display_name"`
	Type        interface{} `json:"type"`
}
