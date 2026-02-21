package models

import (
	"time"
)

// CommissionPayout represents a commission_payout
type CommissionPayout struct {
	RecipientId      string      `json:"recipient_id"`
	PeriodStart      time.Time   `json:"period_start"`
	PeriodEnd        time.Time   `json:"period_end"`
	CommissionCount  int         `json:"commission_count"`
	PaidAt           time.Time   `json:"paid_at,omitempty"`
	Id               string      `json:"id"`
	RecipientType    string      `json:"recipient_type"`
	TotalAmount      *Money      `json:"total_amount,omitempty"`
	Status           interface{} `json:"status"`
	PaymentMethod    string      `json:"payment_method,omitempty"`
	PaymentReference string      `json:"payment_reference,omitempty"`
	AuditInfo        interface{} `json:"audit_info"`
	PayoutNumber     string      `json:"payout_number"`
}
