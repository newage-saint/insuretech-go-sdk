package models

// EndorsementApprovalResponse represents a endorsement_approval_response
type EndorsementApprovalResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
