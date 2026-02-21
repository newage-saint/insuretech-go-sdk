package models

// RefundExecutionResponse represents a refund_execution_response
type RefundExecutionResponse struct {
	Message             string `json:"message,omitempty"`
	Error               *Error `json:"error,omitempty"`
	RefundTransactionId string `json:"refund_transaction_id,omitempty"`
}
