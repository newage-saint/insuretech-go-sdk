package models

import (
	"time"
)

// CommissionPayout represents a commission_payout
type CommissionPayout struct {
	PaidAt           time.Time   `json:"paid_at,omitempty"`
	AuditInfo        interface{} `json:"audit_info"`
	Id               string      `json:"id"`
	PayoutNumber     string      `json:"payout_number"`
	RecipientType    string      `json:"recipient_type"`
	TotalAmount      *Money      `json:"total_amount,omitempty"`
	CommissionCount  int         `json:"commission_count"`
	Status           interface{} `json:"status"`
	PaymentMethod    string      `json:"payment_method,omitempty"`
	PaymentReference string      `json:"payment_reference,omitempty"`
	RecipientId      string      `json:"recipient_id"`
	PeriodStart      time.Time   `json:"period_start"`
	PeriodEnd        time.Time   `json:"period_end"`
}
