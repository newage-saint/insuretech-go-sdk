package models

import (
	"time"
)

// PaymentRefund represents a payment_refund
type PaymentRefund struct {
	ApprovedBy      string      `json:"approved_by,omitempty"`
	ApprovedAt      time.Time   `json:"approved_at,omitempty"`
	ProcessedAt     time.Time   `json:"processed_at,omitempty"`
	UpdatedAt       time.Time   `json:"updated_at"`
	RefundId        string      `json:"refund_id"`
	PaymentId       string      `json:"payment_id"`
	RefundAmount    *Money      `json:"refund_amount"`
	Status          interface{} `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
	RefundPaymentId string      `json:"refund_payment_id,omitempty"`
	Reason          string      `json:"reason"`
}
