package models

import (
	"time"
)

// Telemetry represents a telemetry
type Telemetry struct {
	TelemetryId string                 `json:"telemetry_id"`
	DeviceId    string                 `json:"device_id"`
	Metrics     map[string]interface{} `json:"metrics,omitempty"`
	Location    *Location              `json:"location,omitempty"`
	Timestamp   time.Time              `json:"timestamp"`
	Type        *TelemetryType         `json:"type"`
	Latitude    float64                `json:"latitude,omitempty"`
	Longitude   float64                `json:"longitude,omitempty"`
	RawData     string                 `json:"raw_data,omitempty"`
}
