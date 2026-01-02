package models

import (
	"time"
)

// Telemetry represents a telemetry
type Telemetry struct {
	DeviceId    string                 `json:"device_id"`
	Type        *TelemetryType         `json:"type"`
	Longitude   float64                `json:"longitude,omitempty"`
	Location    *Location              `json:"location,omitempty"`
	Timestamp   time.Time              `json:"timestamp"`
	Metrics     map[string]interface{} `json:"metrics,omitempty"`
	Latitude    float64                `json:"latitude,omitempty"`
	RawData     string                 `json:"raw_data,omitempty"`
	TelemetryId string                 `json:"telemetry_id"`
}
