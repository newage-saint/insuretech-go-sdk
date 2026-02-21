package models

import (
	"time"
)

// Payment represents a payment
type Payment struct {
	FailureReason         string                `json:"failure_reason,omitempty"`
	CompletedAt           time.Time             `json:"completed_at,omitempty"`
	PaymentId             string                `json:"payment_id"`
	TransactionId         string                `json:"transaction_id,omitempty"`
	PolicyId              string                `json:"policy_id,omitempty"`
	Amount                *Money                `json:"amount"`
	Currency              string                `json:"currency"`
	PayeeId               string                `json:"payee_id,omitempty"`
	CreatedAt             time.Time             `json:"created_at"`
	TigerbeetleTransferId string                `json:"tigerbeetle_transfer_id,omitempty"`
	Status                interface{}           `json:"status"`
	PayerId               string                `json:"payer_id"`
	UpdatedAt             time.Time             `json:"updated_at"`
	RetryCount            int                   `json:"retry_count"`
	IdempotencyKey        string                `json:"idempotency_key,omitempty"`
	Method                *PaymentPaymentMethod `json:"method"`
	InitiatedAt           time.Time             `json:"initiated_at"`
	Gateway               string                `json:"gateway,omitempty"`
	ClaimId               string                `json:"claim_id,omitempty"`
	Type                  *PaymentType          `json:"type"`
	GatewayResponse       string                `json:"gateway_response,omitempty"`
	ReceiptUrl            string                `json:"receipt_url,omitempty"`
}
