package models

import (
	"time"
)

// Agent represents a agent
type Agent struct {
	CommissionRate float64     `json:"commission_rate"`
	JoinedAt       time.Time   `json:"joined_at"`
	CreatedAt      time.Time   `json:"created_at"`
	PartnerId      string      `json:"partner_id"`
	FullName       string      `json:"full_name"`
	PhoneNumber    string      `json:"phone_number"`
	Email          string      `json:"email,omitempty"`
	NidNumber      string      `json:"nid_number,omitempty"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      time.Time   `json:"deleted_at,omitempty"`
	AgentId        string      `json:"agent_id"`
	UserId         string      `json:"user_id,omitempty"`
	Status         interface{} `json:"status"`
}
