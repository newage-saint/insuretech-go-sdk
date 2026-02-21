package models

// ReportExecutionRetrievalResponse represents a report_execution_retrieval_response
type ReportExecutionRetrievalResponse struct {
	Error           *Error           `json:"error,omitempty"`
	ReportExecution *ReportExecution `json:"report_execution,omitempty"`
}
