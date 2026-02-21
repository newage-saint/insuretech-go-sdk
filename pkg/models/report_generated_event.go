package models

import (
	"time"
)

// ReportGeneratedEvent represents a report_generated_event
type ReportGeneratedEvent struct {
	ReportId              string    `json:"report_id,omitempty"`
	ReportName            string    `json:"report_name,omitempty"`
	PeriodStart           time.Time `json:"period_start,omitempty"`
	PeriodEnd             time.Time `json:"period_end,omitempty"`
	GenerationTimeSeconds int       `json:"generation_time_seconds,omitempty"`
	GeneratedBy           string    `json:"generated_by,omitempty"`
	ReportType            string    `json:"report_type,omitempty"`
	ReportUrl             string    `json:"report_url,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
	CorrelationId         string    `json:"correlation_id,omitempty"`
	EventId               string    `json:"event_id,omitempty"`
}
