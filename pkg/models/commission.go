package models

import (
	"time"
)

// Commission represents a commission
type Commission struct {
	PolicyId         string          `json:"policy_id"`
	CommissionAmount *Money          `json:"commission_amount"`
	CommissionRate   float64         `json:"commission_rate"`
	PaidAt           time.Time       `json:"paid_at,omitempty"`
	UpdatedAt        time.Time       `json:"updated_at"`
	PartnerId        string          `json:"partner_id,omitempty"`
	AgentId          string          `json:"agent_id,omitempty"`
	Type             *CommissionType `json:"type"`
	Status           interface{}     `json:"status"`
	PaymentId        string          `json:"payment_id,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	CommissionId     string          `json:"commission_id"`
}
