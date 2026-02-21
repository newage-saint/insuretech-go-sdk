package models

// PaymentVerificationRequest represents a payment_verification_request
type PaymentVerificationRequest struct {
	PaymentMethod  string `json:"payment_method,omitempty"`
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	PaymentId      string `json:"payment_id"`
	TransactionId  string `json:"transaction_id"`
}
