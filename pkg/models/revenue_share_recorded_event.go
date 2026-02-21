package models

import (
	"time"
)

// RevenueShareRecordedEvent represents a revenue_share_recorded_event
type RevenueShareRecordedEvent struct {
	RevenueShareId string    `json:"revenue_share_id,omitempty"`
	PolicyId       string    `json:"policy_id,omitempty"`
	InsurerId      string    `json:"insurer_id,omitempty"`
	PlatformShare  *Money    `json:"platform_share,omitempty"`
	InsurerShare   *Money    `json:"insurer_share,omitempty"`
	CorrelationId  string    `json:"correlation_id,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	RevenueModel   string    `json:"revenue_model,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
}
