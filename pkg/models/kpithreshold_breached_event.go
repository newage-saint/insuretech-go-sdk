package models

import (
	"time"
)

// KPIThresholdBreachedEvent represents a kpithreshold_breached_event
type KPIThresholdBreachedEvent struct {
	CurrentValue       float64   `json:"current_value,omitempty"`
	ThresholdValue     float64   `json:"threshold_value,omitempty"`
	ThresholdType      string    `json:"threshold_type,omitempty"`
	NotificationSentTo string    `json:"notification_sent_to,omitempty"`
	CorrelationId      string    `json:"correlation_id,omitempty"`
	EventId            string    `json:"event_id,omitempty"`
	MetricName         string    `json:"metric_name,omitempty"`
	Severity           string    `json:"severity,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}
