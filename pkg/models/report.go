package models

import (
	"time"
)

// Report represents a report
type Report struct {
	PeriodStart time.Time   `json:"period_start"`
	PeriodEnd   time.Time   `json:"period_end"`
	ReportData  string      `json:"report_data"`
	ReportUrl   string      `json:"report_url,omitempty"`
	Status      interface{} `json:"status"`
	GeneratedBy string      `json:"generated_by,omitempty"`
	ReportName  string      `json:"report_name"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ReportId    string      `json:"report_id"`
	Type        *ReportType `json:"type"`
}
