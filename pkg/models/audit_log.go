package models

import (
	"time"
)

// AuditLog represents a audit_log
type AuditLog struct {
	UserAgent  string       `json:"user_agent,omitempty"`
	TraceId    string       `json:"trace_id,omitempty"`
	Timestamp  time.Time    `json:"timestamp"`
	Id         string       `json:"id"`
	EntityType string       `json:"entity_type"`
	EntityId   string       `json:"entity_id"`
	Action     *AuditAction `json:"action"`
	UserId     string       `json:"user_id,omitempty"`
	UserEmail  string       `json:"user_email,omitempty"`
	UserRole   string       `json:"user_role,omitempty"`
	NewValues  string       `json:"new_values,omitempty"`
	OldValues  string       `json:"old_values,omitempty"`
	Changes    string       `json:"changes,omitempty"`
	IpAddress  string       `json:"ip_address,omitempty"`
}
