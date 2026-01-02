package models

import (
	"time"
)

// Report represents a report
type Report struct {
	PeriodStart time.Time   `json:"period_start"`
	CreatedAt   time.Time   `json:"created_at"`
	ReportId    string      `json:"report_id"`
	PeriodEnd   time.Time   `json:"period_end"`
	ReportData  string      `json:"report_data"`
	ReportUrl   string      `json:"report_url,omitempty"`
	Status      interface{} `json:"status"`
	GeneratedBy string      `json:"generated_by,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ReportName  string      `json:"report_name"`
	Type        *ReportType `json:"type"`
}
