package models

import (
	"time"
)

// AuditLog represents a audit_log
type AuditLog struct {
	UserRole   string       `json:"user_role,omitempty"`
	Changes    string       `json:"changes,omitempty"`
	IpAddress  string       `json:"ip_address,omitempty"`
	UserAgent  string       `json:"user_agent,omitempty"`
	EntityId   string       `json:"entity_id,omitempty"`
	Action     *AuditAction `json:"action,omitempty"`
	OldValues  string       `json:"old_values,omitempty"`
	NewValues  string       `json:"new_values,omitempty"`
	TraceId    string       `json:"trace_id,omitempty"`
	Timestamp  time.Time    `json:"timestamp,omitempty"`
	AuditLogId string       `json:"audit_log_id,omitempty"`
	EntityType string       `json:"entity_type,omitempty"`
	UserId     string       `json:"user_id,omitempty"`
	UserEmail  string       `json:"user_email,omitempty"`
}
