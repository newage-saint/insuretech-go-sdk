package models

import (
	"time"
)

// Report represents a report
type Report struct {
	ReportName  string      `json:"report_name"`
	Type        *ReportType `json:"type"`
	PeriodStart time.Time   `json:"period_start"`
	PeriodEnd   time.Time   `json:"period_end"`
	ReportUrl   string      `json:"report_url,omitempty"`
	Status      interface{} `json:"status"`
	ReportId    string      `json:"report_id"`
	ReportData  string      `json:"report_data"`
	GeneratedBy string      `json:"generated_by,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
