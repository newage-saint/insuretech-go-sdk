package models

import (
	"time"
)

// AnomalyDetectedEvent represents a anomaly_detected_event
type AnomalyDetectedEvent struct {
	DeviceId       string                 `json:"device_id,omitempty"`
	DeviceSerial   string                 `json:"device_serial,omitempty"`
	Longitude      float64                `json:"longitude,omitempty"`
	EventId        string                 `json:"event_id,omitempty"`
	AnomalyType    string                 `json:"anomaly_type,omitempty"`
	Severity       string                 `json:"severity,omitempty"`
	AnomalyDetails map[string]interface{} `json:"anomaly_details,omitempty"`
	Timestamp      time.Time              `json:"timestamp,omitempty"`
	Latitude       float64                `json:"latitude,omitempty"`
}
