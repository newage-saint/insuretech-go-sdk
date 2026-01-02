package models

// ReportGenerationRequest represents a report_generation_request
type ReportGenerationRequest struct {
	ReportId  string `json:"report_id"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}
