package models

import (
	"time"
)

// Payment represents a payment
type Payment struct {
	PayeeId               string                `json:"payee_id,omitempty"`
	TigerbeetleTransferId string                `json:"tigerbeetle_transfer_id,omitempty"`
	ClaimId               string                `json:"claim_id,omitempty"`
	Type                  *PaymentType          `json:"type"`
	Method                *PaymentPaymentMethod `json:"method"`
	Status                interface{}           `json:"status"`
	Amount                *Money                `json:"amount"`
	Gateway               string                `json:"gateway,omitempty"`
	PaymentId             string                `json:"payment_id"`
	InitiatedAt           time.Time             `json:"initiated_at"`
	UpdatedAt             time.Time             `json:"updated_at"`
	ReceiptUrl            string                `json:"receipt_url,omitempty"`
	IdempotencyKey        string                `json:"idempotency_key,omitempty"`
	Currency              string                `json:"currency"`
	CompletedAt           time.Time             `json:"completed_at,omitempty"`
	TransactionId         string                `json:"transaction_id,omitempty"`
	PolicyId              string                `json:"policy_id,omitempty"`
	GatewayResponse       string                `json:"gateway_response,omitempty"`
	RetryCount            int                   `json:"retry_count"`
	FailureReason         string                `json:"failure_reason,omitempty"`
	PayerId               string                `json:"payer_id"`
	CreatedAt             time.Time             `json:"created_at"`
}
