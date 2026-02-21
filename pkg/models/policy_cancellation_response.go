package models

// PolicyCancellationResponse represents a policy_cancellation_response
type PolicyCancellationResponse struct {
	RefundAmount *Money `json:"refund_amount,omitempty"`
	Error        *Error `json:"error,omitempty"`
	Message      string `json:"message,omitempty"`
}
