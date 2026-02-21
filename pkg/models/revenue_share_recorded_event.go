package models

import (
	"time"
)

// RevenueShareRecordedEvent represents a revenue_share_recorded_event
type RevenueShareRecordedEvent struct {
	InsurerId      string    `json:"insurer_id,omitempty"`
	InsurerShare   *Money    `json:"insurer_share,omitempty"`
	RevenueModel   string    `json:"revenue_model,omitempty"`
	CorrelationId  string    `json:"correlation_id,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
	PlatformShare  *Money    `json:"platform_share,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	RevenueShareId string    `json:"revenue_share_id,omitempty"`
	PolicyId       string    `json:"policy_id,omitempty"`
}
