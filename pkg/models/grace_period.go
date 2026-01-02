package models

import (
	"time"
)

// GracePeriod represents a grace_period
type GracePeriod struct {
	DaysRemaining  int         `json:"days_remaining,omitempty"`
	Status         interface{} `json:"status"`
	StartDate      time.Time   `json:"start_date"`
	EndDate        time.Time   `json:"end_date"`
	CoverageActive bool        `json:"coverage_active,omitempty"`
	RevivedAt      time.Time   `json:"revived_at,omitempty"`
	AuditInfo      *AuditInfo  `json:"audit_info,omitempty"`
	Id             string      `json:"id"`
	PolicyId       string      `json:"policy_id"`
}
