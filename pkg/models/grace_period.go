package models

import (
	"time"
)

// GracePeriod represents a grace_period
type GracePeriod struct {
	StartDate      time.Time   `json:"start_date"`
	EndDate        time.Time   `json:"end_date"`
	DaysRemaining  int         `json:"days_remaining,omitempty"`
	RevivedAt      time.Time   `json:"revived_at,omitempty"`
	Status         interface{} `json:"status"`
	CoverageActive bool        `json:"coverage_active,omitempty"`
	AuditInfo      interface{} `json:"audit_info"`
	Id             string      `json:"id"`
	PolicyId       string      `json:"policy_id"`
}
