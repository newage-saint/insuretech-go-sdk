package models

import (
	"time"
)

// Commission represents a commission
type Commission struct {
	PolicyId         string          `json:"policy_id"`
	PartnerId        string          `json:"partner_id,omitempty"`
	CommissionAmount *Money          `json:"commission_amount"`
	Status           interface{}     `json:"status"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	AgentId          string          `json:"agent_id,omitempty"`
	Type             *CommissionType `json:"type"`
	CommissionRate   float64         `json:"commission_rate"`
	PaymentId        string          `json:"payment_id,omitempty"`
	PaidAt           time.Time       `json:"paid_at,omitempty"`
	CommissionId     string          `json:"commission_id"`
}
