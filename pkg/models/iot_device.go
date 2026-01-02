package models

import (
	"time"
)

// IoTDevice represents a iot_device
type IoTDevice struct {
	CreatedAt    time.Time              `json:"created_at"`
	DeviceSerial string                 `json:"device_serial"`
	Type         *IotDeviceType         `json:"type"`
	Manufacturer string                 `json:"manufacturer"`
	RegisteredAt time.Time              `json:"registered_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeviceId     string                 `json:"device_id"`
	Model        string                 `json:"model"`
	PolicyId     string                 `json:"policy_id,omitempty"`
	OwnerId      string                 `json:"owner_id"`
	Status       interface{}            `json:"status"`
	LastSeenAt   time.Time              `json:"last_seen_at,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}
