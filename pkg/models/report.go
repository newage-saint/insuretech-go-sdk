package models

import (
	"time"
)

// Report represents a report
type Report struct {
	ReportId    string      `json:"report_id"`
	ReportName  string      `json:"report_name"`
	Type        *ReportType `json:"type"`
	ReportUrl   string      `json:"report_url,omitempty"`
	GeneratedBy string      `json:"generated_by,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at"`
	PeriodStart time.Time   `json:"period_start"`
	PeriodEnd   time.Time   `json:"period_end"`
	ReportData  string      `json:"report_data"`
	Status      interface{} `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}
