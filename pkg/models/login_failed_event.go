package models

import (
	"time"
)

// LoginFailedEvent represents a login_failed_event
type LoginFailedEvent struct {
	FailedAttemptsCount int       `json:"failed_attempts_count,omitempty"`
	EventId             string    `json:"event_id,omitempty"`
	UserId              string    `json:"user_id,omitempty"`
	MobileNumber        string    `json:"mobile_number,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
	IpAddress           string    `json:"ip_address,omitempty"`
	FailureReason       string    `json:"failure_reason,omitempty"`
	DeviceType          string    `json:"device_type,omitempty"`
	UserAgent           string    `json:"user_agent,omitempty"`
}
