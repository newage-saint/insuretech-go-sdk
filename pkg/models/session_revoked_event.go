package models

import (
	"time"
)

// SessionRevokedEvent represents a session_revoked_event
type SessionRevokedEvent struct {
	SessionId   string    `json:"session_id,omitempty"`
	SessionType string    `json:"session_type,omitempty"`
	RevokedBy   string    `json:"revoked_by,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	EventId     string    `json:"event_id,omitempty"`
	UserId      string    `json:"user_id,omitempty"`
}
