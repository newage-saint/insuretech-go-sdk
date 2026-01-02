package models

import (
	"time"
)

// Commission represents a commission
type Commission struct {
	Status           interface{}     `json:"status"`
	PaymentId        string          `json:"payment_id,omitempty"`
	PaidAt           time.Time       `json:"paid_at,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	CommissionId     string          `json:"commission_id"`
	PolicyId         string          `json:"policy_id"`
	PartnerId        string          `json:"partner_id,omitempty"`
	AgentId          string          `json:"agent_id,omitempty"`
	Type             *CommissionType `json:"type"`
	CommissionAmount *Money          `json:"commission_amount"`
	CommissionRate   float64         `json:"commission_rate"`
}
