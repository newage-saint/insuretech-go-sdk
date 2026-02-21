package models

// RequestRefundResponse represents a request_refund_response
type RequestRefundResponse struct {
	Message      string `json:"message,omitempty"`
	Error        *Error `json:"error,omitempty"`
	RefundId     string `json:"refund_id,omitempty"`
	RefundNumber string `json:"refund_number,omitempty"`
}
