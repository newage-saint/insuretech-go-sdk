package models

import (
	"time"
)

// UserLoggedOutEvent represents a user_logged_out_event
type UserLoggedOutEvent struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	EventId   string    `json:"event_id,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	SessionId string    `json:"session_id,omitempty"`
}
