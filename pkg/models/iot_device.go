package models

import (
	"time"
)

// IoTDevice represents a iot_device
type IoTDevice struct {
	LastSeenAt   time.Time              `json:"last_seen_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeviceSerial string                 `json:"device_serial"`
	PolicyId     string                 `json:"policy_id,omitempty"`
	OwnerId      string                 `json:"owner_id"`
	Status       interface{}            `json:"status"`
	RegisteredAt time.Time              `json:"registered_at"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	DeviceId     string                 `json:"device_id"`
	Type         *IotDeviceType         `json:"type"`
	Manufacturer string                 `json:"manufacturer"`
	Model        string                 `json:"model"`
}
