package models

import (
	"time"
)

// CommissionPayout represents a commission_payout
type CommissionPayout struct {
	TotalAmount      *Money      `json:"total_amount,omitempty"`
	CommissionCount  int         `json:"commission_count"`
	Status           interface{} `json:"status"`
	PaymentMethod    string      `json:"payment_method,omitempty"`
	RecipientId      string      `json:"recipient_id"`
	PeriodStart      time.Time   `json:"period_start"`
	PaymentReference string      `json:"payment_reference,omitempty"`
	PaidAt           time.Time   `json:"paid_at,omitempty"`
	AuditInfo        *AuditInfo  `json:"audit_info,omitempty"`
	Id               string      `json:"id"`
	PayoutNumber     string      `json:"payout_number"`
	RecipientType    string      `json:"recipient_type"`
	PeriodEnd        time.Time   `json:"period_end"`
}
