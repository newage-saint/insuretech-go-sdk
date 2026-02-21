package models

import (
	"time"
)

// Refund represents a refund
type Refund struct {
	PolicyId           string        `json:"policy_id"`
	Reason             *RefundReason `json:"reason"`
	TotalPremiumPaid   *Money        `json:"total_premium_paid,omitempty"`
	RequestedBy        string        `json:"requested_by"`
	ApprovedBy         string        `json:"approved_by,omitempty"`
	AuditInfo          interface{}   `json:"audit_info"`
	RefundNumber       string        `json:"refund_number"`
	PremiumUsed        *Money        `json:"premium_used,omitempty"`
	ProcessedAt        time.Time     `json:"processed_at,omitempty"`
	CancellationCharge *Money        `json:"cancellation_charge,omitempty"`
	RefundableAmount   *Money        `json:"refundable_amount,omitempty"`
	Status             interface{}   `json:"status"`
	PaymentMethod      string        `json:"payment_method,omitempty"`
	PaymentReference   string        `json:"payment_reference,omitempty"`
	Id                 string        `json:"id"`
	ReasonDetails      string        `json:"reason_details,omitempty"`
	CalculationDetails string        `json:"calculation_details,omitempty"`
	PaymentRefundId    string        `json:"payment_refund_id,omitempty"`
}
