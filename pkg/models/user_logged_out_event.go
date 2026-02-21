package models

import (
	"time"
)

// UserLoggedOutEvent represents a user_logged_out_event
type UserLoggedOutEvent struct {
	EventId      string    `json:"event_id,omitempty"`
	UserId       string    `json:"user_id,omitempty"`
	SessionId    string    `json:"session_id,omitempty"`
	SessionType  string    `json:"session_type,omitempty"`
	LogoutReason string    `json:"logout_reason,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	IpAddress    string    `json:"ip_address,omitempty"`
	DeviceType   string    `json:"device_type,omitempty"`
}
