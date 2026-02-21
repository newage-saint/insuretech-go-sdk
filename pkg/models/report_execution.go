package models

import (
	"time"
)

// ReportExecution represents a report_execution
type ReportExecution struct {
	ErrorMessage       string      `json:"error_message,omitempty"`
	AuditInfo          interface{} `json:"audit_info"`
	Id                 string      `json:"id"`
	ReportDefinitionId string      `json:"report_definition_id"`
	FileFormat         string      `json:"file_format,omitempty"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	RowCount           int         `json:"row_count,omitempty"`
	ExecutedBy         string      `json:"executed_by,omitempty"`
	StartedAt          time.Time   `json:"started_at"`
	CompletedAt        time.Time   `json:"completed_at,omitempty"`
	ReportScheduleId   string      `json:"report_schedule_id,omitempty"`
	Parameters         string      `json:"parameters,omitempty"`
	Status             interface{} `json:"status"`
	FileUrl            string      `json:"file_url,omitempty"`
}
