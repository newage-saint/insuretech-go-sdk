package models

import (
	"time"
)

// GracePeriod represents a grace_period
type GracePeriod struct {
	Id             string      `json:"id"`
	PolicyId       string      `json:"policy_id"`
	StartDate      time.Time   `json:"start_date"`
	EndDate        time.Time   `json:"end_date"`
	DaysRemaining  int         `json:"days_remaining,omitempty"`
	Status         interface{} `json:"status"`
	AuditInfo      interface{} `json:"audit_info"`
	CoverageActive bool        `json:"coverage_active,omitempty"`
	RevivedAt      time.Time   `json:"revived_at,omitempty"`
}
