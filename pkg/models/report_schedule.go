package models

import (
	"time"
)

// ReportSchedule represents a report_schedule
type ReportSchedule struct {
	Name               string             `json:"name"`
	Frequency          *ScheduleFrequency `json:"frequency"`
	Parameters         string             `json:"parameters,omitempty"`
	Recipients         []string           `json:"recipients,omitempty"`
	IsActive           bool               `json:"is_active,omitempty"`
	LastRunAt          time.Time          `json:"last_run_at,omitempty"`
	Id                 string             `json:"id"`
	ReportDefinitionId string             `json:"report_definition_id"`
	CronExpression     string             `json:"cron_expression,omitempty"`
	NextRunAt          time.Time          `json:"next_run_at,omitempty"`
	AuditInfo          *AuditInfo         `json:"audit_info,omitempty"`
}
