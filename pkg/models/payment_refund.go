package models

import (
	"time"
)

// PaymentRefund represents a payment_refund
type PaymentRefund struct {
	UpdatedAt       time.Time   `json:"updated_at"`
	RefundId        string      `json:"refund_id"`
	RefundAmount    *Money      `json:"refund_amount"`
	ApprovedBy      string      `json:"approved_by,omitempty"`
	ApprovedAt      time.Time   `json:"approved_at,omitempty"`
	ProcessedAt     time.Time   `json:"processed_at,omitempty"`
	PaymentId       string      `json:"payment_id"`
	RefundPaymentId string      `json:"refund_payment_id,omitempty"`
	Reason          string      `json:"reason"`
	Status          interface{} `json:"status"`
	CreatedAt       time.Time   `json:"created_at"`
}
