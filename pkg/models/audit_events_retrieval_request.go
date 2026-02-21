package models

// AuditEventsRetrievalRequest represents a audit_events_retrieval_request
type AuditEventsRetrievalRequest struct {
	Severity  string `json:"severity,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
	Category  string `json:"category"`
}
