package models

import (
	"time"
)

// ReportExecution represents a report_execution
type ReportExecution struct {
	Id                 string      `json:"id"`
	ReportDefinitionId string      `json:"report_definition_id"`
	Parameters         string      `json:"parameters,omitempty"`
	FileUrl            string      `json:"file_url,omitempty"`
	FileFormat         string      `json:"file_format,omitempty"`
	FileSizeBytes      string      `json:"file_size_bytes,omitempty"`
	RowCount           int         `json:"row_count,omitempty"`
	ErrorMessage       string      `json:"error_message,omitempty"`
	ReportScheduleId   string      `json:"report_schedule_id,omitempty"`
	Status             interface{} `json:"status"`
	ExecutedBy         string      `json:"executed_by,omitempty"`
	StartedAt          time.Time   `json:"started_at"`
	CompletedAt        time.Time   `json:"completed_at,omitempty"`
	AuditInfo          *AuditInfo  `json:"audit_info,omitempty"`
}
