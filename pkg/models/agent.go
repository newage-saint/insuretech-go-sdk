package models

import (
	"time"
)

// Agent represents a agent
type Agent struct {
	PhoneNumber    string      `json:"phone_number"`
	CommissionRate float64     `json:"commission_rate"`
	CreatedAt      time.Time   `json:"created_at"`
	DeletedAt      time.Time   `json:"deleted_at,omitempty"`
	PartnerId      string      `json:"partner_id"`
	Email          string      `json:"email,omitempty"`
	Status         interface{} `json:"status"`
	NidNumber      string      `json:"nid_number,omitempty"`
	JoinedAt       time.Time   `json:"joined_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	AgentId        string      `json:"agent_id"`
	UserId         string      `json:"user_id,omitempty"`
	FullName       string      `json:"full_name"`
}
