package models

import (
	"time"
)

// EmailVerifiedEvent represents a email_verified_event
type EmailVerifiedEvent struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	EventId   string    `json:"event_id,omitempty"`
	UserId    string    `json:"user_id,omitempty"`
	Email     string    `json:"email,omitempty"`
}
