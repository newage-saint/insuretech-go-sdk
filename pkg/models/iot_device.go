package models

import (
	"time"
)

// IoTDevice represents a iot_device
type IoTDevice struct {
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeviceId     string                 `json:"device_id"`
	DeviceSerial string                 `json:"device_serial"`
	Type         *IotDeviceType         `json:"type"`
	Manufacturer string                 `json:"manufacturer"`
	Model        string                 `json:"model"`
	Status       interface{}            `json:"status"`
	RegisteredAt time.Time              `json:"registered_at"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	PolicyId     string                 `json:"policy_id,omitempty"`
	OwnerId      string                 `json:"owner_id"`
	LastSeenAt   time.Time              `json:"last_seen_at,omitempty"`
}
