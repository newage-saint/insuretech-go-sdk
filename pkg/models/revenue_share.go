package models

import (
	"time"
)

// RevenueShare represents a revenue_share
type RevenueShare struct {
	GrossPremium  *Money     `json:"gross_premium,omitempty"`
	PlatformShare *Money     `json:"platform_share,omitempty"`
	RevenueModel  string     `json:"revenue_model"`
	InsurerShare  *Money     `json:"insurer_share,omitempty"`
	SplitConfig   string     `json:"split_config,omitempty"`
	RecordedAt    time.Time  `json:"recorded_at"`
	AuditInfo     *AuditInfo `json:"audit_info,omitempty"`
	Id            string     `json:"id"`
	PolicyId      string     `json:"policy_id"`
	InsurerId     string     `json:"insurer_id"`
}
