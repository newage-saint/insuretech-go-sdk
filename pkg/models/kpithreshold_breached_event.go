package models

import (
	"time"
)

// KPIThresholdBreachedEvent represents a kpithreshold_breached_event
type KPIThresholdBreachedEvent struct {
	MetricName         string    `json:"metric_name,omitempty"`
	ThresholdValue     float64   `json:"threshold_value,omitempty"`
	ThresholdType      string    `json:"threshold_type,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	EventId            string    `json:"event_id,omitempty"`
	CurrentValue       float64   `json:"current_value,omitempty"`
	Severity           string    `json:"severity,omitempty"`
	NotificationSentTo string    `json:"notification_sent_to,omitempty"`
	CorrelationId      string    `json:"correlation_id,omitempty"`
}
