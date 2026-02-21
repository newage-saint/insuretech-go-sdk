package models

import (
	"time"
)

// Commission represents a commission
type Commission struct {
	PaidAt           time.Time       `json:"paid_at,omitempty"`
	UpdatedAt        time.Time       `json:"updated_at"`
	PartnerId        string          `json:"partner_id,omitempty"`
	Type             *CommissionType `json:"type"`
	CommissionAmount *Money          `json:"commission_amount"`
	Status           interface{}     `json:"status"`
	CreatedAt        time.Time       `json:"created_at"`
	CommissionId     string          `json:"commission_id"`
	PolicyId         string          `json:"policy_id"`
	AgentId          string          `json:"agent_id,omitempty"`
	CommissionRate   float64         `json:"commission_rate"`
	PaymentId        string          `json:"payment_id,omitempty"`
}
