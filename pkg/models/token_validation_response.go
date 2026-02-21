package models

import (
	"time"
)

// TokenValidationResponse represents a token_validation_response
type TokenValidationResponse struct {
	Roles       []string  `json:"roles,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
	Error       *Error    `json:"error,omitempty"`
	Valid       bool      `json:"valid,omitempty"`
	UserId      string    `json:"user_id,omitempty"`
	SessionId   string    `json:"session_id,omitempty"`
	SessionType string    `json:"session_type,omitempty"`
}
