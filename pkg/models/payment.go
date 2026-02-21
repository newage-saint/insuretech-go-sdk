package models

import (
	"time"
)

// Payment represents a payment
type Payment struct {
	CreatedAt             time.Time             `json:"created_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
	TigerbeetleTransferId string                `json:"tigerbeetle_transfer_id,omitempty"`
	RetryCount            int                   `json:"retry_count"`
	IdempotencyKey        string                `json:"idempotency_key,omitempty"`
	PayerId               string                `json:"payer_id"`
	InitiatedAt           time.Time             `json:"initiated_at"`
	Gateway               string                `json:"gateway,omitempty"`
	GatewayResponse       string                `json:"gateway_response,omitempty"`
	ReceiptUrl            string                `json:"receipt_url,omitempty"`
	TransactionId         string                `json:"transaction_id,omitempty"`
	Type                  *PaymentType          `json:"type"`
	Status                interface{}           `json:"status"`
	Currency              string                `json:"currency"`
	CompletedAt           time.Time             `json:"completed_at,omitempty"`
	FailureReason         string                `json:"failure_reason,omitempty"`
	PaymentId             string                `json:"payment_id"`
	PolicyId              string                `json:"policy_id,omitempty"`
	ClaimId               string                `json:"claim_id,omitempty"`
	Method                *PaymentPaymentMethod `json:"method"`
	Amount                *Money                `json:"amount"`
	PayeeId               string                `json:"payee_id,omitempty"`
}
