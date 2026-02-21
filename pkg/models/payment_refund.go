package models

import (
	"time"
)

// PaymentRefund represents a payment_refund
type PaymentRefund struct {
	PaymentId       string      `json:"payment_id"`
	RefundPaymentId string      `json:"refund_payment_id,omitempty"`
	RefundAmount    *Money      `json:"refund_amount"`
	Reason          string      `json:"reason"`
	ApprovedAt      time.Time   `json:"approved_at,omitempty"`
	ProcessedAt     time.Time   `json:"processed_at,omitempty"`
	RefundId        string      `json:"refund_id"`
	Status          interface{} `json:"status"`
	ApprovedBy      string      `json:"approved_by,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
