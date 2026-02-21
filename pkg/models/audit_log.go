package models

import (
	"time"
)

// AuditLog represents a audit_log
type AuditLog struct {
	IpAddress  string       `json:"ip_address,omitempty"`
	UserAgent  string       `json:"user_agent,omitempty"`
	TraceId    string       `json:"trace_id,omitempty"`
	Timestamp  time.Time    `json:"timestamp,omitempty"`
	AuditLogId string       `json:"audit_log_id,omitempty"`
	EntityType string       `json:"entity_type,omitempty"`
	EntityId   string       `json:"entity_id,omitempty"`
	UserId     string       `json:"user_id,omitempty"`
	NewValues  string       `json:"new_values,omitempty"`
	Changes    string       `json:"changes,omitempty"`
	Action     *AuditAction `json:"action,omitempty"`
	UserEmail  string       `json:"user_email,omitempty"`
	UserRole   string       `json:"user_role,omitempty"`
	OldValues  string       `json:"old_values,omitempty"`
}
