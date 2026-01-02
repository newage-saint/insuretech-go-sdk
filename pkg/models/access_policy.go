package models

import (
	"time"
)

// AccessPolicy represents a access_policy
type AccessPolicy struct {
	Resource   string       `json:"resource"`
	Conditions []*Condition `json:"conditions,omitempty"`
	Effect     *Effect      `json:"effect"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	IsActive   bool         `json:"is_active"`
	PolicyId   string       `json:"policy_id"`
	PolicyName string       `json:"policy_name"`
}
