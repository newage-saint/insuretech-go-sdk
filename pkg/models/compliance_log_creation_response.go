package models

// ComplianceLogCreationResponse represents a compliance_log_creation_response
type ComplianceLogCreationResponse struct {
	Error           *Error `json:"error,omitempty"`
	ComplianceLogId string `json:"compliance_log_id,omitempty"`
}
