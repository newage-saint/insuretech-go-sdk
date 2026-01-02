package models

import (
	"time"
)

// PaymentRefund represents a payment_refund
type PaymentRefund struct {
	RefundPaymentId string      `json:"refund_payment_id,omitempty"`
	RefundAmount    *Money      `json:"refund_amount"`
	Status          interface{} `json:"status"`
	ApprovedBy      string      `json:"approved_by,omitempty"`
	ApprovedAt      time.Time   `json:"approved_at,omitempty"`
	RefundId        string      `json:"refund_id"`
	PaymentId       string      `json:"payment_id"`
	Reason          string      `json:"reason"`
	ProcessedAt     time.Time   `json:"processed_at,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
