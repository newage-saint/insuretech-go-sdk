package models

import (
	"time"
)

// DeviceRegisteredEvent represents a device_registered_event
type DeviceRegisteredEvent struct {
	Timestamp    time.Time `json:"timestamp,omitempty"`
	DeviceType   string    `json:"device_type,omitempty"`
	OwnerId      string    `json:"owner_id,omitempty"`
	EventId      string    `json:"event_id,omitempty"`
	DeviceId     string    `json:"device_id,omitempty"`
	DeviceSerial string    `json:"device_serial,omitempty"`
	Manufacturer string    `json:"manufacturer,omitempty"`
	Model        string    `json:"model,omitempty"`
	PolicyId     string    `json:"policy_id,omitempty"`
}
