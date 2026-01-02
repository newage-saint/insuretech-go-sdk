package models

import (
	"time"
)

// User represents a user
type User struct {
	PasswordHash  string      `json:"password_hash"`
	Status        interface{} `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	LastLoginAt   time.Time   `json:"last_login_at,omitempty"`
	CreatedBy     string      `json:"created_by,omitempty"`
	UpdatedBy     string      `json:"updated_by,omitempty"`
	LoginAttempts int         `json:"login_attempts"`
	UserId        string      `json:"user_id"`
	MobileNumber  string      `json:"mobile_number"`
	Email         string      `json:"email,omitempty"`
	LockedUntil   time.Time   `json:"locked_until,omitempty"`
	DeletedAt     time.Time   `json:"deleted_at,omitempty"`
}
