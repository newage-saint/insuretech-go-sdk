package models

// InitiateRefundRequest represents a initiate_refund_request
type InitiateRefundRequest struct {
	RefundAmount *Money `json:"refund_amount,omitempty"`
	Reason       string `json:"reason,omitempty"`
	InitiatedBy  string `json:"initiated_by,omitempty"`
	PaymentId    string `json:"payment_id"`
}
