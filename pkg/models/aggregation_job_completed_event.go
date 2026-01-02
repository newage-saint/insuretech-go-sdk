package models

import (
	"time"
)

// AggregationJobCompletedEvent represents a aggregation_job_completed_event
type AggregationJobCompletedEvent struct {
	ExecutionTimeSeconds int       `json:"execution_time_seconds,omitempty"`
	CorrelationId        string    `json:"correlation_id,omitempty"`
	PeriodEnd            time.Time `json:"period_end,omitempty"`
	RecordsProcessed     string    `json:"records_processed,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
	EventId              string    `json:"event_id,omitempty"`
	JobId                string    `json:"job_id,omitempty"`
	JobType              string    `json:"job_type,omitempty"`
	PeriodStart          time.Time `json:"period_start,omitempty"`
}
