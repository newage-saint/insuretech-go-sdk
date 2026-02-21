package models

import (
	"time"
)

// DeviceActivatedEvent represents a device_activated_event
type DeviceActivatedEvent struct {
	EventId      string    `json:"event_id,omitempty"`
	DeviceId     string    `json:"device_id,omitempty"`
	DeviceSerial string    `json:"device_serial,omitempty"`
	PolicyId     string    `json:"policy_id,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}
