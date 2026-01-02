package models

import (
	"time"
)

// DataExportCompletedEvent represents a data_export_completed_event
type DataExportCompletedEvent struct {
	Timestamp             time.Time `json:"timestamp,omitempty"`
	CorrelationId         string    `json:"correlation_id,omitempty"`
	ExportId              string    `json:"export_id,omitempty"`
	ExportUrl             string    `json:"export_url,omitempty"`
	RecordCount           string    `json:"record_count,omitempty"`
	GenerationTimeSeconds int       `json:"generation_time_seconds,omitempty"`
	EventId               string    `json:"event_id,omitempty"`
	UserId                string    `json:"user_id,omitempty"`
	FileSizeBytes         string    `json:"file_size_bytes,omitempty"`
}
