package models

import (
	"time"
)

// AccessDeniedEvent represents a access_denied_event
type AccessDeniedEvent struct {
	UserId      string    `json:"user_id,omitempty"`
	Resource    string    `json:"resource,omitempty"`
	Action      string    `json:"action,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	IpAddress   string    `json:"ip_address,omitempty"`
	RequestPath string    `json:"request_path,omitempty"`
	EventId     string    `json:"event_id,omitempty"`
}
