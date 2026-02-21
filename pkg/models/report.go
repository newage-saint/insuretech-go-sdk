package models

import (
	"time"
)

// Report represents a report
type Report struct {
	UpdatedAt   time.Time   `json:"updated_at"`
	PeriodEnd   time.Time   `json:"period_end"`
	Status      interface{} `json:"status"`
	GeneratedBy string      `json:"generated_by,omitempty"`
	ReportId    string      `json:"report_id"`
	ReportName  string      `json:"report_name"`
	Type        *ReportType `json:"type"`
	PeriodStart time.Time   `json:"period_start"`
	ReportData  string      `json:"report_data"`
	ReportUrl   string      `json:"report_url,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
}
