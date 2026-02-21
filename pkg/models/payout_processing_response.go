package models

// PayoutProcessingResponse represents a payout_processing_response
type PayoutProcessingResponse struct {
	PaidAt  string `json:"paid_at,omitempty"`
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
