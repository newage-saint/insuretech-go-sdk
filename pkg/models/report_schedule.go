package models

import (
	"time"
)

// ReportSchedule represents a report_schedule
type ReportSchedule struct {
	Id                 string             `json:"id"`
	Frequency          *ScheduleFrequency `json:"frequency"`
	CronExpression     string             `json:"cron_expression,omitempty"`
	Recipients         []string           `json:"recipients,omitempty"`
	AuditInfo          interface{}        `json:"audit_info"`
	ReportDefinitionId string             `json:"report_definition_id"`
	Name               string             `json:"name"`
	Parameters         string             `json:"parameters,omitempty"`
	IsActive           bool               `json:"is_active,omitempty"`
	NextRunAt          time.Time          `json:"next_run_at,omitempty"`
	LastRunAt          time.Time          `json:"last_run_at,omitempty"`
}
