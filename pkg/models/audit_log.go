package models

import (
	"time"
)

// AuditLog represents a audit_log
type AuditLog struct {
	UserEmail  string       `json:"user_email,omitempty"`
	UserRole   string       `json:"user_role,omitempty"`
	OldValues  string       `json:"old_values,omitempty"`
	NewValues  string       `json:"new_values,omitempty"`
	AuditLogId string       `json:"audit_log_id,omitempty"`
	Changes    string       `json:"changes,omitempty"`
	IpAddress  string       `json:"ip_address,omitempty"`
	UserAgent  string       `json:"user_agent,omitempty"`
	TraceId    string       `json:"trace_id,omitempty"`
	Timestamp  time.Time    `json:"timestamp,omitempty"`
	EntityType string       `json:"entity_type,omitempty"`
	EntityId   string       `json:"entity_id,omitempty"`
	Action     *AuditAction `json:"action,omitempty"`
	UserId     string       `json:"user_id,omitempty"`
}
