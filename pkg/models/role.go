package models

import (
	"time"
)

// Role represents a role
type Role struct {
	RoleId      string      `json:"role_id"`
	RoleName    string      `json:"role_name"`
	DisplayName string      `json:"display_name"`
	Permissions []string    `json:"permissions,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	IsActive    bool        `json:"is_active"`
	Description string      `json:"description,omitempty"`
	Type        interface{} `json:"type"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
