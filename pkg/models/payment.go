package models

import (
	"time"
)

// Payment represents a payment
type Payment struct {
	TigerbeetleTransferId string                `json:"tigerbeetle_transfer_id,omitempty"`
	Type                  *PaymentType          `json:"type"`
	Currency              string                `json:"currency"`
	InitiatedAt           time.Time             `json:"initiated_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
	ClaimId               string                `json:"claim_id,omitempty"`
	Status                interface{}           `json:"status"`
	Amount                *Money                `json:"amount"`
	PayerId               string                `json:"payer_id"`
	CompletedAt           time.Time             `json:"completed_at,omitempty"`
	RetryCount            int                   `json:"retry_count"`
	PaymentId             string                `json:"payment_id"`
	Method                *PaymentPaymentMethod `json:"method"`
	CreatedAt             time.Time             `json:"created_at"`
	FailureReason         string                `json:"failure_reason,omitempty"`
	IdempotencyKey        string                `json:"idempotency_key,omitempty"`
	TransactionId         string                `json:"transaction_id,omitempty"`
	PolicyId              string                `json:"policy_id,omitempty"`
	PayeeId               string                `json:"payee_id,omitempty"`
	Gateway               string                `json:"gateway,omitempty"`
	GatewayResponse       string                `json:"gateway_response,omitempty"`
	ReceiptUrl            string                `json:"receipt_url,omitempty"`
}
