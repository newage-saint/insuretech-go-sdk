package models

import (
	"time"
)

// Agent represents a agent
type Agent struct {
	FullName       string      `json:"full_name"`
	PhoneNumber    string      `json:"phone_number"`
	Status         interface{} `json:"status"`
	JoinedAt       time.Time   `json:"joined_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      time.Time   `json:"deleted_at,omitempty"`
	AgentId        string      `json:"agent_id"`
	Email          string      `json:"email,omitempty"`
	CommissionRate float64     `json:"commission_rate"`
	NidNumber      string      `json:"nid_number,omitempty"`
	CreatedAt      time.Time   `json:"created_at"`
	PartnerId      string      `json:"partner_id"`
	UserId         string      `json:"user_id,omitempty"`
}
