package models

// UnderwritingRejectionResponse represents a underwriting_rejection_response
type UnderwritingRejectionResponse struct {
	Message    string `json:"message,omitempty"`
	Error      *Error `json:"error,omitempty"`
	DecisionId string `json:"decision_id,omitempty"`
}
