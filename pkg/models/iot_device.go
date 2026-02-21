package models

import (
	"time"
)

// IoTDevice represents a iot_device
type IoTDevice struct {
	OwnerId      string                 `json:"owner_id"`
	CreatedAt    time.Time              `json:"created_at"`
	Type         *IotDeviceType         `json:"type"`
	Model        string                 `json:"model"`
	Status       interface{}            `json:"status"`
	RegisteredAt time.Time              `json:"registered_at"`
	LastSeenAt   time.Time              `json:"last_seen_at,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeviceId     string                 `json:"device_id"`
	DeviceSerial string                 `json:"device_serial"`
	Manufacturer string                 `json:"manufacturer"`
	PolicyId     string                 `json:"policy_id,omitempty"`
}
