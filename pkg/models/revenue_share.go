package models

import (
	"time"
)

// RevenueShare represents a revenue_share
type RevenueShare struct {
	AuditInfo     interface{} `json:"audit_info"`
	PolicyId      string      `json:"policy_id"`
	PlatformShare *Money      `json:"platform_share,omitempty"`
	SplitConfig   string      `json:"split_config,omitempty"`
	RecordedAt    time.Time   `json:"recorded_at"`
	Id            string      `json:"id"`
	InsurerId     string      `json:"insurer_id"`
	GrossPremium  *Money      `json:"gross_premium,omitempty"`
	InsurerShare  *Money      `json:"insurer_share,omitempty"`
	RevenueModel  string      `json:"revenue_model"`
}
