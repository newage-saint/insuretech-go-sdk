package models

import (
	"time"
)

// RevenueShare represents a revenue_share
type RevenueShare struct {
	InsurerId     string      `json:"insurer_id"`
	PlatformShare *Money      `json:"platform_share,omitempty"`
	InsurerShare  *Money      `json:"insurer_share,omitempty"`
	SplitConfig   string      `json:"split_config,omitempty"`
	Id            string      `json:"id"`
	PolicyId      string      `json:"policy_id"`
	GrossPremium  *Money      `json:"gross_premium,omitempty"`
	RevenueModel  string      `json:"revenue_model"`
	RecordedAt    time.Time   `json:"recorded_at"`
	AuditInfo     interface{} `json:"audit_info"`
}
