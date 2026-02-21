package models

import (
	"time"
)

// CommissionPayout represents a commission_payout
type CommissionPayout struct {
	CommissionCount  int         `json:"commission_count"`
	PaymentMethod    string      `json:"payment_method,omitempty"`
	PaymentReference string      `json:"payment_reference,omitempty"`
	PaidAt           time.Time   `json:"paid_at,omitempty"`
	PayoutNumber     string      `json:"payout_number"`
	RecipientId      string      `json:"recipient_id"`
	PeriodStart      time.Time   `json:"period_start"`
	TotalAmount      *Money      `json:"total_amount,omitempty"`
	Status           interface{} `json:"status"`
	AuditInfo        interface{} `json:"audit_info"`
	Id               string      `json:"id"`
	RecipientType    string      `json:"recipient_type"`
	PeriodEnd        time.Time   `json:"period_end"`
}
