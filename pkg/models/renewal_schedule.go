package models

import (
	"time"
)

// RenewalSchedule represents a renewal_schedule
type RenewalSchedule struct {
	Id              string       `json:"id"`
	PolicyId        string       `json:"policy_id"`
	RenewalType     *RenewalType `json:"renewal_type"`
	GracePeriodDays int          `json:"grace_period_days,omitempty"`
	RenewalDueDate  time.Time    `json:"renewal_due_date"`
	RenewalPremium  *Money       `json:"renewal_premium,omitempty"`
	Status          interface{}  `json:"status"`
	GracePeriodEnd  time.Time    `json:"grace_period_end,omitempty"`
	RenewedAt       time.Time    `json:"renewed_at,omitempty"`
	RenewedPolicyId string       `json:"renewed_policy_id,omitempty"`
	AuditInfo       *AuditInfo   `json:"audit_info,omitempty"`
}
