package models

// ReportDownloadResponse represents a report_download_response
type ReportDownloadResponse struct {
	Filename    string `json:"filename,omitempty"`
	Error       *Error `json:"error,omitempty"`
	Content     string `json:"content,omitempty"`
	ContentType string `json:"content_type,omitempty"`
}
