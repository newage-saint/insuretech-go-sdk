package models

import (
	"time"
)

// CommissionPayout represents a commission_payout
type CommissionPayout struct {
	PaymentReference string      `json:"payment_reference,omitempty"`
	AuditInfo        interface{} `json:"audit_info"`
	PayoutNumber     string      `json:"payout_number"`
	RecipientId      string      `json:"recipient_id"`
	CommissionCount  int         `json:"commission_count"`
	Status           interface{} `json:"status"`
	PaidAt           time.Time   `json:"paid_at,omitempty"`
	Id               string      `json:"id"`
	RecipientType    string      `json:"recipient_type"`
	PeriodStart      time.Time   `json:"period_start"`
	PeriodEnd        time.Time   `json:"period_end"`
	TotalAmount      *Money      `json:"total_amount,omitempty"`
	PaymentMethod    string      `json:"payment_method,omitempty"`
}
