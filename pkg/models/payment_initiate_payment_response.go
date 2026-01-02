package models

import (
	"time"
)

// PaymentInitiatePaymentResponse represents a payment_initiate_payment_response
type PaymentInitiatePaymentResponse struct {
	PaymentId     string    `json:"payment_id,omitempty"`
	TransactionId string    `json:"transaction_id,omitempty"`
	PaymentUrl    string    `json:"payment_url,omitempty"`
	Status        string    `json:"status,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	Error         *Error    `json:"error,omitempty"`
}
