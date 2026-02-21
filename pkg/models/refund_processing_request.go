package models

// RefundProcessingRequest represents a refund_processing_request
type RefundProcessingRequest struct {
	PaymentReference string `json:"payment_reference,omitempty"`
	RefundId         string `json:"refund_id"`
	PaymentMethod    string `json:"payment_method,omitempty"`
}
