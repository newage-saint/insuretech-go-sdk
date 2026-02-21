package models

import (
	"time"
)

// ReportExecution represents a report_execution
type ReportExecution struct {
	RowCount           int         `json:"row_count,omitempty"`
	ExecutedBy         string      `json:"executed_by,omitempty"`
	StartedAt          time.Time   `json:"started_at"`
	AuditInfo          interface{} `json:"audit_info"`
	Id                 string      `json:"id"`
	ReportScheduleId   string      `json:"report_schedule_id,omitempty"`
	Status             interface{} `json:"status"`
	FileUrl            string      `json:"file_url,omitempty"`
	FileFormat         string      `json:"file_format,omitempty"`
	ErrorMessage       string      `json:"error_message,omitempty"`
	CompletedAt        time.Time   `json:"completed_at,omitempty"`
	ReportDefinitionId string      `json:"report_definition_id"`
	Parameters         string      `json:"parameters,omitempty"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
}
