package models

import (
	"time"
)

// Payment represents a payment
type Payment struct {
	PayeeId               string                `json:"payee_id,omitempty"`
	GatewayResponse       string                `json:"gateway_response,omitempty"`
	TransactionId         string                `json:"transaction_id,omitempty"`
	Type                  *PaymentType          `json:"type"`
	InitiatedAt           time.Time             `json:"initiated_at"`
	Gateway               string                `json:"gateway,omitempty"`
	ReceiptUrl            string                `json:"receipt_url,omitempty"`
	ClaimId               string                `json:"claim_id,omitempty"`
	Method                *PaymentPaymentMethod `json:"method"`
	Currency              string                `json:"currency"`
	CreatedAt             time.Time             `json:"created_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
	RetryCount            int                   `json:"retry_count"`
	FailureReason         string                `json:"failure_reason,omitempty"`
	TigerbeetleTransferId string                `json:"tigerbeetle_transfer_id,omitempty"`
	Status                interface{}           `json:"status"`
	Amount                *Money                `json:"amount"`
	CompletedAt           time.Time             `json:"completed_at,omitempty"`
	IdempotencyKey        string                `json:"idempotency_key,omitempty"`
	PaymentId             string                `json:"payment_id"`
	PolicyId              string                `json:"policy_id,omitempty"`
	PayerId               string                `json:"payer_id"`
}
