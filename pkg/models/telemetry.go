package models

import (
	"time"
)

// Telemetry represents a telemetry
type Telemetry struct {
	TelemetryId string                 `json:"telemetry_id"`
	Type        *TelemetryType         `json:"type"`
	Metrics     map[string]interface{} `json:"metrics,omitempty"`
	Latitude    float64                `json:"latitude,omitempty"`
	Longitude   float64                `json:"longitude,omitempty"`
	Location    *Location              `json:"location,omitempty"`
	DeviceId    string                 `json:"device_id"`
	Timestamp   time.Time              `json:"timestamp"`
	RawData     string                 `json:"raw_data,omitempty"`
}
