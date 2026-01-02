package models

import (
	"time"
)

// Refund represents a refund
type Refund struct {
	ProcessedAt        time.Time     `json:"processed_at,omitempty"`
	AuditInfo          *AuditInfo    `json:"audit_info,omitempty"`
	Id                 string        `json:"id"`
	PolicyId           string        `json:"policy_id"`
	ReasonDetails      string        `json:"reason_details,omitempty"`
	RefundableAmount   *Money        `json:"refundable_amount,omitempty"`
	PaymentMethod      string        `json:"payment_method,omitempty"`
	PaymentReference   string        `json:"payment_reference,omitempty"`
	RefundNumber       string        `json:"refund_number"`
	TotalPremiumPaid   *Money        `json:"total_premium_paid,omitempty"`
	PremiumUsed        *Money        `json:"premium_used,omitempty"`
	CancellationCharge *Money        `json:"cancellation_charge,omitempty"`
	ApprovedBy         string        `json:"approved_by,omitempty"`
	RequestedBy        string        `json:"requested_by"`
	PaymentRefundId    string        `json:"payment_refund_id,omitempty"`
	Reason             *RefundReason `json:"reason"`
	CalculationDetails string        `json:"calculation_details,omitempty"`
	Status             interface{}   `json:"status"`
}
