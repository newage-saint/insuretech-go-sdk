package models

// ScheduleReportResponse represents a schedule_report_response
type ScheduleReportResponse struct {
	Error     *Error `json:"error,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
	Message   string `json:"message,omitempty"`
	NextRunAt string `json:"next_run_at,omitempty"`
}
