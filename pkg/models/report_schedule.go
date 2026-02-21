package models

import (
	"time"
)

// ReportSchedule represents a report_schedule
type ReportSchedule struct {
	CronExpression     string             `json:"cron_expression,omitempty"`
	Parameters         string             `json:"parameters,omitempty"`
	Recipients         []string           `json:"recipients,omitempty"`
	NextRunAt          time.Time          `json:"next_run_at,omitempty"`
	Id                 string             `json:"id"`
	Name               string             `json:"name"`
	IsActive           bool               `json:"is_active,omitempty"`
	LastRunAt          time.Time          `json:"last_run_at,omitempty"`
	AuditInfo          interface{}        `json:"audit_info"`
	ReportDefinitionId string             `json:"report_definition_id"`
	Frequency          *ScheduleFrequency `json:"frequency"`
}
