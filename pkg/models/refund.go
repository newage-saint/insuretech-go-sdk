package models

import (
	"time"
)

// Refund represents a refund
type Refund struct {
	TotalPremiumPaid   *Money        `json:"total_premium_paid,omitempty"`
	RefundableAmount   *Money        `json:"refundable_amount,omitempty"`
	ApprovedBy         string        `json:"approved_by,omitempty"`
	PaymentMethod      string        `json:"payment_method,omitempty"`
	ProcessedAt        time.Time     `json:"processed_at,omitempty"`
	RefundNumber       string        `json:"refund_number"`
	Reason             *RefundReason `json:"reason"`
	PremiumUsed        *Money        `json:"premium_used,omitempty"`
	CalculationDetails string        `json:"calculation_details,omitempty"`
	Status             interface{}   `json:"status"`
	Id                 string        `json:"id"`
	CancellationCharge *Money        `json:"cancellation_charge,omitempty"`
	RequestedBy        string        `json:"requested_by"`
	PaymentReference   string        `json:"payment_reference,omitempty"`
	PolicyId           string        `json:"policy_id"`
	ReasonDetails      string        `json:"reason_details,omitempty"`
	PaymentRefundId    string        `json:"payment_refund_id,omitempty"`
	AuditInfo          interface{}   `json:"audit_info"`
}
