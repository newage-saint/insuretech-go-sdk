package models

import (
	"time"
)

// Refund represents a refund
type Refund struct {
	RefundNumber       string        `json:"refund_number"`
	ApprovedBy         string        `json:"approved_by,omitempty"`
	PaymentReference   string        `json:"payment_reference,omitempty"`
	Id                 string        `json:"id"`
	PolicyId           string        `json:"policy_id"`
	Reason             *RefundReason `json:"reason"`
	RefundableAmount   *Money        `json:"refundable_amount,omitempty"`
	CalculationDetails string        `json:"calculation_details,omitempty"`
	AuditInfo          interface{}   `json:"audit_info"`
	ReasonDetails      string        `json:"reason_details,omitempty"`
	PremiumUsed        *Money        `json:"premium_used,omitempty"`
	Status             interface{}   `json:"status"`
	PaymentMethod      string        `json:"payment_method,omitempty"`
	PaymentRefundId    string        `json:"payment_refund_id,omitempty"`
	TotalPremiumPaid   *Money        `json:"total_premium_paid,omitempty"`
	CancellationCharge *Money        `json:"cancellation_charge,omitempty"`
	RequestedBy        string        `json:"requested_by"`
	ProcessedAt        time.Time     `json:"processed_at,omitempty"`
}
