package models

import (
	"time"
)

// UserRegisteredEvent represents a user_registered_event
type UserRegisteredEvent struct {
	UserId       string    `json:"user_id,omitempty"`
	MobileNumber string    `json:"mobile_number,omitempty"`
	Email        string    `json:"email,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	IpAddress    string    `json:"ip_address,omitempty"`
	DeviceType   string    `json:"device_type,omitempty"`
	EventId      string    `json:"event_id,omitempty"`
}
