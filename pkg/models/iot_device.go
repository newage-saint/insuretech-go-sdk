package models

import (
	"time"
)

// IoTDevice represents a iot_device
type IoTDevice struct {
	Status       interface{}            `json:"status"`
	LastSeenAt   time.Time              `json:"last_seen_at,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Model        string                 `json:"model"`
	PolicyId     string                 `json:"policy_id,omitempty"`
	RegisteredAt time.Time              `json:"registered_at"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeviceId     string                 `json:"device_id"`
	DeviceSerial string                 `json:"device_serial"`
	Type         *IotDeviceType         `json:"type"`
	Manufacturer string                 `json:"manufacturer"`
	OwnerId      string                 `json:"owner_id"`
}
