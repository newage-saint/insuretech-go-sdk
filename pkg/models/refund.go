package models

import (
	"time"
)

// Refund represents a refund
type Refund struct {
	TotalPremiumPaid   *Money        `json:"total_premium_paid,omitempty"`
	PremiumUsed        *Money        `json:"premium_used,omitempty"`
	RequestedBy        string        `json:"requested_by"`
	AuditInfo          interface{}   `json:"audit_info"`
	Reason             *RefundReason `json:"reason"`
	PaymentReference   string        `json:"payment_reference,omitempty"`
	PaymentMethod      string        `json:"payment_method,omitempty"`
	RefundableAmount   *Money        `json:"refundable_amount,omitempty"`
	CalculationDetails string        `json:"calculation_details,omitempty"`
	ApprovedBy         string        `json:"approved_by,omitempty"`
	PaymentRefundId    string        `json:"payment_refund_id,omitempty"`
	Id                 string        `json:"id"`
	PolicyId           string        `json:"policy_id"`
	CancellationCharge *Money        `json:"cancellation_charge,omitempty"`
	Status             interface{}   `json:"status"`
	ProcessedAt        time.Time     `json:"processed_at,omitempty"`
	RefundNumber       string        `json:"refund_number"`
	ReasonDetails      string        `json:"reason_details,omitempty"`
}
