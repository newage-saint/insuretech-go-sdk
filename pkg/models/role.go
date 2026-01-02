package models

import (
	"time"
)

// Role represents a role
type Role struct {
	Type        interface{} `json:"type"`
	RoleId      string      `json:"role_id"`
	DisplayName string      `json:"display_name"`
	Description string      `json:"description,omitempty"`
	Permissions []string    `json:"permissions,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	IsActive    bool        `json:"is_active"`
	RoleName    string      `json:"role_name"`
}
