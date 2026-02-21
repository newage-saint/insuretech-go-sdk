package models

import (
	"time"
)

// Agent represents a agent
type Agent struct {
	AgentId        string      `json:"agent_id"`
	UserId         string      `json:"user_id,omitempty"`
	FullName       string      `json:"full_name"`
	PhoneNumber    string      `json:"phone_number"`
	Email          string      `json:"email,omitempty"`
	Status         interface{} `json:"status"`
	CommissionRate float64     `json:"commission_rate"`
	NidNumber      string      `json:"nid_number,omitempty"`
	PartnerId      string      `json:"partner_id"`
	JoinedAt       time.Time   `json:"joined_at"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      time.Time   `json:"deleted_at,omitempty"`
}
