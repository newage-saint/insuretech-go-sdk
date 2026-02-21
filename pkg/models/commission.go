package models

import (
	"time"
)

// Commission represents a commission
type Commission struct {
	Type             *CommissionType `json:"type"`
	CommissionRate   float64         `json:"commission_rate"`
	Status           interface{}     `json:"status"`
	PaymentId        string          `json:"payment_id,omitempty"`
	PaidAt           time.Time       `json:"paid_at,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	CommissionId     string          `json:"commission_id"`
	PolicyId         string          `json:"policy_id"`
	PartnerId        string          `json:"partner_id,omitempty"`
	CommissionAmount *Money          `json:"commission_amount"`
	AgentId          string          `json:"agent_id,omitempty"`
}
