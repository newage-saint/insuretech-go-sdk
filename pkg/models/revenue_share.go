package models

import (
	"time"
)

// RevenueShare represents a revenue_share
type RevenueShare struct {
	PlatformShare *Money      `json:"platform_share,omitempty"`
	InsurerShare  *Money      `json:"insurer_share,omitempty"`
	SplitConfig   string      `json:"split_config,omitempty"`
	RevenueModel  string      `json:"revenue_model"`
	RecordedAt    time.Time   `json:"recorded_at"`
	AuditInfo     interface{} `json:"audit_info"`
	Id            string      `json:"id"`
	PolicyId      string      `json:"policy_id"`
	InsurerId     string      `json:"insurer_id"`
	GrossPremium  *Money      `json:"gross_premium,omitempty"`
}
